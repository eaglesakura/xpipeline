package mipmap

import (
	"../utils"
	"../image"
	"errors"
	"github.com/urfave/cli"
	"strings"
	"fmt"
	"os"
)

var _DPI_TABLE = []*image.DotPerInch{
	&image.DotPerInch{Name: "ldpi"},
	&image.DotPerInch{Name: "mdpi"},
	&image.DotPerInch{Name: "hdpi"},
	&image.DotPerInch{Name: "xhdpi"},
	&image.DotPerInch{Name: "xxhdpi"},
	&image.DotPerInch{Name: "xxxhdpi"},
}

//
// 画像のMipmap処理を行う
// リサイズのみを担当し、webp等への変換は担当しない
//
type Task struct {
	ctx *cli.Context // CommandLine Context

	configDirectory string // config.yaml が格納されているディレクトリ, "/"は含まない
	config          Configure
}

func NewInstance(ctx *cli.Context) (*Task, error) {
	result := &Task{
	}

	if configFile := ctx.String("config"); configFile == "" {
		return nil, errors.New("-config path/to/config.yaml")
	} else {
		// load config
		if parseError := utils.ParseYamlFromFile(configFile, &result.config); parseError != nil {
			return nil, parseError
		}

		// set yaml path
		if strings.Index(configFile, "/") > 0 {
			result.configDirectory = configFile[0:strings.LastIndex(configFile, "/")]
		}
	}

	return result, nil
}

/*
 処理対象のdpi一覧を取得する
*/
func (it *Task) getDpiList(path string) []*image.DotPerInch {
	var result []*image.DotPerInch
	for _, dir := range utils.ListDirectories(it.configDirectory + "/" + path) {
		result = append(result, &image.DotPerInch{
			Name: dir.Name(),
		})
	}
	return result
}

/*
 1ファイル単位でmipmapを生成する
*/
func (it *Task) generateMipmap(srcDpi *image.DotPerInch, src os.FileInfo, outputPath string, request Request) error {

	// 出力ファイル名を決定する
	dstFileName := src.Name()
	if len(request.Format) > 0 {
		// フォーマット変換が必要
		dstFileName = src.Name()[0:strings.LastIndex(src.Name(), ".")] + "." + strings.ToLower(request.Format)
	}

	srcFilePath := it.configDirectory + "/" + request.Path + "/" + srcDpi.Name + "/" + src.Name()

	// 画像情報
	info, err := image.NewImageInstance(srcFilePath)
	if err != nil {
		return err
	}

	fmt.Printf("convert %v[%vx%v:%v] %v\n", srcDpi.Name, info.Width, info.Height, info.Format, srcFilePath)
	for _, dstDpi := range _DPI_TABLE {

		dstWidth := srcDpi.GetResizePixels(info.Width, dstDpi)
		dstHeight := srcDpi.GetResizePixels(info.Height, dstDpi)

		if dstWidth <= 0 || dstHeight <= 0 {
			continue
		}

		// convert経由で出力する
		os.MkdirAll(outputPath+"/"+request.Type+"-"+dstDpi.Name, os.ModePerm)
		dstFilePath := outputPath + "/" + request.Type + "-" + dstDpi.Name + "/" + dstFileName

		// 出力ファイルが存在したらskip
		if _, err := os.Stat(dstFilePath); err == nil {
			// ファイルが存在するので、出力しない
			fmt.Printf("  - %v[%vx%v] Skip \n", dstDpi.Name, dstWidth, dstHeight)
			continue
		}

		fmt.Printf("  * %v[%vx%v] %v \n", dstDpi.Name, dstWidth, dstHeight, dstFilePath)
		cmd := &utils.ExternalCommand{
			Commands: []string{
				"convert", srcFilePath,
			},
		}
		// リサイズの必要があるなら設定
		if dstWidth != info.Width || dstHeight != info.Height {
			cmd.Commands = append(cmd.Commands, "-resize", fmt.Sprintf("%vx%v", dstWidth, dstHeight))
		}

		// 引数を追加する
		for _, arg := range request.ConvertArgs {
			cmd.Commands = append(cmd.Commands, arg)
		}
		// 出力ファイルパス
		cmd.Commands = append(cmd.Commands, dstFilePath)

		// execute!
		_, stdErr, err := cmd.RunStdout()
		if err != nil {
			return errors.New(fmt.Sprintf("%v %v", err, stdErr))
		}
	}

	return nil
}

/*
 1リクエストを処理する
*/
func (it *Task) executeAndroid(request Request) error {
	// 出力先ディレクトリを作成する
	outputDirectoryPath := it.configDirectory + "/" + request.OutputPath
	os.MkdirAll(outputDirectoryPath, os.ModePerm)

	// dpi一覧を取得する
	for _, dpi := range it.getDpiList(request.Path) {
		// dpi内部のファイルを列挙する
		path := it.configDirectory + "/" + request.Path + "/" + dpi.Name
		for _, srcFile := range utils.ListFiles(path) {
			// 1ファイルの生成を行う
			if err := it.generateMipmap(dpi, srcFile, outputDirectoryPath, request); err != nil {
				return err
			}
		}
	}

	return nil
}

func (it *Task) Execute() {
	for _, req := range it.config.Mipmap.Requests {
		if req.Platform == "android" {
			err := it.executeAndroid(req)
			if err != nil {
				fmt.Errorf("mipmap failed : path=%v %v\n", req.Path, err)
				return
			}
		} else {
			fmt.Errorf("mipmap failed : unknown type[%v] path=%v\n", req.Type, req.Path)
			return
		}
	}
}

type Request struct {
	Path        string   `yaml:"path"`         // ファイル一覧へのパス
	Platform    string   `yaml:"platform"`     // 対象プラットフォーム
	Type        string   `yaml:"type"`         // 出力タイプ
	OutputPath  string   `yaml:"outpath"`      // 出力先ディレクトリ
	Format      string   `yaml:"format"`       // 出力ファイルフォーマット
	ConvertArgs []string `yaml:"convert_args"` // convertコマンドへ渡される引数
}

// mipmap出力設定ファイル
type Configure struct {
	Mipmap struct {
		Requests []Request
	}
}

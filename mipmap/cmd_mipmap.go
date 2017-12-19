package mipmap

import (
	"../utils"
	"errors"
	"github.com/urfave/cli"
)

//
// 画像のMipmap処理を行う
// リサイズのみを担当し、webp等への変換は担当しない
//
type MipmapTask struct {
	ctx *cli.Context // CommandLine Context

	config Configure
}

func NewInstance(ctx *cli.Context) (*MipmapTask, error) {
	result := &MipmapTask{
	}

	if configFile := ctx.String("config"); configFile == "" {
		return nil, errors.New("-config path/to/config.yaml")
	} else {
		// load config
		if parseError := utils.ParseYamlFromFile(configFile, &result.config); parseError != nil {
			return nil, parseError
		}
	}

	return result, nil
}

func (it *MipmapTask) Execute() {

}

// mipmap出力設定ファイル
type Configure struct {
	Mipmap struct {
		Requests []*struct {
			// ファイル一覧へのパス
			Path string `yaml:"path"`

			// 対象プラットフォーム
			Platform string `yaml:"platform"`

			// 出力タイプ
			Type string `yaml:"type"`
		}
	}
}

package image

import (
	"../utils"
	"fmt"
	"errors"
	"strings"
)

/*
 画像情報構造体
*/
type ImageInfo struct {
	Path     string // リクエストされたパス
	Width    int    // 幅ピクセル数
	Height   int    // 高さピクセル数
	Format   string // 画像フォーマット
	Channels string // カラーチャンネル
}

/*
 画像情報を読み込む
*/
func LoadImageInfo(path string) (*ImageInfo, error) {
	//if currentDir, _ := os.Getwd(); currentDir != "" {
	//	path = currentDir + "/" + path
	//}

	cmd := &utils.ExternalCommand{
		Commands: []string{
			"identify",
			"-format", "%w,%h,%m,%[channels]",
			path,
		},
	}

	if info, errInfo, err := cmd.RunStdout(); err != nil {
		return nil, errors.New(fmt.Sprintf("%v / %v", err, errInfo))
	} else {
		infoList := strings.Split(info, ",")
		//fmt.Printf("Image[%v] Check[%v]", path, info)
		return &ImageInfo{
			Path:     path,
			Width:    utils.Atoi(infoList[0]),
			Height:   utils.Atoi(infoList[1]),
			Format:   infoList[2],
			Channels: infoList[3],
		}, nil
	}
}

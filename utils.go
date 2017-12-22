package main

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strconv"
	"os"
)

/*
 YAMLからstructへのパースを行う
*/
func ParseYamlFromFile(path string, result interface{}) error {
	if bytes, e := ioutil.ReadFile(path); e != nil {
		return e
	} else {
		return yaml.Unmarshal(bytes, result)
	}
}

func Atoi(ascii string) int {
	result, err := strconv.Atoi(ascii)
	if err != nil {
		return 0
	} else {
		return result
	}
}

/*
 指定パスに存在するディレクトリ一覧を取得する
*/
func ListDirectories(path string) []os.FileInfo {
	files, _ := ioutil.ReadDir(path)
	var result []os.FileInfo

	for _, info := range files {
		if info.IsDir() {
			result = append(result, info)
		}
	}
	return result
}

/*
 指定パスに存在するディレクトリ一覧を取得する
*/
func ListFiles(path string) []os.FileInfo {
	files, _ := ioutil.ReadDir(path)
	var result []os.FileInfo
	for _, info := range files {
		if !info.IsDir() {
			result = append(result, info)
		}
	}
	return result
}

/*
 一時ファイルの格納パスを取得する
*/
func GetTempFilePath(path string) string {
	return ".xpipeline/temp/" + path
}

/*
 一時ファイル格納パスを生成する
*/
func NewTempFileDirectory() {
	os.MkdirAll(".xpipeline/temp", os.ModePerm)
	ioutil.WriteFile(".xpipeline/temp/.gitignore", []byte(`
*.bin
*.json
*.yaml
*.dat
*.db
`), os.ModePerm)

}

/*
 一時ファイルを削除する
*/
func DeleteTempFiles() {
	os.RemoveAll(".xpipeline/temp")
}

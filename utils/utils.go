package utils

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

package utils

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strconv"
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

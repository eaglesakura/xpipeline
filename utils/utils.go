package utils

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

func ParseYamlFromFile(path string, result interface{}) error {
	if bytes, e := ioutil.ReadFile(path); e != nil {
		return e
	} else {
		return yaml.Unmarshal(bytes, result)
	}
}

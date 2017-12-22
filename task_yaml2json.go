package main

import (
	"github.com/urfave/cli"
	"github.com/ghodss/yaml"
	"errors"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

type Yaml2JsonTask struct {
	ctx *cli.Context

	srcPath string
	dstPath string
}

func newYaml2JsonTask(ctx *cli.Context) (*Yaml2JsonTask, error) {

	source := ctx.String("input")
	if source == "" {
		return nil, errors.New("-input \"path/to/input.[yaml|yml]\"")
	}

	output := ctx.String("output")
	return &Yaml2JsonTask{
		ctx:     ctx,
		srcPath: source,
		dstPath: output,
	}, nil
}

func (it *Yaml2JsonTask) Execute() {
	yamlBytes, err := ioutil.ReadFile(it.srcPath)
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}

	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}

	if it.dstPath == "" {
		// 標準出力に出力しておしまい
		fmt.Print(string(jsonBytes))
	} else {
		// ファイルとして書き出す
		os.MkdirAll(it.dstPath[0:strings.LastIndex(it.dstPath, "/")], os.ModePerm)
		err = ioutil.WriteFile(it.dstPath, jsonBytes, os.ModePerm)
		if err != nil {
			fmt.Errorf("%v\n", err)
			return
		}
		fmt.Printf("%v[%v bytes] -> %v[%v bytes]", it.srcPath, len(yamlBytes), it.dstPath, len(jsonBytes))
	}
}

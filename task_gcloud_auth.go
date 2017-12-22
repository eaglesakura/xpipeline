package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type GcloudServiceAccount struct {
	ProjectId   string `json:"project_id,omitempty"`   // GCP プロジェクトID
	ClientEmail string `json:"client_email,omitempty"` // GCP サービスアカウントメール
}

/*
 gcloud サービスアカウント認証をサポートする
*/
type GcloudAuthTask struct {
	ctx *cli.Context

	serviceAccountPath      string
	serviceAccountFileBytes []byte
	serviceAccount          *GcloudServiceAccount
}

func newGcloudAuthTask(ctx *cli.Context) (*GcloudAuthTask, error) {

	serviceAccountPath := ctx.String("key-file")
	var serviceAccountFileBytes []byte
	if serviceAccountPath == "" {
		// 標準入力から取得する
		serviceAccountFileBytes, _ = ioutil.ReadAll(os.Stdin)
	} else {
		// ファイルからロードする
		bytes, err := ioutil.ReadFile(serviceAccountPath)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%v load failed[err:%v]", serviceAccountPath, err))
		}
		serviceAccountFileBytes = bytes
	}
	serviceAccount := &GcloudServiceAccount{}

	if err := json.Unmarshal(serviceAccountFileBytes, serviceAccount);
		err != nil || serviceAccount.ProjectId == "" || serviceAccount.ClientEmail == "" {
		return nil, errors.New(fmt.Sprintf("%v not valid", serviceAccountPath))
	}

	return &GcloudAuthTask{
		ctx:                     ctx,
		serviceAccountPath:      serviceAccountPath,
		serviceAccountFileBytes: serviceAccountFileBytes,
		serviceAccount:          serviceAccount,
	}, nil
}

func (it *GcloudAuthTask) Execute() {

	if it.serviceAccountPath == "" {
		// 一時ファイルに書き出す
		path := GetTempFilePath("temp.bin")
		ioutil.WriteFile(path, it.serviceAccountFileBytes, os.ModePerm)
		it.serviceAccountPath = path
	}

	// 認証
	{
		shell := &Shell{
			Commands: []string{
				"gcloud", "auth",
				"activate-service-account", it.serviceAccount.ClientEmail,
				"--key-file", it.serviceAccountPath,
				"--project", it.serviceAccount.ProjectId,
			},
		}
		_, stderr, err := shell.RunStdout()
		if err != nil {
			fmt.Errorf("%v %v", err, stderr)
			return
		}
	}
	// config
	{
		shell := &Shell{
			Commands: []string{
				"gcloud", "config",
				"set", "project", it.serviceAccount.ProjectId,
			},
		}
		_, stderr, err := shell.RunStdout()
		if err != nil {
			fmt.Errorf("%v %v", err, stderr)
			return
		}
	}
	// config
	{
		shell := &Shell{
			Commands: []string{
				"gcloud", "config",
				"set", "account", it.serviceAccount.ClientEmail,
			},
		}
		_, stderr, err := shell.RunStdout()
		if err != nil {
			fmt.Errorf("%v %v", err, stderr)
			return
		}
	}

	fmt.Printf("gcloud auth\n")
	fmt.Printf("   * project-id   %v\n", it.serviceAccount.ProjectId)
	fmt.Printf("   * client-email %v\n", it.serviceAccount.ClientEmail)

}

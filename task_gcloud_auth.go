package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
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

	serviceAccountPath string
	serviceAccount     *GcloudServiceAccount
}

func newGcloudAuthTask(ctx *cli.Context) (*GcloudAuthTask, error) {

	serviceAccountPath := ctx.String("key-file")
	serviceAccountFileBytes, err := ioutil.ReadFile(serviceAccountPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v load failed[err:%v]", serviceAccountPath, err))
	}
	serviceAccount := &GcloudServiceAccount{}

	if err := json.Unmarshal(serviceAccountFileBytes, serviceAccount);
		err != nil || serviceAccount.ProjectId == "" || serviceAccount.ClientEmail == "" {
		return nil, errors.New(fmt.Sprintf("%v not valid", serviceAccountPath))
	}

	return &GcloudAuthTask{
		ctx:                ctx,
		serviceAccountPath: serviceAccountPath,
		serviceAccount:     serviceAccount,
	}, nil
}

func (it *GcloudAuthTask) Execute() {
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
}

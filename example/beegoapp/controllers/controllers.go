package controllers

import (
	"github.com/jungju/circle_manager/example/beegoapp/envs"
	"github.com/jungju/cognitor"
)

var cognitoClient *cognitor.CognitoClient

func init() {
	if envs.AWSClientID != "" &&
		envs.AWSClientSecret != "" &&
		envs.AWSRegion != "" && envs.AWSUserPoolID != "" {

		cognitoClient = &cognitor.CognitoClient{
			AWSClientID:     envs.AWSClientID,
			AWSClientSecret: envs.AWSClientSecret,
			AWSRegion:       envs.AWSRegion,
			AWSUserPoolID:   envs.AWSUserPoolID,
			AWSLogLevel:     1,
		}

		if err := cognitoClient.Init(); err != nil {
			panic(err)
		}
	}
}

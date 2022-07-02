package helixf_env

import (
	"errors"
	"os"
)

var HelixfEnv string
var RootPath string = os.Getenv("HELIXF_ROOT")
var ChannelAccessToken string = os.Getenv("LINE_ACCESS_TOKEN")
var UrlHost string = os.Getenv("URL_HOST")

func init() {
	HelixfEnv = os.Getenv("HELIXF_ENV")
	if HelixfEnv == "" {
		HelixfEnv = "development"
	}

	if RootPath == "" {
		panic(errors.New("env variables is not configed: HELIXF_ROOT"))
	}
}

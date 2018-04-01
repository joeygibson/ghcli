package commands

import (
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/sirupsen/logrus"
	"strings"
)

func Login(args []string) {
	if len(args) == 0 {
		logrus.Fatal("you must specify your Github token")
	}

	token := args[0]

	token = strings.TrimSpace(token)

	if len(token) == 0 {
		logrus.Fatalf("invalid token")
	}

	conf, err := config.LoadFromDisk()
	if err != nil {
		logrus.Fatalf("reading config file: %v", err)
	}

	conf.Token = token

	err = conf.Save()
	if err != nil {
		logrus.Fatal(err)
	}
}

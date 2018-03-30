package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

type Config struct {
	Org      string
	User     string
	Password string
	Token    string
	Top      int
}

var conf *Config

// GetConfig returns the current config, loading it from CLI/env/config the first time.
func GetConfig() *Config {
	if conf == nil {
		token := readTokenFile(viper.GetString("token.file"))

		conf = &Config{
			Org:      viper.GetString("org"),
			User:     viper.GetString("user"),
			Password: viper.GetString("password"),
			Token:    token,
			Top:      viper.GetInt("top"),
		}

		if conf.Org == "" {
			logrus.Fatal("no organization specified")
		}
	}

	return conf
}

func readTokenFile(tokenFile string) string {
	_, err := os.Stat(tokenFile)

	if os.IsNotExist(err) {
		if tokenFile == GetDefaultTokenFileName() {
			return ""
		} else {
			logrus.Fatalf("OAuth token file %s not found", tokenFile)
		}
	}

	contents, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		logrus.Fatalf("%s exists, but cannot be read: %v", tokenFile, err)
	}

	return strings.TrimSpace(string(contents))
}

// GetDefaultTokenFileName crafts a filename for the OAuth token, which lives in the user's
// home directory.
func GetDefaultTokenFileName() string {
	currentUser, err := user.Current()
	if err != nil {
		logrus.Fatalf("unable to determine current user: %v", err)
	}

	return fmt.Sprintf("%s/.ghcli-token", currentUser.HomeDir)
}

package config

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"os/user"
)

type Config struct {
	Org      string `yaml:"org,omitempty"`
	User     string `yaml:"-"`
	Password string `yaml:"-"`
	Token    string `yaml:"token,omitempty"`
	Top      int    `yaml:"top,omitempty"`
}

var conf *Config

// GetConfig returns the current config, loading it from CLI/env/config the first time.
func GetConfig() *Config {
	if conf == nil {
		conf = &Config{
			Org:      viper.GetString("org"),
			User:     viper.GetString("user"),
			Password: viper.GetString("password"),
			Token:    viper.GetString("token"),
			Top:      viper.GetInt("top"),
		}

		if conf.Org == "" {
			logrus.Fatal("no organization specified")
		}
	}

	return conf
}

// Save writes out the current configuration to disk
func (c *Config) Save() error {
	currentUser, err := user.Current()
	if err != nil {
		logrus.Fatalf("unable to determine current user: %v", err)
	}

	fileName := fmt.Sprintf("%s/.ghcli.yaml", currentUser.HomeDir)

	file, err := os.Create(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("creating %s: %v", fileName))
	}

	defer file.Close()

	encoder := yaml.NewEncoder(file)

	err = encoder.Encode(c)
	if err != nil {
		return errors.New(fmt.Sprintf("writing config: %v", err))
	}

	return nil
}

// LoadFromDisk will load a config file from the user's home directory,
// if it exists.
func LoadFromDisk() (*Config, error) {
	currentUser, err := user.Current()
	if err != nil {
		logrus.Fatalf("unable to determine current user: %v", err)
	}

	fileName := fmt.Sprintf("%s/.ghcli.yaml", currentUser.HomeDir)

	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		return &Config{}, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("opening %s: %v", fileName))
	}

	defer file.Close()

	var conf Config

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&conf)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("decoding config %s: %v", fileName, err))
	}

	return &conf, nil
}

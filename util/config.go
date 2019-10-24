package util

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func findConfigFile(configFile string) (string, error) {
	if FileExists(configFile) == nil {
		return configFile, nil
	}

	if ex, err := os.Executable(); err == nil {
		exPath := filepath.Dir(ex)
		configFile = filepath.Join(exPath, "config.toml")
	}

	if FileExists(configFile) == nil {
		return configFile, nil
	}

	return "", errors.New("unable to find config file")
}

func loadConfig(configFile string) (config Settings, err error) {
	if file, err := findConfigFile(configFile); err != nil {
		return config, err
	} else if _, err = toml.DecodeFile(file, &config); err != nil {
		logrus.Errorf("DecodeFile error %v", err)
	}

	return config, err
}

func MustLoadConfig(configFile string) (config Settings) {
	config, _ = loadConfig(configFile)
	ViperToStruct(&config)

	return config
}

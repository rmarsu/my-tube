package config

import (
	"myTube/pkg/log"
	"os"

	"github.com/VandiKond/vanerrors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     string `yaml:"port"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Salt      string `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

func LoadConfig(path string) *Config {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		err = vanerrors.NewWrap("error reading config file", err, vanerrors.EmptyHandler)
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		err = vanerrors.NewWrap("error unmarshaling config file", err, vanerrors.EmptyHandler)
		log.Fatal(err)
	}

	return &config
}

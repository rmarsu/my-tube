package config

import (
	"fmt"
	"myTube/pkg/log"
	"os"

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
     Salt string `yaml:"salt"`
     JWTSecret string `yaml:"jwt_secret"`
}

func LoadConfig(path string) *Config {
	var config Config

     data, err := os.ReadFile(path)
     if err != nil {
          log.Fatal(fmt.Sprintf("Error reading config file: %v", err))
     }

     err = yaml.Unmarshal(data, &config)
     if err != nil {
          log.Fatal(fmt.Sprintf("Error unmarshalling config file: %v", err))
     }

     return &config
}
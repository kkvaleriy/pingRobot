package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type service struct {
	name string `yaml:"name"`
	url  string `yaml:"url"`
}

type server struct {
	port string `yaml:"port"`
}

type Config struct {
	Server   server    `yaml:"server"`
	Services []service `yaml:"services"`
}

func FromFile() *Config {
	file, err := os.ReadFile("../configs/config.yaml")
	if err != nil {
		log.Fatalf("cant read config file: %s", err.Error())
	}
	cfg := &Config{}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		log.Fatalf("cant parse yaml file: %s", err.Error())
	}

	return  cfg
}
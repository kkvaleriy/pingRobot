package config

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var defConfigPath string = "../configs/config.yaml"

type service struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type server struct {
	Port int `yaml:"port"`
}

type Config struct {
	Server   server    `yaml:"server"`
	Services []service `yaml:"services"`
}

func FromFile() *Config {
	cfgFilePath := os.Getenv("PINGROBOT_CONFIG_PATH")
	if len(cfgFilePath) < 1 {
		cfgFilePath = defConfigPath
		log.Printf("env config path dont set, use def path: %s", cfgFilePath)
	}
	file, err := os.ReadFile(cfgFilePath)
	if err != nil {
		log.Fatalf("cant read config file: %s", err.Error())
	}
	cfg := &Config{}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		log.Fatalf("cant parse yaml file: %s", err.Error())
	}

	return cfg
}

func (c *Config) Port() string {
	defPort := "80"

	if c.Server.Port < 10 || c.Server.Port > 65535 {
		log.Printf("invalid port value: %v, will use default port: %s", c.Server.Port, defPort)
		return defPort
	}
	return strconv.Itoa(c.Server.Port)
}

func (c *Config) ServicesForCheck() []service {
	return c.Services
}

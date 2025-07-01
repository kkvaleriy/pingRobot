package config

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type service struct {
	name string `yaml:"name"`
	url  string `yaml:"url"`
}

type server struct {
	Port int `yaml:"port"`
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

func (c *Config) Port() string{ 
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
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

// TODO: Check isDigit
func (c *Config) Port() string{ 
	defPort := "80"
	if len(c.Server.port) < 2 {
		return defPort
	}
	pDigit, err := strconv.Atoi(c.Server.port)
	if err != nil || pDigit < 10 || pDigit > 65535 {
		log.Printf("invalid port value: %s, will use default port: %s", c.Server.port, defPort)
		return defPort
	}
	return c.Server.port
}
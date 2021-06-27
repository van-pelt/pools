package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	App      string   `yaml:"app"`
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	AllowDir string   `yaml:"allowDir"`
}

type Server struct {
	Url  string `yaml:"url"`
	Port int    `yaml:"port"`
}

type Database struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DBName string `yaml:"dbname"`
}

func NewConfig() (*Config, error) {
	conf := Config{}
	data, err := ioutil.ReadFile("config/cfg.yaml")
	if err != nil {
		return nil, fmt.Errorf("Config:ReadFile:%w", err)
	}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		return nil, fmt.Errorf("Config:Unmarshal:%w", err)
	}
	return &conf, nil
}

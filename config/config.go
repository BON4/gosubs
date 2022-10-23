package config

import (
	"os"
	"time"

	"github.com/BON4/timedQ/pkg/ttlstore"
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	AppConfig struct {
		Port    string `yaml:"port"`
		LogFile string `yaml:"log-file"`
	} `yaml:"app"`

	Token struct {
		AcessDuration   time.Duration `yaml:"acess_duration"`
		RefreshDuration time.Duration `yaml:"refresh_duration"`
	} `yaml:"token"`

	Auth struct {
		HeaderKey  string `yaml:"header_key"`
		PaylaodKey string `yaml:"payload_key"`
	} `yaml:"auth"`

	DBconn string `yaml:"db_conn"`

	ttlstore.TTLStoreConfig `yaml:"store"`
}

func LoadServerConfig(path string) (ServerConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return ServerConfig{}, err
	}
	defer f.Close()

	var cfg ServerConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return ServerConfig{}, err
	}

	return cfg, nil
}

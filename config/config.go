package config

import (
	"time"

	"github.com/BON4/timedQ/pkg/ttlstore"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port    string `yaml:"port" mapstructure:"PORT"`
	LogFile string `yaml:"log-file" mapstructure:"LOG_PATH"`

	AcessDuration   time.Duration `yaml:"acess_duration" mapstructure:"ACESS_DUR"`
	RefreshDuration time.Duration `yaml:"refresh_duration" mapstructure:"REFRESH_DUR"`

	HeaderKey  string `yaml:"header_key" mapstructure:"HEADER_KEY"`
	PaylaodKey string `yaml:"payload_key" mapstructure:"PAYLOAD_KEY"`

	DBconn string `yaml:"db_source" mapstructure:"DB_SOURCE"`

	Store ttlstore.TTLStoreConfig
}

func LoadServerConfig(path string) (config ServerConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("cfg")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	var storeCfg ttlstore.TTLStoreConfig

	err = viper.Unmarshal(&storeCfg)

	err = viper.Unmarshal(&config)
	config.Store = storeCfg
	return
}

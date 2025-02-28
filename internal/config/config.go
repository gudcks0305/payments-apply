package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Mode string `mapstructure:"mode"`
	} `mapstructure:"server"`
	PortOne struct {
		ImpKey    string `mapstructure:"imp_key"`
		ImpSecret string `mapstructure:"imp_secret"`
		BaseURL   string `mapstructure:"base_url"`
	} `mapstructure:"portone"`
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일을 읽는데 실패했습니다: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("설정을 구조체로 변환하는데 실패했습니다: %v", err)
	}

	return &config
}

package config

import (
	"log"
	"strings"

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
	// 기본 설정 파일 설정
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("PAYMENTS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	bindEnvs()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("설정 파일을 읽는데 실패했습니다: %v, 환경 변수만 사용합니다", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("설정을 구조체로 변환하는데 실패했습니다: %v", err)
	}

	return &config
}

// 특정 환경 변수들을 명시적으로 바인딩
func bindEnvs() {
	viper.BindEnv("server.port", "PAYMENTS_SERVER_PORT")
	viper.BindEnv("server.mode", "PAYMENTS_SERVER_MODE")

	viper.BindEnv("portone.imp_key", "PAYMENTS_PORTONE_IMP_KEY")
	viper.BindEnv("portone.imp_secret", "PAYMENTS_PORTONE_IMP_SECRET")
	viper.BindEnv("portone.base_url", "PAYMENTS_PORTONE_BASE_URL")

	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
}

// 보안을 위해 민감한 정보 마스킹
func maskString(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	// 첫 4자리만 표시
	return s[:4] + "****"
}

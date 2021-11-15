package telegram

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type Config struct {
	ChatId int
	Token  string
	Url    string `mapstructure:"telegram_bot_url"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	log.Println("Parsing environments")
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	if err := viper.BindEnv("chat_id"); err != nil {
		return err
	}
	cfg.Token = viper.GetString("token")
	s, err := strconv.Atoi(viper.GetString("chat_id"))
	if err != nil {
		return err
	}
	cfg.ChatId = s
	return nil
}

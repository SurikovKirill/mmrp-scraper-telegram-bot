package telegram

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	ChatID int
	Token  string
}

func Init() (*Config, error) {
	var cfg Config
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}
	log.Println("Telegram config done")
	return &cfg, nil
}

func parseEnv(cfg *Config) error {
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
	cfg.ChatID = s

	return nil
}

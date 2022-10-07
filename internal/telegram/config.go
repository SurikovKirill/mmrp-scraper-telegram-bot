package telegram

import (
	"log"
	"os"
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
	// For testing
	os.Setenv("TOKEN", "2102541865:AAFdbDr_mclUw_aCLXadmn1aO9T6sLR3WcQ")
	os.Setenv("CHAT_ID", "-1001559971169")

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

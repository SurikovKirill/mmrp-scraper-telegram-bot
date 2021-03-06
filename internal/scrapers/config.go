package scrapers

import "github.com/spf13/viper"

//Config ...
type Config struct {
	MapmUrl string `mapstructure:"mapm_url"`
	MmrpUrl string `mapstructure:"mmrp_url"`
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

	return &cfg, nil
}

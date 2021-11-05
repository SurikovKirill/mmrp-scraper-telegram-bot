package telegram

type Config struct {
	chatId int
	token  string
	url    string
}

func NewConfig() *Config {
	return &Config{}
}

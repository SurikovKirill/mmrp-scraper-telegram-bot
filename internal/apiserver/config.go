package apiserver

type Config struct {
	Addr string
}

func Init() *Config {
	return &Config{
		Addr: ":8080",
	}
}

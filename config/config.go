package config

type Config struct {
	PublicApiKey  string
	PrivateApiKey string
}

func GetConfig() Config {
	return Config{
		PublicApiKey:  "",
		PrivateApiKey: "",
	}
}

package apiserver

type Config struct {
	BindAddr    string `json:"bind_addr"`
	LogLevel    string `json:"log_level"`
	DatabaseUrl string `json:"database_url"`
}

//func NewConfig() *Config {
//	return &Config{
//		BindAddr: ":8080",
//		LogLevel: "debug",
//		Store:    store.NewConfig(),
//	}}

package main

import "github.com/spf13/viper"

const (
	cfg_serveraddr = "SERVER_ADDR"
	cfg_serverport = "SERVER_PORT"

	default_serveraddr = "127.0.0.1"
	default_serverport = 8080
)

type config struct {
	serverAddr string
	serverPort int
}

func readConfig() config {
	viper.SetDefault(cfg_serveraddr, default_serveraddr)
	viper.SetDefault(cfg_serverport, default_serverport)

	viper.AutomaticEnv()

	return config{
		serverAddr: viper.GetString(cfg_serveraddr),
		serverPort: viper.GetInt(cfg_serverport),
	}
}

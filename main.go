package main

import (
	"pokemontrainer/configs"
	"pokemontrainer/routers"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	configs.InitDB()
	routers.InitRoute()
}

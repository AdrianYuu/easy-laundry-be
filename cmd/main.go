package main

import (
	"fmt"
	"github.com/AdrianYuu/easy-laundry-be/internal/config"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber()

	config.Bootstrap(&config.BootstrapConfig{
		Config: viper,
		App:    app,
	})

	port := viper.GetInt("app.port")

	err := app.Listen(fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}
}

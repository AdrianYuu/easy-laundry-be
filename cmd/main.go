package main

import (
	"fmt"
	"github.com/AdrianYuu/easy-laundry-be/internal/config"
)

func main() {
	viper := config.NewViper()
	fiber := config.NewFiber()
	log := config.NewLogrus(viper)
	database := config.NewGorm(viper, log)

	config.Bootstrap(&config.BootstrapConfig{
		Viper: viper,
		Fiber: fiber,
		Log:   log,
		Db:    database,
	})

	port := viper.GetInt("app.port")

	err := fiber.Listen(fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

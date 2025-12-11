package main

import (
	"fmt"
	"realtime_chat_server/db"
	"realtime_chat_server/internal/handler"
	"realtime_chat_server/internal/repository"
	"realtime_chat_server/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	db := db.NewDatabase()

	userRepository := repository.NewUserRepositoryDB(db)

	userServer := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(userServer)

	app := fiber.New()

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

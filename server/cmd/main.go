package main

import (
	"fmt"
	"realtime_chat_server/db"
	"realtime_chat_server/internal/handler"
	"realtime_chat_server/internal/repository"
	"realtime_chat_server/internal/service"
	"realtime_chat_server/internal/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	db := db.NewDatabase()

	userRepository := repository.NewUserRepositoryDB(db)
	userServer := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userServer)

	hub := websocket.NewHub()
	wsHandler := handler.NewRoomHandler(hub)

	app := fiber.New()

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	app.Get("/logout", userHandler.Logout)

	app.Post("/ws/createRoom", wsHandler.CreateRoom)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

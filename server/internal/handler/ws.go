package handler

import (
	ws "realtime_chat_server/internal/websocket"

	"github.com/gofiber/fiber/v2"
)

type roomHandler struct {
	hub *ws.Hub
}

func NewRoomHandler(h *ws.Hub) *roomHandler {
	return &roomHandler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *roomHandler) CreateRoom(c *fiber.Ctx) error {
	body := new(CreateRoomReq)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})

	}

	h.hub.Rooms[body.ID] = &ws.Room{
		ID:      body.ID,
		Name:    body.Name,
		Clients: make(map[string]*ws.Client),
	}

	return handleSuccess(c, body)
}

package websocket

import "github.com/gofiber/fiber/v2"

type roomHandler struct {
	hub *Hub
}

func NewRoomHandler(h *Hub) *roomHandler {
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

	h.hub.Rooms[body.ID] = &Room{
		ID:      body.ID,
		Name:    body.Name,
		Clients: make(map[string]*Client),
	}
	return nil

	// return handleSuccess(c, body)
}

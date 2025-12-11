package handler

import (
	ws "realtime_chat_server/internal/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type wsHandler struct {
	hub *ws.Hub
}

func NewWSHandler(h *ws.Hub) *wsHandler {
	return &wsHandler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *wsHandler) CreateRoom(c *fiber.Ctx) error {
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

func (h *wsHandler) JoinRoom(conn *websocket.Conn) {
	defer conn.Close()

	roomID := conn.Params("roomId")
	clientID := conn.Query("userId")
	username := conn.Query("username")

	cl := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	// register new client throught register channel
	h.hub.Register <- cl
	// bradcast msg
	h.hub.Broadcast <- m

	// writeMsg
	go cl.WriteMsg()
	// readMsg
	cl.ReadMsg(h.hub)

}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *wsHandler) GetRoom(c *fiber.Ctx) error {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return handleSuccess(c, rooms)

}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *wsHandler) GetClient(c *fiber.Ctx) error {
	clients := make([]ClientRes, 0)

	roomId := c.Params("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		return handleSuccess(c, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	return handleSuccess(c, clients)

}

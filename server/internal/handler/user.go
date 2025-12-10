package handler

import (
	m "realtime_chat_server/internal/model"
	"realtime_chat_server/internal/service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) Register(c *fiber.Ctx) error {
	body := new(m.RegisterReq)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	user, err := h.userSrv.Register(c.Context(), body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return handleSuccess(c, user)
}

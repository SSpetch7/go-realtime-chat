package handler

import (
	m "realtime_chat_server/internal/model"
	"realtime_chat_server/internal/service"
	"time"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Register fail, username or email is already exists."})
	}

	return handleSuccess(c, user)
}

func (h userHandler) Login(c *fiber.Ctx) error {

	body := new(m.LoginReq)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	res, err := h.userSrv.Login(c.Context(), body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Login fail"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		Expires:  time.Now().Add(8 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
	})

	return handleSuccess(c, &m.LoginRes{Username: res.Username, ID: res.ID})

}

func (h userHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	return handleSuccess(c, nil)
}

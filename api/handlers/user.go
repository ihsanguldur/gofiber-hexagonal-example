package handlers

import (
	"github.com/gofiber/fiber/v2"
	"weatherApp/api/utils"
	"weatherApp/domain"
)

type handler struct {
	userService domain.Service
}

func NewHandler(userService domain.Service) *handler {
	return &handler{
		userService: userService,
	}
}

func (h *handler) Create(c *fiber.Ctx) error {
	var err error
	user := new(domain.User)

	if err = c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	if err = h.userService.Create(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Creating.")
	}

	return utils.SuccessPresenter(c, "User Created.", user)
}

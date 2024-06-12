package handler

import (
	"bocchikitsunei_webportfolio/internal/service"
	"github.com/gofiber/fiber/v2"
)

type EmailHandler struct {
	emailService service.EmailService
}

func NewEmailHandler(emailService service.EmailService) *EmailHandler {
	return &EmailHandler{emailService: emailService}
}

func (h *EmailHandler) SendMail(c *fiber.Ctx) error {
	type MailRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
		MSG   string `json:"msg"`
	}

	var request MailRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if request.Name == "" || request.Email == "" || request.Phone == "" || request.MSG == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing required fields"})
	}

	err := h.emailService.SendMail(request.Name, request.Email, request.Phone, request.MSG)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot send email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "email sent successfully"})
}


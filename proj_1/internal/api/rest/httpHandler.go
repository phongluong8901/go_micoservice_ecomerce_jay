package rest

import (
	"proj_1/configs"
	"proj_1/internal/helper"
	"proj_1/pkg/payment"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App    *fiber.App
	DB     *gorm.DB
	Auth   helper.Auth
	Config configs.AppConfig
	Pc     payment.PaymentClient
}

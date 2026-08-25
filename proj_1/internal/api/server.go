package api

import (
	"net/http"
	"proj_1/configs"
	"proj_1/internal/api/rest"
	"proj_1/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	resthandler := &rest.RestHandler{ //Khởi tạo struct mới và trả về địa chỉ ô nhớ của nó (con trỏ).
		App: app, //đưa instance của Fiber vào trường App trong struct
	}

	setupRoutes(resthandler)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I am breathing!",
	})
}

func setupRoutes(rh *rest.RestHandler) { //Khai báo tham số nhận vào là một con trỏ trỏ đến kiểu RestHandler.
	//user handler
	handlers.SetupUserRoutes(rh)

	//transactions

	//catalog
}

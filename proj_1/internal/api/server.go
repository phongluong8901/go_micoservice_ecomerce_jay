package api

import (
	"log"
	"net/http"
	"proj_1/configs"
	"proj_1/internal/api/rest"
	"proj_1/internal/api/rest/handlers"
	"proj_1/internal/domain"
	"proj_1/internal/helper"

	"gorm.io/driver/postgres"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	log.Printf("config DSN %v", config.Dsn)

	app.Get("/health", HealthCheck)

	//connection DB
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v")
	}

	//check conenction
	log.Println("database connected")
	log.Print(db)

	//run migration
	db.AutoMigrate(&domain.User{})

	//auth helper
	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{ //Khởi tạo struct mới và trả về địa chỉ ô nhớ của nó (con trỏ).
		App:  app,  //đưa instance của Fiber vào trường App trong struct
		DB:   db,   //instance cua db postgres
		Auth: auth, //instance authe helper
	}

	setupRoutes(rh)

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

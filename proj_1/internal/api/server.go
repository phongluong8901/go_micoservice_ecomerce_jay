package api

import (
	"log"
	"net/http"
	"proj_1/configs"
	"proj_1/internal/api/rest"
	"proj_1/internal/api/rest/handlers"
	"proj_1/internal/domain"
	"proj_1/internal/helper"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	log.Printf("config DSN %v", config.Dsn)

	//connection DB
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v")
	}

	//check conenction
	log.Println("database connected")
	log.Print(db)

	//run migration - create tables
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Address{},
		&domain.BankAccount{},
		&domain.Category{},
		&domain.Product{},
		&domain.Cart{},
	)
	if err != nil {
		log.Fatalf("error on runing migration %v", err.Error())
	}

	log.Println("migration was successful")

	//cors configuration
	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:3030",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})

	app.Use(c)

	app.Get("/", func(c *fiber.Ctx) error {
		return rest.SuccessResponse(c, "I am Healthy", &fiber.Map{
			"status": "ok with 200 status code",
		})
	})

	//auth helper
	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{ //Khởi tạo struct mới và trả về địa chỉ ô nhớ của nó (con trỏ).
		App:    app,  //đưa instance của Fiber vào trường App trong struct
		DB:     db,   //instance cua db postgres
		Auth:   auth, //instance authe helper
		Config: config,
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
	handlers.SetupCatalogRoutes(rh)
}

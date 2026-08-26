package handlers

import (
	"log"
	"proj_1/internal/api/rest"
	"proj_1/internal/repository"
	"proj_1/service"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	//svc UserService
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	//create an instance of user service & inject to handler
	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := CatalogHandler{
		svc: svc,
	}

	//public
	//listing products and categories
	app.Get("/products")
	app.Get("/products/:id")
	app.Get("/categories")
	app.Get("/categories/:id")

	//private
	//manage products and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	//Categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	//Products
	selRoutes.Post("/products", handler.CreateProducts)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Patch("/products/:id", handler.EditProducts)
	selRoutes.Patch("/products/:id", handler.UpdateStock) //update stock
	selRoutes.Delete("/products/:id", handler.DeleteProducts)
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	log.Printf("current user %v", user.ID)

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) CreateProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) EditProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) DeleteProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

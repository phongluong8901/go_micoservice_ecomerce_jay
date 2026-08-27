package handlers

import (
	"errors"
	"log"
	"net/http"
	"proj_1/internal/api/rest"
	"proj_1/internal/dto"
	"proj_1/internal/repository"
	"proj_1/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	//svc UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	//create an instance of user service & inject to handler
	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		CRepo:  repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := UserHandler{
		svc: svc,
	}

	//- create prefix for public gorup
	pubRoutes := app.Group("/users")
	//public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	//- create prefix for private gorup
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	//private endpoints
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Get("/profile", handler.GetProfile)
	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Patch("/profile", handler.UpdateProfile)
	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/orders", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)
	pvtRoutes.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	//to create user
	user := dto.UserSignup{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})

}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "please provide correct user id and password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	//create vertification code and updte to user profile in DB
	err := h.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate vertification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	//request
	var req dto.VertificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		log.Printf("%v", err)
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	req := dto.ProfileInput{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a  valid input",
		})
	}

	log.Printf("User %v", user)

	//create profile
	err := h.svc.CreateProfile(user.ID, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to create profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)

	//call user service and perform get profile
	profile, err := h.svc.GetProfile(user.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to get profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"user":    profile,
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	req := dto.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a  valid input",
		})
	}

	err := h.svc.UpdateProfile(user.ID, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to update profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "update profile",
		"cart":    user,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	req := dto.CreateCartRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid product and qty",
		})
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)

	//call user service and  perform create cart
	cartItems, err := h.svc.CreateCart(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "cart created successfully", cartItems)

}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	cart, err := h.svc.FindCart(user.ID)
	if err != nil {
		return rest.InternalError(ctx, errors.New("cart doese not exist"))
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "GetCart",
		"cart":    cart,
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "CreateOrder",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "GetOrders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "GetOrder",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "request parameters are not valid",
		})
	}

	token, err := h.svc.BecomeSeller(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "fail to become seller",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
		"token":   token,
	})
}

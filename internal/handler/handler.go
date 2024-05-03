package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ngobrut/cat-tinder-api/config"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/middleware"
	"github.com/ngobrut/cat-tinder-api/internal/usecase"
)

type Handler struct {
	uc usecase.IFaceUsecase
}

func InitHTTPHandler(cnf *config.Config, uc usecase.IFaceUsecase) *fiber.App {
	h := Handler{
		uc: uc,
	}

	m := middleware.Middleware{
		JWTSecret: cnf.JWTSecret,
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} | ${method} | ${url}\n",
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(response.JsonResponse{
			Message: "hello!",
			Data: fiber.Map{
				"app-name": "cat-tinder-api",
			},
		})
	})

	api := app.Group("/v1")

	user := api.Group("/user")
	user.Post("/register", h.Register)
	user.Post("/login", h.Login)

	profile := user.Group("/profile", m.Authorize())
	profile.Get("/", h.GetProfile)

	manageCat := api.Group("/cat", m.Authorize())
	manageCat.Post("/", h.CreateCat)
	manageCat.Get("/", h.GetListCat)
	manageCat.Put("/:id", h.UpdateCat)
	manageCat.Delete("/:id", h.DeleteCat)

	catMatch := api.Group("/cat/match", m.Authorize())
	catMatch.Post("/", h.CreateCatMatch)
	catMatch.Get("/", h.GetListCatMatch)
	catMatch.Post("/approve", h.ApproveCatMatch)
	catMatch.Post("/reject", h.RejectCatMatch)
	catMatch.Delete("/:id", h.DeleteCatMatch)

	return app
}

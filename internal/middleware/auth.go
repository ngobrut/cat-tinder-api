package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	custom_jwt "github.com/ngobrut/cat-tinder-api/pkg/jwt"
)

type Middleware struct {
	JWTSecret string
}

func (m Middleware) Authorize() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(http.StatusUnauthorized).JSON(response.JsonResponse{
				Message: "error:unauthorized",
				Error: &response.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.ErrorMessageMap[http.StatusUnauthorized],
				},
			})
		}

		parsed, err := jwt.ParseWithClaims(token[7:], &custom_jwt.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JWTSecret), nil
		})

		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.JsonResponse{
				Message: "error:unauthorized",
				Error: &response.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.ErrorMessageMap[http.StatusUnauthorized],
				},
			})
		}

		claims, ok := parsed.Claims.(*custom_jwt.CustomClaims)
		if !ok && !parsed.Valid {
			return c.Status(http.StatusUnauthorized).JSON(response.JsonResponse{
				Message: "error:unauthorized",
				Error: &response.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.ErrorMessageMap[http.StatusUnauthorized],
				},
			})
		}

		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_validator"
	"github.com/ngobrut/cat-tinder-api/pkg/util"
)

func (h *Handler) Register(c *fiber.Ctx) error {
	var req request.Register
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	res, err := h.uc.Register(&req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, res, http.StatusCreated, "User registered successfully")
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req request.Login
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	res, err := h.uc.Login(&req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, res, http.StatusOK, "User logged-in successfully")
}

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	userID, err := uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	res, err := h.uc.GetProfile(userID)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, res, http.StatusOK, "Get profile success")
}

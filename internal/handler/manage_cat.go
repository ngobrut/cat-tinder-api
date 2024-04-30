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

func (h *Handler) CreateCat(c *fiber.Ctx) error {
	userID, err := uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	var req request.CreateCat
	err = custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	res, err := h.uc.CreateCat(&req, userID)
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, res, http.StatusCreated, "successfully add cat")
}

func (h *Handler) UpdateCat(c *fiber.Ctx) error {
	// to do
	return nil
}

func (h *Handler) GetCats(c *fiber.Ctx) error {
	// to do
	return nil
}

func (h *Handler) DeleteCat(c *fiber.Ctx) error {
	// to do
	return nil
}

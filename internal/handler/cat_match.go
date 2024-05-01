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

func (h *Handler) CreateCatMatch(c *fiber.Ctx) error {
	var req request.CreateCatMatch
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	req.UserID, err = uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	err = h.uc.CreateCatMatch(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusCreated, "Success")
}

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"

	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_validator"
	"github.com/ngobrut/cat-tinder-api/pkg/util"
)

func (h *Handler) CreateCat(c *fiber.Ctx) error {
	var req request.CreateCat
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	req.UserID, err = uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	res, err := h.uc.CreateCat(&req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, res, http.StatusCreated, "Cat added successfully")
}

func (h *Handler) GetListCat(c *fiber.Ctx) error {
	var params request.ListCatQuery
	err := c.QueryParser(&params)
	if err != nil {
		return response.Error(c, err)
	}

	if params.Owned != "" {
		params.UserID, err = uuid.Parse(util.GetUserIDFromHeader(c))
		if err != nil {
			return response.Error(c, err)
		}
	}

	res, err := h.uc.GetListCat(&params)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, res, http.StatusOK, "Success")
}

func (h *Handler) UpdateCat(c *fiber.Ctx) error {
	var req request.UpdateCat
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	req.CatID, err = uuid.Parse(c.Params("id"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return response.Error(c, err)
	}

	req.UserID, err = uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	err = h.uc.UpdateCat(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusOK, "Cat updated successfully")
}

func (h *Handler) DeleteCat(c *fiber.Ctx) error {
	catID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return response.Error(c, err)
	}

	err = h.uc.DeleteCat(c, catID)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusOK, "Cat deleted successfully")
}

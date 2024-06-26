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

func (h *Handler) CreateCatMatch(c *fiber.Ctx) error {
	var req request.CreateCatMatch
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	err = uuid.Validate(req.MatchCatID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return response.Error(c, err)
	}

	err = uuid.Validate(req.UserCatID)
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

	err = h.uc.CreateCatMatch(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusCreated, "Cat match request created successfully")
}

func (h *Handler) GetListCatMatch(c *fiber.Ctx) error {
	userID, err := uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	params := &request.ListCatMatchQuery{
		UserID: userID,
	}

	res, err := h.uc.GetListCatMatch(params)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, res, http.StatusOK, "Success")
}

func (h *Handler) ApproveCatMatch(c *fiber.Ctx) error {
	var req request.ApproveCatMatch
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	err = uuid.Validate(req.MatchID.String())
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return response.Error(c, err)
	}

	req.UserId, err = uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return response.Error(c, err)
	}

	err = h.uc.ApproveCatMatch(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusOK, "successfully matches the cat match request")
}

func (h *Handler) RejectCatMatch(c *fiber.Ctx) error {
	var req request.RejectCatMatch
	err := custom_validator.ValidateStruct(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	err = uuid.Validate(req.MatchID.String())
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

	err = h.uc.RejectCatMatch(c, &req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusOK, "successfully reject the cat match request")
}

func (h *Handler) DeleteCatMatch(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return response.Error(c, err)
	}

	err = h.uc.DeleteCatMatch(c, ID)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, nil, http.StatusOK, "Cat match request deleted successfully")
}

package usecase

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
	"github.com/ngobrut/cat-tinder-api/pkg/util"
)

// CreateCatMatch implements IFaceUsecase.
func (u *Usecase) CreateCatMatch(c *fiber.Ctx, req *request.CreateCatMatch) error {
	issuerCat, err := u.repo.FindOneCatByID(req.UserCatID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "id is not found",
		})
		return err
	}

	if issuerCat == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "issuer cat not found",
		})

		return err
	}

	if issuerCat.UserID != req.UserID {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "cannot issue a match request because this isn't your cat",
		})

		return err
	}

	receiverCat, err := u.repo.FindOneCatByID(req.MatchCatID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "id is not found",
		})
		return err
	}

	if receiverCat == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "receiver cat not found",
		})

		return err
	}

	if issuerCat.HasMatched {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "your cat has already matched with other cat",
		})

		return err
	}

	if receiverCat.HasMatched {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "the cat you want to match with has already matched",
		})

		return err
	}

	if issuerCat.UserID == receiverCat.UserID {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "cannot match cats from the same owner",
		})

		return err
	}

	if issuerCat.Sex == receiverCat.Sex {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "both of the cats cannot be of the same sex",
		})

		return err
	}

	err = u.repo.CheckDuplicateMatchRequest(issuerCat.CatID, receiverCat.CatID)
	if err != nil {
		return err
	}

	data := &model.CatMatch{
		ID:             uuid.New(),
		IssuerUserID:   issuerCat.UserID,
		IssuerCatID:    issuerCat.CatID,
		ReceiverUserID: receiverCat.UserID,
		ReceiverCatID:  receiverCat.CatID,
		Message:        req.Message,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = u.repo.CreateCatMatch(data)
	if err != nil {
		return err
	}

	return nil
}

// GetListCatMatch implements IFaceUsecase.
func (u *Usecase) GetListCatMatch(params *request.ListCatMatchQuery) ([]*response.CatMatchResponse, error) {
	var res = make([]*response.CatMatchResponse, 0)

	res, err := u.repo.FindCatMatch(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ApproveCatMatch implements IFaceUsecase.
func (u *Usecase) ApproveCatMatch(c *fiber.Ctx, req *request.ApproveCatMatch) error {
	catMatch, err := u.repo.FindOneCatMatchByID(req.MatchID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return err
	}

	if req.UserId != catMatch.ReceiverUserID {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return err
	}

	if catMatch.IsApproved != nil || catMatch.DeletedAt != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "matchId is no longer valid",
		})

		return err
	}

	err = u.repo.ApproveCatMatch(req.MatchID)
	if err != nil {
		return err
	}

	return nil
}

// RejectCatMatch implements IFaceUsecase.
func (u *Usecase) RejectCatMatch(c *fiber.Ctx, req *request.RejectCatMatch) error {
	cm, err := u.repo.FindOneCatMatchByID(req.MatchID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return err
	}

	if req.UserID != cm.ReceiverUserID {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusUnauthorized,
			Message:  "matchId is not found",
		})

		return err
	}

	if cm.IsApproved != nil || cm.DeletedAt != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "matchId is no longer valid",
		})

		return err
	}

	err = u.repo.RejectCatMatch(req.MatchID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCatMatch implements IFaceUsecase.
func (u *Usecase) DeleteCatMatch(c *fiber.Ctx, ID uuid.UUID) error {
	cm, err := u.repo.FindOneCatMatchByID(ID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})
		return err
	}

	if cm == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return err
	}

	userID, err := uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return err
	}

	if cm.IssuerUserID != userID {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusUnauthorized,
			Message:  "you are not the issuer of this match request",
		})

		return err
	}

	if cm.IsApproved != nil {
		if *cm.IsApproved {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusBadRequest,
				Message:  "this match request was approved",
			})
		} else {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusBadRequest,
				Message:  "this match request was rejected",
			})
		}

		return err
	}

	return u.repo.DeleteCatMatchByID(ID)
}

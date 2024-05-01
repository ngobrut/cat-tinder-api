package usecase

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/internal/repository"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

// CreateCatMatch implements IFaceUsecase.
func (u *Usecase) CreateCatMatch(c *fiber.Ctx, req *request.CreateCatMatch) error {
	issuerCat, err := u.repo.FindOneCatByID(req.UserCatID)
	if err != nil && !repository.IsRecordNotFound(err) {
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
			HTTPCode: http.StatusUnauthorized,
			Message:  "cannot issue a match request because this isn't your cat",
		})

		return err
	}

	receiverCat, err := u.repo.FindOneCatByID(req.MatchCatID)
	if err != nil {
		return err
	}

	err = u.repo.CheckDuplicateMatchRequest(issuerCat.CatID, receiverCat.CatID)
	if err != nil {
		return err
	}

	if receiverCat == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "issuer cat not found",
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

	return res, nil
}

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
	if err != nil && !repository.IsRecordNotFound(err) {
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

	list, err := u.repo.FindCatMatch(params)
	if err != nil {
		return res, err
	}

	for _, v := range list {
		res = append(res, &response.CatMatchResponse{
			ID: v.ID,
			IssuedBy: &response.IssuedBy{
				Name:      v.Issuer.Name,
				Email:     v.Issuer.Email,
				CreatedAt: v.Issuer.CreatedAt,
			},
			MatchCatDetail: &response.CatResponse{
				CatID:       v.ReceiverCat.CatID,
				Name:        v.ReceiverCat.Name,
				Race:        v.ReceiverCat.Race,
				Sex:         v.ReceiverCat.Sex,
				AgeInMonth:  v.ReceiverCat.AgeInMonth,
				ImageUrls:   v.ReceiverCat.ImageURLs,
				Description: v.ReceiverCat.Description,
				HasMatched:  v.ReceiverCat.HasMatched,
				CreatedAt:   v.ReceiverCat.CreatedAt,
			},
			UserCatDetail: &response.CatResponse{
				CatID:       v.IssuerCat.CatID,
				Name:        v.IssuerCat.Name,
				Race:        v.IssuerCat.Race,
				Sex:         v.IssuerCat.Sex,
				AgeInMonth:  v.IssuerCat.AgeInMonth,
				ImageUrls:   v.IssuerCat.ImageURLs,
				Description: v.IssuerCat.Description,
				HasMatched:  v.IssuerCat.HasMatched,
				CreatedAt:   v.IssuerCat.CreatedAt,
			},
			Message:   v.Message,
			CreatedAt: v.CreatedAt,
		})
	}

	return res, nil
}

func (u *Usecase) ApproveCatMatch(c *fiber.Ctx, req *request.ApproveCatMatch) error {
	catMatch, err := u.repo.FindOneCatMatchByMatchID(req.MatchID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})

		return err
	}

	if req.UserId == catMatch.ReceiverUserID {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "matchId is not found",
		})
	}

	if catMatch.IsApproved != nil && catMatch.DeletedAt != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "matchId is no longer valid",
		})
	}
	err = u.repo.ApproveCatMatch(req.MatchID)
	if err != nil {
		return err
	}
	return nil
}

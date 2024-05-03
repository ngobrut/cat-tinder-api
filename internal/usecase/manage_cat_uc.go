package usecase

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/internal/repository"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
	"github.com/ngobrut/cat-tinder-api/pkg/util"
)

// CreateCat implements IFaceUsecase.
func (u *Usecase) CreateCat(req *request.CreateCat) (*response.CreateCat, error) {
	cat := &model.Cat{
		CatID:       uuid.New(),
		UserID:      req.UserID,
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
		ImageURLs:   req.ImageURLs,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := u.repo.CreateCat(cat)
	if err != nil {
		return nil, err
	}

	res := &response.CreateCat{
		CatID:     cat.CatID,
		CreatedAt: cat.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

// GetListCat implements IFaceUsecase.
func (u *Usecase) GetListCat(params *request.ListCatQuery) ([]*response.CatResponse, error) {
	var res = make([]*response.CatResponse, 0)

	list, err := u.repo.FindCat(params)
	if err != nil {
		return res, err
	}

	for _, v := range list {
		res = append(res, &response.CatResponse{
			CatID:       v.CatID,
			Name:        v.Name,
			Race:        v.Race,
			Sex:         v.Sex,
			AgeInMonth:  v.AgeInMonth,
			ImageUrls:   v.ImageURLs,
			Description: v.Description,
			HasMatched:  v.HasMatched,
			CreatedAt:   v.CreatedAt,
		})
	}

	return res, nil
}

// UpdateCat implements IFaceUsecase.
func (u *Usecase) UpdateCat(c *fiber.Ctx, req *request.UpdateCat) error {
	cat, err := u.repo.FindOneCatByID(req.CatID.String())
	if err != nil && !repository.IsRecordNotFound(err) {
		return err
	}

	if cat == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "cat not found",
		})

		return err
	}

	err = u.repo.CheckOwnCat(req.UserID, req.CatID)
	if err != nil {
		return err
	}

	cm, err := u.repo.FindOneCatMatchByCatID(req.CatID)
	if err != nil && !repository.IsRecordNotFound(err) {
		return err
	}

	if cm != nil && req.Sex != cat.Sex {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "cannot change cat's sex because your cat's match request has been issued",
		})

		return err
	}

	data := map[string]interface{}{
		"name":         req.Name,
		"race":         req.Race,
		"sex":          req.Sex,
		"age_in_month": req.AgeInMonth,
		"description":  req.Description,
		"image_urls":   pq.Array(req.ImageURLs),
		"updated_at":   time.Now(),
	}

	err = u.repo.UpdateCatByID(data, req.CatID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCat implements IFaceUsecase.
func (u *Usecase) DeleteCat(c *fiber.Ctx, catID uuid.UUID) error {
	cat, err := u.repo.FindOneCatByID(catID.String())
	if err != nil && !repository.IsRecordNotFound(err) {
		return err
	}

	if cat == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "cat not found",
		})

		return err
	}

	userID, err := uuid.Parse(util.GetUserIDFromHeader(c))
	if err != nil {
		return err
	}

	err = u.repo.CheckOwnCat(userID, catID)
	if err != nil {
		return err
	}

	return u.repo.DeleteCatByID(catID)
}

package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

// CreateCat implements IFaceUsecase.
func (u *Usecase) CreateCat(req *request.CreateCat, userID uuid.UUID) (*response.CreateCat, error) {
	cat := &model.Cat{
		CatID:       uuid.New(),
		UserID:      userID,
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
		ImageUrl:    req.ImageUrls,
	}

	err := u.repo.CreateCat(cat)
	if err != nil {
		return nil, err
	}
	fmt.Println(cat)
	res := &response.CreateCat{
		CatID:     cat.CatID,
		CreatedAt: cat.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

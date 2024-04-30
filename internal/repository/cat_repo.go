package repository

import (
	"net/http"

	"github.com/lib/pq"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

// CreateCat implements IFaceRepository.
func (r *Repository) CreateCat(data *model.Cat) error {
	args := []interface{}{
		data.CatID,
		data.UserID,
		data.Name,
		data.Race,
		data.Sex,
		data.AgeInMonth,
		data.Description,
		pq.Array(data.ImageUrl),
	}
	err := r.db.
		QueryRow("INSERT INTO cats(cat_id, user_id, name, race, sex, age_in_month, description, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning has_matched, created_at, updated_at", args...).
		Scan(&data.HasMatched, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.ErrorMessageMap[http.StatusInternalServerError],
		})
	}
	return nil
}

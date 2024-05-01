package repository

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

// FindCatMatch implements IFaceRepository.
func (r *Repository) FindCatMatch(params *request.ListCatMatchQuery) ([]*model.CatMatch, error) {
	var res = make([]*model.CatMatch, 0)

	return res, nil
}

// CreateCatMatch implements IFaceRepository.
func (r *Repository) CreateCatMatch(data *model.CatMatch) error {
	args := []interface{}{
		data.ID,
		data.IssuerUserID,
		data.IssuerCatID,
		data.ReceiverUserID,
		data.ReceiverCatID,
		data.Message,
		data.CreatedAt.Format("2006-01-02 15:04:05"),
		data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	_, err := r.db.Exec("INSERT INTO cat_matches(id, issuer_user_id, issuer_cat_id, receiver_user_id, receiver_cat_id, message, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", args...)
	if err != nil {
		return err
	}

	return nil
}

// CheckDuplicateMatchRequest implements IFaceRepository.
func (r *Repository) CheckDuplicateMatchRequest(issuerCatID uuid.UUID, receiverCatID uuid.UUID) error {
	var cnt int

	if err := r.db.
		QueryRow("SELECT COUNT(*) FROM cat_matches WHERE issuer_cat_id = $1 AND receiver_cat_id = $2", issuerCatID, receiverCatID).
		Scan(&cnt); err != nil {
		return err
	}

	if cnt > 0 {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "match request for your cat to the cat you want to match with has already been issued",
		})

		return err
	}

	return nil
}

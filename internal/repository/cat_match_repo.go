package repository

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

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

// FindCatMatch implements IFaceRepository.
func (r *Repository) FindCatMatch(params *request.ListCatMatchQuery) ([]*model.CatMatch, error) {
	var res = make([]*model.CatMatch, 0)

	rows, err := r.db.Query("SELECT * FROM cat_matches WHERE (issuer_user_id = $1 OR receiver_user_id = $2) ORDER BY created_at DESC", params.UserID, params.UserID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cm := &model.CatMatch{}
		if err := rows.Scan(
			&cm.ID,
			&cm.IssuerUserID,
			&cm.IssuerCatID,
			&cm.ReceiverUserID,
			&cm.ReceiverCatID,
			&cm.Message,
			&cm.IsApproved,
			&cm.CreatedAt,
			&cm.UpdatedAt,
			&cm.DeletedAt,
		); err != nil {
			return nil, err
		}

		ic := &model.Cat{}
		if err := r.db.
			QueryRow("SELECT * FROM cats WHERE cat_id = $1", cm.IssuerCatID).
			Scan(
				&ic.CatID,
				&ic.UserID,
				&ic.Name,
				&ic.Race,
				&ic.Sex,
				&ic.AgeInMonth,
				&ic.Description,
				&ic.HasMatched,
				pq.Array(&ic.ImageURLs),
				&ic.CreatedAt,
				&ic.UpdatedAt,
				&ic.DeletedAt,
			); err != nil {
			return nil, err
		}

		rc := &model.Cat{}
		if err := r.db.
			QueryRow("SELECT * FROM cats WHERE cat_id = $1", cm.ReceiverCatID).
			Scan(
				&rc.CatID,
				&rc.UserID,
				&rc.Name,
				&rc.Race,
				&rc.Sex,
				&rc.AgeInMonth,
				&rc.Description,
				&rc.HasMatched,
				pq.Array(&rc.ImageURLs),
				&rc.CreatedAt,
				&rc.UpdatedAt,
				&rc.DeletedAt,
			); err != nil {
			return nil, err
		}

		issuer := &model.User{}
		if err := r.db.
			QueryRow("SELECT * FROM users WHERE user_id = $1", cm.IssuerUserID).
			Scan(
				&issuer.UserID,
				&issuer.Name,
				&issuer.Email,
				&issuer.Password,
				&issuer.CreatedAt,
				&issuer.UpdatedAt,
				&issuer.DeletedAt,
			); err != nil {
			return nil, err
		}

		cm.IssuerCat = ic
		cm.ReceiverCat = rc
		cm.Issuer = issuer

		res = append(res, cm)
	}

	return res, nil
}

// FindOneCatMatchByCatID implements IFaceRepository.
func (r *Repository) FindOneCatMatchByCatID(catID uuid.UUID) (*model.CatMatch, error) {
	res := &model.CatMatch{}

	if err := r.db.
		QueryRow("SELECT * FROM cat_matches WHERE (issuer_cat_id = $1 OR receiver_cat_id = $2)", catID, catID).
		Scan(
			&res.ID,
			&res.IssuerUserID,
			&res.IssuerCatID,
			&res.ReceiverUserID,
			&res.ReceiverCatID,
			&res.Message,
			&res.IsApproved,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
		return nil, err
	}

	return res, nil
}

// FindOneCatMatchByMatchID implements IFaceRepository.
func (r *Repository) FindOneCatMatchByMatchID(matchID uuid.UUID) (*model.CatMatch, error) {
	res := &model.CatMatch{}

	if err := r.db.
		QueryRow("SELECT * FROM cat_matches WHERE match_id = $1", matchID).
		Scan(
			&res.ID,
			&res.IssuerUserID,
			&res.IssuerCatID,
			&res.ReceiverUserID,
			&res.ReceiverCatID,
			&res.Message,
			&res.IsApproved,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repository) ApproveCatMatch(matchID uuid.UUID) error {
	var issuerCatID uuid.UUID
	var receiverCatID uuid.UUID
	err := r.db.
		QueryRow("UPDATE cat_matches SET is_approved = true WHERE match_id = $1 returning issuer_cat_id, receiver_cat_id", matchID).
		Scan(
			&issuerCatID,
			&receiverCatID,
		)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("UPDATE cat_matches SET deleted_at = $1 WHERE issuer_cat_id in ($2, $3) or receiver_cat_id in ($2, $3) ", time.Now().Format("2006-01-02 15:04:05"), issuerCatID, receiverCatID)
	if err != nil {
		return err
	}

	return nil
}

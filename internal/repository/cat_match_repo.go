package repository

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
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
		QueryRow("SELECT COUNT(*) FROM cat_matches WHERE issuer_cat_id = $1 AND receiver_cat_id = $2 AND deleted_at iS NULL", issuerCatID, receiverCatID).
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

// FindOneCatMatchByCatID implements IFaceRepository.
func (r *Repository) FindOneCatMatchByCatID(catID uuid.UUID) (*model.CatMatch, error) {
	res := &model.CatMatch{}

	if err := r.db.
		QueryRow("SELECT * FROM cat_matches WHERE (issuer_cat_id = $1 OR receiver_cat_id = $2) AND deleted_at IS NULL", catID, catID).
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
			&res.DeletedAt,
		); err != nil {
		return nil, err
	}

	return res, nil
}

// FindCatMatch implements IFaceRepository.
func (r *Repository) FindCatMatch(params *request.ListCatMatchQuery) ([]*response.CatMatchResponse, error) {
	var res = make([]*response.CatMatchResponse, 0)

	query := `
		SELECT cm.id        as cm_id,
		u.name              AS u_name,
		u.email             AS u_email,
		TO_CHAR(U.created_at, 'YYYY-MM-DD HH24:MI:SS')        AS u_created_at,
		c2.cat_id           AS c2_cat_id,
		c2.name             AS c2_name,
		c2.race             AS c2_race,
		c2.sex              AS c2_sex,
		c2.description      AS c2_description,
		c2.age_in_month     AS c2_age_in_month,
		c2.image_urls       AS c2_image_urls,
		c2.has_matched      AS c2_has_matched,
		TO_CHAR(c2.created_at, 'YYYY-MM-DD HH24:MI:SS')       AS c2_created_at,
		c1.cat_id           AS c1_cat_id,
		c1.name             AS c1_name,
		c1.race             AS c1_race,
		c1.sex              AS c1_sex,
		c1.description      AS c1_description,
		c1.age_in_month     AS c1_age_in_month,
		c1.image_urls       AS c1_image_urls,
		c1.has_matched      AS c1_has_matched,
		TO_CHAR(c1.created_at, 'YYYY-MM-DD HH24:MI:SS')       AS c1_created_at,
		cm.message          as cm_message,
		TO_CHAR(cm.created_at, 'YYYY-MM-DD HH24:MI:SS')       as cm_created_at
		FROM cat_matches cm
				LEFT JOIN cats c1 ON c1.cat_id = cm.issuer_cat_id
				LEFT JOIN cats c2 ON c2.cat_id = cm.receiver_cat_id
				LEFT JOIN users u ON u.user_id = cm.issuer_user_id
		WHERE (cm.issuer_user_id = $1 OR cm.receiver_user_id = $2)
		AND cm.deleted_at IS NULL
		ORDER BY cm.created_at DESC
	`

	rows, err := r.db.Query(query, params.UserID, params.UserID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cmr := &response.CatMatchResponse{}

		if err := rows.Scan(
			&cmr.ID,
			&cmr.IssuedBy.Name,
			&cmr.IssuedBy.Email,
			&cmr.IssuedBy.CreatedAt,
			&cmr.MatchCatDetail.CatID,
			&cmr.MatchCatDetail.Name,
			&cmr.MatchCatDetail.Race,
			&cmr.MatchCatDetail.Sex,
			&cmr.MatchCatDetail.Description,
			&cmr.MatchCatDetail.AgeInMonth,
			pq.Array(&cmr.MatchCatDetail.ImageURLs),
			&cmr.MatchCatDetail.HasMatched,
			&cmr.MatchCatDetail.CreatedAt,
			&cmr.UserCatDetail.CatID,
			&cmr.UserCatDetail.Name,
			&cmr.UserCatDetail.Race,
			&cmr.UserCatDetail.Sex,
			&cmr.UserCatDetail.Description,
			&cmr.UserCatDetail.AgeInMonth,
			pq.Array(&cmr.UserCatDetail.ImageURLs),
			&cmr.UserCatDetail.HasMatched,
			&cmr.UserCatDetail.CreatedAt,
			&cmr.Message,
			&cmr.CreatedAt,
		); err != nil {
			return nil, err
		}

		res = append(res, cmr)
	}

	return res, nil
}

// FindOneCatMatchByID implements IFaceRepository.
func (r *Repository) FindOneCatMatchByID(ID uuid.UUID) (*model.CatMatch, error) {
	res := &model.CatMatch{}
	err := r.db.
		QueryRow("SELECT * FROM cat_matches WHERE id = $1", ID).
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
			&res.DeletedAt,
		)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) ApproveCatMatch(matchID uuid.UUID) error {
	var issuerCatID uuid.UUID
	var receiverCatID uuid.UUID
	err := r.db.
		QueryRow("UPDATE cat_matches SET is_approved = true WHERE id = $1 returning issuer_cat_id, receiver_cat_id", matchID).
		Scan(
			&issuerCatID,
			&receiverCatID,
		)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("UPDATE cat_matches SET deleted_at = $1 WHERE (issuer_cat_id in ($2, $3) or receiver_cat_id in ($2, $3)) AND NOT id = $4 ", time.Now().Format("2006-01-02 15:04:05"), issuerCatID, receiverCatID, matchID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("UPDATE cats SET has_matched = true WHERE cat_id in ($1, $2)  ", issuerCatID, receiverCatID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RejectCatMatch(matchID uuid.UUID) error {
	_, err := r.db.Exec("UPDATE cat_matches SET is_approved = false WHERE id = $1 ", matchID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCatMatchByID implements IFaceRepository.
func (r *Repository) UpdateCatMatchByID(data map[string]interface{}, ID uuid.UUID) error {
	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	for k, v := range data {
		clause = append(clause, fmt.Sprintf("%s = $%d", k, counter))
		args = append(args, v)
		counter++
	}

	query := `UPDATE cat_matches SET `

	if len(clause) > 0 {
		query += strings.Join(clause, ", ")
	}

	query += fmt.Sprintf(" WHERE id = $%d", counter)
	args = append(args, ID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCatMatchByID implements IFaceRepository.
func (r *Repository) DeleteCatMatchByID(ID uuid.UUID) error {
	_, err := r.db.Exec("UPDATE cat_matches SET deleted_at = $1 WHERE id = $2", time.Now().Format("2006-01-02 15:04:05"), ID)
	if err != nil {
		return err
	}

	return nil
}

package repository

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
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
		pq.Array(data.ImageURLs),
		data.CreatedAt.Format("2006-01-02 15:04:05"),
		data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	_, err := r.db.Exec("INSERT INTO cats(cat_id, user_id, name, race, sex, age_in_month, description, image_urls, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", args...)
	if err != nil {
		return err
	}

	return nil
}

// FindCat implements IFaceRepository.
func (r *Repository) FindCat(params *request.ListCatQuery) ([]*model.Cat, error) {
	var res = make([]*model.Cat, 0)

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	if params.ID != "" {
		clause = append(clause, fmt.Sprintf(" cat_id = $%d", counter))
		args = append(args, params.ID)
		counter++
	} else if params.ID == "" {
		if params.Race != "" {
			clause = append(clause, fmt.Sprintf(" race = $%d", counter))
			args = append(args, params.Race)
			counter++
		}

		if params.Sex != "" {
			clause = append(clause, fmt.Sprintf(" sex = $%d", counter))
			args = append(args, params.Sex)
			counter++
		}

		if params.HasMatched != "" {
			matched, err := strconv.ParseBool(params.HasMatched)
			if err != nil {
				return nil, err
			}

			clause = append(clause, fmt.Sprintf(" has_matched = $%d", counter))
			args = append(args, matched)
			counter++
		}

		if params.AgeInMonth != "" {
			switch params.AgeInMonth {
			case ">4":
				clause = append(clause, fmt.Sprintf(" age_in_month > $%d", counter))
				args = append(args, 4)
			case "<4":
				clause = append(clause, fmt.Sprintf(" age_in_month < $%d", counter))
				args = append(args, 4)
			default:
				clause = append(clause, fmt.Sprintf(" age_in_month = $%d", counter))
				args = append(args, params.AgeInMonth)
			}

			counter++
		}

		if params.Owned != "" {
			owned, err := strconv.ParseBool(params.Owned)
			if err != nil {
				return nil, err
			}

			if owned {
				clause = append(clause, fmt.Sprintf(" user_id = $%d", counter))
				args = append(args, params.UserID)
			} else {
				clause = append(clause, fmt.Sprintf(" user_id != $%d", counter))
				args = append(args, params.UserID)
			}

			counter++
		}

		if params.Search != "" {
			clause = append(clause, fmt.Sprintf(" name LIKE $%d", counter))
			args = append(args, "%"+params.Search+"%")
			counter++
		}
	}

	clause = append(clause, " deleted_at IS NULL")

	query := `SELECT * FROM cats`

	if len(clause) > 0 {
		query += " WHERE" + strings.Join(clause, " AND")
	}

	if params.ID == "" {
		query += " ORDER BY created_at DESC"

		if params.Limit == 0 || params.Limit > 0 {
			if params.Limit == 0 {
				params.Limit = 5
			}

			query += fmt.Sprintf(" LIMIT $%d", counter)
			args = append(args, params.Limit)
			counter++
		}

		if params.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", counter)
			args = append(args, params.Offset)
		}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cat := &model.Cat{}

		if err := rows.Scan(
			&cat.CatID,
			&cat.UserID,
			&cat.Name,
			&cat.Race,
			&cat.Sex,
			&cat.AgeInMonth,
			&cat.Description,
			&cat.HasMatched,
			pq.Array(&cat.ImageURLs),
			&cat.CreatedAt,
			&cat.UpdatedAt,
			&cat.DeletedAt,
		); err != nil {
			return nil, err
		}

		res = append(res, cat)
	}

	return res, nil
}

// FindOneCatByID implements IFaceRepository.
func (r *Repository) FindOneCatByID(catID string) (*model.Cat, error) {
	res := &model.Cat{}

	if err := r.db.
		QueryRow("SELECT * FROM cats WHERE cat_id = $1 AND deleted_at IS NULL", catID).
		Scan(
			&res.CatID,
			&res.UserID,
			&res.Name,
			&res.Race,
			&res.Sex,
			&res.AgeInMonth,
			&res.Description,
			&res.HasMatched,
			pq.Array(&res.ImageURLs),
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		); err != nil {
		return nil, err
	}

	return res, nil
}

// CheckOwnCat implement IFaceRepository
func (r *Repository) CheckOwnCat(userID uuid.UUID, catID uuid.UUID) error {
	var cnt int

	query := `SELECT count(*) FROM cats where user_id = $1 and cat_id = $2 and deleted_at IS NULL`
	err := r.db.QueryRow(query, userID, catID).Scan(&cnt)
	if err != nil {
		return err
	}

	if cnt == 0 {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  constant.ErrorMessageMap[http.StatusNotFound],
		})

		return err
	}

	return nil
}

// UpdateCatByID implements IFaceRepository.
func (r *Repository) UpdateCatByID(data map[string]interface{}, catID uuid.UUID) error {
	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	for k, v := range data {
		clause = append(clause, fmt.Sprintf("%s = $%d", k, counter))
		args = append(args, v)
		counter++
	}

	query := `UPDATE cats SET `

	if len(clause) > 0 {
		query += strings.Join(clause, ", ")
	}

	query += fmt.Sprintf(" WHERE cat_id = $%d", counter)
	args = append(args, catID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCatByID implements IFaceRepository.
func (r *Repository) DeleteCatByID(catID uuid.UUID) error {
	_, err := r.db.Exec("UPDATE cats SET deleted_at = $1 WHERE cat_id = $2", time.Now().Format("2006-01-02 15:04:05"), catID)
	if err != nil {
		return err
	}

	return nil
}

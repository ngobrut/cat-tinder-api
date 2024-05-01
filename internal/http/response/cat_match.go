package response

import (
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
)

type IssuedByRes struct {
	Name      string `json:"name" db:"u_name"`
	Email     string `json:"email" db:"u_email"`
	CreatedAt string `json:"createdAt" db:"u_created_at"`
}

type MatchCatDetailRes struct {
	CatID       uuid.UUID        `json:"id" db:"c2_cat_id"`
	Name        string           `json:"name" db:"c2_name"`
	Race        constant.CatRace `json:"race" db:"c2_race"`
	Sex         constant.CatSex  `json:"sex" db:"c2_sex"`
	Description string           `json:"description" db:"c2_description"`
	AgeInMonth  int              `json:"ageInMonth" db:"c2_age_in_month"`
	ImageURLs   []string         `json:"imageUrls" db:"c2_image_urls"`
	HasMatched  bool             `json:"hasMatched" db:"c2_has_matched"`
	CreatedAt   string           `json:"createdAt" db:"c2_created_at"`
}

type UserCatDetailRes struct {
	CatID       uuid.UUID        `json:"id" db:"c1_cat_id"`
	Name        string           `json:"name" db:"c1_name"`
	Race        constant.CatRace `json:"race" db:"c1_race"`
	Sex         constant.CatSex  `json:"sex" db:"c1_sex"`
	Description string           `json:"description" db:"c1_description"`
	AgeInMonth  int              `json:"ageInMonth" db:"c1_age_in_month"`
	ImageURLs   []string         `json:"imageUrls" db:"c1_image_urls"`
	HasMatched  bool             `json:"hasMatched" db:"c1_has_matched"`
	CreatedAt   string           `json:"createdAt" db:"c1_created_at"`
}

type CatMatchResponse struct {
	ID             uuid.UUID         `json:"id" db:"id"`
	IssuedBy       IssuedByRes       `json:"issuedBy" db:"-"`
	MatchCatDetail MatchCatDetailRes `json:"matchCatDetail" db:"-"`
	UserCatDetail  UserCatDetailRes  `json:"userCatDetail" db:"-"`
	Message        string            `json:"message" db:"cm_message"`
	CreatedAt      string            `json:"createdAt" db:"cm_created_at"`
}

// type CatMatchRes struct {
// 	CmID             uuid.UUID  `db:"cm_id"`
// 	CmIssuerUserID   uuid.UUID  `db:"cm_issuer_user_id"`
// 	CmIssuerCatID    uuid.UUID  `db:"cm_issuer_cat_id"`
// 	CmReceiverUserID uuid.UUID  `db:"cm_receiver_user_id"`
// 	CmReceiverCatID  uuid.UUID  `db:"cm_receiver_cat_id"`
// 	CmMessage        string     `db:"cm_message"`
// 	CmIsApproved     *bool      `db:"cm_is_approved"`
// 	CmCreatedAt      time.Time  `db:"cm_created_at"`
// 	CmUpdatedAt      time.Time  `db:"cm_updated_at"`
// 	CmDeletedAt      *time.Time `db:"cm_deleted_at"`

// 	IcCatID       uuid.UUID        `db:"c1_cat_id"`
// 	IcUserID      uuid.UUID        `db:"c1_user_id"`
// 	IcName        string           `db:"c1_name"`
// 	IcRace        constant.CatRace `db:"c1_race"`
// 	IcSex         constant.CatSex  `db:"c1_sex"`
// 	IcAgeInMonth  int              `db:"c1_age_in_month"`
// 	IcDescription string           `db:"c1_description"`
// 	IcHasMatched  bool             `db:"c1_has_matched"`
// 	IcImageURLs   []string         `db:"c1_image_urls"`
// 	IcCreatedAt   time.Time        `db:"c1_created_at"`
// 	IcUpdatedAt   time.Time        `db:"c1_updated_at"`
// 	IcDeletedAt   *time.Time       `db:"c1_deleted_at"`

// 	RcCatID       uuid.UUID        `db:"c2_cat_id"`
// 	RcUserID      uuid.UUID        `db:"c2_user_id"`
// 	RcName        string           `db:"c2_name"`
// 	RcRace        constant.CatRace `db:"c2_race"`
// 	RcSex         constant.CatSex  `db:"c2_sex"`
// 	RcAgeInMonth  int              `db:"c2_age_in_month"`
// 	RcDescription string           `db:"c2_description"`
// 	RcHasMatched  bool             `db:"c2_has_matched"`
// 	RcImageURLs   []string         `db:"c2_image_urls"`
// 	RcCreatedAt   time.Time        `db:"c2_created_at"`
// 	RcUpdatedAt   time.Time        `db:"c2_updated_at"`
// 	RcDeletedAt   *time.Time       `db:"c2_deleted_at"`

// 	IssuerUserID    uuid.UUID  `db:"u_user_id"`
// 	IssuerName      string     `db:"u_name"`
// 	IssuerEmail     string     `db:"u_email"`
// 	IssuerPassword  string     `db:"u_password"`
// 	IssuerCreatedAt time.Time  `db:"u_created_at"`
// 	IssuerUpdatedAt time.Time  `db:"u_updated_at"`
// 	IssuerDeletedAt *time.Time `db:"u_deleted_at"`
// }

package model

import (
	"database/sql"
	"privy-test/param/cake"
)

type Cake struct {
	ID          sql.NullInt64   `json:"id" db:"id"`
	Title       sql.NullString  `json:"title" db:"title"`
	Description sql.NullString  `json:"description" db:"description"`
	Rating      sql.NullFloat64 `json:"rating" db:"rating"`
	Image       sql.NullString  `json:"image" db:"image"`
	CreatedAt   sql.NullTime    `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime    `json:"updated_at" db:"updated_at"`
}

func (c *Cake) ParamToModel(param *cake.CreateUpdateRequest) *Cake {
	return &Cake{
		Title:       sql.NullString{String: param.Title, Valid: true},
		Description: sql.NullString{String: param.Description, Valid: true},
		Rating:      sql.NullFloat64{Float64: param.Rating, Valid: true},
		Image:       sql.NullString{String: param.Image, Valid: true},
	}
}

func (c *Cake) ConvertResponse() cake.DetailCakeResponse {
	return cake.DetailCakeResponse{
		ID:          c.ID.Int64,
		Title:       c.Title.String,
		Description: c.Description.String,
		Rating:      c.Rating.Float64,
		Image:       c.Image.String,
		CreatedAt:   c.CreatedAt.Time.String(),
		UpdatedAt:   c.CreatedAt.Time.String(),
	}
}

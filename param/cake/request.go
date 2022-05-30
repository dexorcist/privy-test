package cake

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"privy-test/enum"
	"privy-test/utils"
)

type CreateUpdateRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
}

func (r *CreateUpdateRequest) Validate() error {
	if err := validation.Validate(r.Title, validation.Required); err != nil {
		return utils.MultiStringBadError(enum.HTTPErrorTitleRequired)
	}

	if err := validation.Validate(r.Description, validation.Required); err != nil {
		return utils.MultiStringBadError(enum.HTTPErrorDescriptionRequired)
	}

	if err := validation.Validate(r.Image, validation.Required); err != nil {
		return utils.MultiStringBadError(enum.HTTPErrorImageRequired)
	}

	if r.Rating <= 0 {
		return utils.MultiStringBadError(enum.HTTPErrorRatingRequired)
	}

	return nil
}

type FindAllRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	MinRating   float64 `json:"min_rating"`
	MaxRating   float64 `json:"max_rating"`
}

func (r *FindAllRequest) Validate() error {
	r.Title = "%" + r.Title + "%"
	r.Description = "%" + r.Description + "%"

	if r.MaxRating == 0 {
		r.MaxRating = 10
	}

	if r.MaxRating > 10 {
		return utils.MultiStringBadError(enum.HTTPErrorMaxRatingBad)
	}

	if r.MinRating > r.MaxRating {
		return utils.MultiStringBadError(enum.HTTPErrorMinBiggerMaxRattingBad)
	}

	return nil
}

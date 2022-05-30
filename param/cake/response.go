package cake

type HTTPGetDetailCakeResponse struct {
	Data DetailCakeResponse `json:"data"`
}

type HTTPGetListCakeResponse struct {
	Data []DetailCakeResponse `json:"data"`
}

type DetailCakeResponse struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

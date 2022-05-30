package param

type Response struct {
	Status     string        `json:"status"`
	HTTPCode   int           `json:"-"`
	Payload    interface{}   `json:"payload,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
	Message    *ErrorMessage `json:"message,omitempty"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ErrorMessage struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type CommonErrorResponse struct {
	ErrorID string `json:"error_id"`
	Message struct {
		EN string `json:"en"`
		ID string `json:"id"`
	} `json:"message"`
}

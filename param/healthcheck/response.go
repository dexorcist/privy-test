package healthcheck

type HTTPHealthCheckResponse struct {
	Data DataResponse `json:"data"`
}

type DataResponse struct {
	Environment    string `json:"environment"`
	DatabaseStatus bool   `json:"database_status"`
}

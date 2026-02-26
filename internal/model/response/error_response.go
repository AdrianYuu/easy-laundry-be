package response

type ErrorResponse struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

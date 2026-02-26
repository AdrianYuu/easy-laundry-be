package response

type WebResponse[T any] struct {
	Data   T               `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

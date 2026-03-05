package request

type VerifyRequest struct {
	Token string `validate:"required"`
}

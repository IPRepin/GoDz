package verify

type EmailRequest struct {
	Email string `json:"email" validate:"required, email"`
}

type VerifyRequest struct {
	Hash string `json:"hash" validate:"required"`
}

package verify

type EmailRequest struct {
	Email string `json:"email"`
}

type VerifyRequest struct {
	Hash string `json:"hash"`
}

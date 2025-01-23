package verify

type EmailRequest struct {
	Email   string `json:"email"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type VerifyRequest struct {
	Hash string `json:"hash"`
}

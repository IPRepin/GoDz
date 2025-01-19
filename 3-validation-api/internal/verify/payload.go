package verify

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type VerifyRequest struct {
	Hash string `json:"hash"`
}

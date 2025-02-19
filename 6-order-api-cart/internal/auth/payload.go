package auth

type AuthResponse struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id,omitempty"`
}

type AuthRequest struct {
	Phone     string `json:"phone" validate:"required, phone"`
	SessionId string `json:"session_id" validate:"required"`
	Code      string `json:"code"`
}

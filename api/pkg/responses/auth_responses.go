package responses

// AuthResponse is the response for authentication
type AuthResponse struct {
	response
	Data AuthData `json:"data"`
}

// AuthData contains authentication data
type AuthData struct {
	APIKey string `json:"api_key" example:"uk_abc123..."`
	Email  string `json:"email" example:"user@example.com"`
	UserID string `json:"user_id" example:"WB7DRDWrJZRGbYrv2CKGkqbzvqdC"`
}


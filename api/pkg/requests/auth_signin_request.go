package requests

import "strings"

// AuthSignIn is the payload for user sign in
type AuthSignIn struct {
	request
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securePassword123"`
}

// Sanitize sanitizes the AuthSignIn request
func (input *AuthSignIn) Sanitize() AuthSignIn {
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	input.Password = strings.TrimSpace(input.Password)
	return *input
}


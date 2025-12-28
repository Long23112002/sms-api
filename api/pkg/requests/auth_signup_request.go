package requests

import "strings"

// AuthSignUp is the payload for user sign up
type AuthSignUp struct {
	request
	Email            string `json:"email" example:"user@example.com"`
	Password         string `json:"password" example:"securePassword123"`
	ConfirmPassword  string `json:"confirm_password" example:"securePassword123"`
}

// Sanitize sanitizes the AuthSignUp request
func (input *AuthSignUp) Sanitize() AuthSignUp {
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	input.Password = strings.TrimSpace(input.Password)
	input.ConfirmPassword = strings.TrimSpace(input.ConfirmPassword)
	return *input
}


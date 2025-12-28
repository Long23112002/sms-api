package validators

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/NdoleStudio/httpsms/pkg/requests"
	"github.com/NdoleStudio/httpsms/pkg/telemetry"
	"github.com/thedevsaddam/govalidator"
)

// AuthHandlerValidator validates models used in handlers.AuthHandler
type AuthHandlerValidator struct {
	validator
	logger telemetry.Logger
	tracer telemetry.Tracer
}

// NewAuthHandlerValidator creates a new handlers.AuthHandler validator
func NewAuthHandlerValidator(
	logger telemetry.Logger,
	tracer telemetry.Tracer,
) *AuthHandlerValidator {
	return &AuthHandlerValidator{
		logger: logger.WithService(fmt.Sprintf("%T", &AuthHandlerValidator{})),
		tracer: tracer,
	}
}

// ValidateSignUp validates requests.AuthSignUp
func (validator *AuthHandlerValidator) ValidateSignUp(_ context.Context, request requests.AuthSignUp) url.Values {
	v := govalidator.New(govalidator.Options{
		Data: &request,
		Rules: govalidator.MapData{
			"email": []string{
				"required",
				"email",
				"min:3",
				"max:255",
			},
			"password": []string{
				"required",
				"min:6",
				"max:100",
			},
			"confirm_password": []string{
				"required",
				"min:6",
				"max:100",
			},
		},
	})

	errors := v.ValidateStruct()

	// Additional password validation
	if request.Password != "" && len(request.Password) < 6 {
		if errors == nil {
			errors = url.Values{}
		}
		errors.Add("password", "The password field must be at least 6 characters long")
	}

	// Confirm password validation
	if request.ConfirmPassword != "" && request.Password != request.ConfirmPassword {
		if errors == nil {
			errors = url.Values{}
		}
		errors.Add("confirm_password", "The confirm password field must match the password field")
	}

	// Email format validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if request.Email != "" && !emailRegex.MatchString(request.Email) {
		if errors == nil {
			errors = url.Values{}
		}
		errors.Add("email", "The email field must be a valid email address")
	}

	return errors
}

// ValidateSignIn validates requests.AuthSignIn
func (validator *AuthHandlerValidator) ValidateSignIn(_ context.Context, request requests.AuthSignIn) url.Values {
	v := govalidator.New(govalidator.Options{
		Data: &request,
		Rules: govalidator.MapData{
			"email": []string{
				"required",
				"email",
				"min:3",
				"max:255",
			},
			"password": []string{
				"required",
				"min:1",
			},
		},
	})

	return v.ValidateStruct()
}


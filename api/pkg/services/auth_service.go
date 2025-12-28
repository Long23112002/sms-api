package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/NdoleStudio/httpsms/pkg/repositories"
	"github.com/NdoleStudio/httpsms/pkg/telemetry"
	"github.com/palantir/stacktrace"
)

const (
	// ErrCodeUserExists is thrown when a user already exists
	ErrCodeUserExists = stacktrace.ErrorCode(2000)
	// ErrCodeInvalidCredentials is thrown when credentials are invalid
	ErrCodeInvalidCredentials = stacktrace.ErrorCode(2001)
)

// AuthService handles authentication operations
type AuthService struct {
	logger         telemetry.Logger
	tracer         telemetry.Tracer
	userRepository repositories.UserRepository
}

// NewAuthService creates a new AuthService
func NewAuthService(
	logger telemetry.Logger,
	tracer telemetry.Tracer,
	userRepository repositories.UserRepository,
) *AuthService {
	return &AuthService{
		logger:         logger.WithService("AuthService"),
		tracer:         tracer,
		userRepository: userRepository,
	}
}

// SignUpParams are parameters for signing up a user
type SignUpParams struct {
	Email    string
	Password string
}

// SignInParams are parameters for signing in a user
type SignInParams struct {
	Email    string
	Password string
}

// SignUpResult contains the result of a sign up operation
type SignUpResult struct {
	APIKey string
	Email  string
	UserID string
}

// SignInResult contains the result of a sign in operation
type SignInResult struct {
	APIKey string
	Email  string
	UserID string
}

// generateAPIKey generates a random API key
func (service *AuthService) generateAPIKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", stacktrace.Propagate(err, "cannot generate random bytes")
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// hashPassword hashes a password using bcrypt
func (service *AuthService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", stacktrace.Propagate(err, "cannot hash password")
	}
	return string(hash), nil
}

// comparePassword compares a password with a hash
func (service *AuthService) comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// SignUp creates a new user account
func (service *AuthService) SignUp(ctx context.Context, params SignUpParams) (*SignUpResult, error) {
	ctx, span, ctxLogger := service.tracer.StartWithLogger(ctx, service.logger)
	defer span.End()

	// Check if user already exists
	existingUser, err := service.userRepository.LoadByEmail(ctx, params.Email)
	if err == nil && existingUser != nil {
		msg := fmt.Sprintf("user with email [%s] already exists", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.PropagateWithCode(errors.New("user already exists"), ErrCodeUserExists, msg))
	}
	// If error is not "not found", it's a real error
	if err != nil && stacktrace.GetCode(err) != repositories.ErrCodeNotFound {
		msg := fmt.Sprintf("cannot check if user exists with email [%s]", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.Propagate(err, msg))
	}

	// Hash password
	passwordHash, err := service.hashPassword(params.Password)
	if err != nil {
		msg := fmt.Sprintf("cannot hash password for email [%s]", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.Propagate(err, msg))
	}

	// Generate API key
	apiKeyBytes, err := service.generateAPIKey(64)
	if err != nil {
		msg := fmt.Sprintf("cannot generate API key for email [%s]", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.Propagate(err, msg))
	}

	// Generate user ID
	userID := entities.UserID(uuid.New().String())

	// Create user
	user := &entities.User{
		ID:               userID,
		Email:            params.Email,
		PasswordHash:     passwordHash,
		APIKey:           "uk_" + apiKeyBytes,
		SubscriptionName: entities.SubscriptionNameFree,
		Timezone:         "UTC",
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}

	// Store user in database
	if err := service.userRepository.Store(ctx, user); err != nil {
		msg := fmt.Sprintf("cannot create user in database for email [%s]", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.Propagate(err, msg))
	}

	ctxLogger.Info(fmt.Sprintf("new user created with ID [%s] and email [%s]", user.ID, user.Email))

	return &SignUpResult{
		APIKey: user.APIKey,
		Email:  user.Email,
		UserID: string(user.ID),
	}, nil
}

// SignIn authenticates an existing user
func (service *AuthService) SignIn(ctx context.Context, params SignInParams) (*SignInResult, error) {
	ctx, span, ctxLogger := service.tracer.StartWithLogger(ctx, service.logger)
	defer span.End()

	// Load user by email
	user, err := service.userRepository.LoadByEmail(ctx, params.Email)
	if err != nil {
		// Check if user not found
		if stacktrace.GetCode(err) == repositories.ErrCodeNotFound {
			msg := fmt.Sprintf("user with email [%s] does not exist", params.Email)
			return nil, service.tracer.WrapErrorSpan(span, stacktrace.PropagateWithCode(errors.New("invalid credentials"), ErrCodeInvalidCredentials, msg))
		}
		msg := fmt.Sprintf("cannot load user with email [%s]", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.Propagate(err, msg))
	}

	// Check if user has password hash (custom auth user)
	if user.PasswordHash == "" {
		msg := fmt.Sprintf("user with email [%s] does not have password set", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.PropagateWithCode(errors.New("invalid credentials"), ErrCodeInvalidCredentials, msg))
	}

	// Verify password
	if !service.comparePassword(user.PasswordHash, params.Password) {
		msg := fmt.Sprintf("invalid password for email [%s]", params.Email)
		return nil, service.tracer.WrapErrorSpan(span, stacktrace.PropagateWithCode(errors.New("invalid credentials"), ErrCodeInvalidCredentials, msg))
	}

	ctxLogger.Info(fmt.Sprintf("user signed in with ID [%s] and email [%s]", user.ID, user.Email))

	return &SignInResult{
		APIKey: user.APIKey,
		Email:  user.Email,
		UserID: string(user.ID),
	}, nil
}

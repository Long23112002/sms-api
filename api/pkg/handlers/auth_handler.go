package handlers

import (
	"fmt"
	"net/url"

	"github.com/NdoleStudio/httpsms/pkg/requests"
	"github.com/NdoleStudio/httpsms/pkg/responses"
	"github.com/NdoleStudio/httpsms/pkg/services"
	"github.com/NdoleStudio/httpsms/pkg/telemetry"
	"github.com/NdoleStudio/httpsms/pkg/validators"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/palantir/stacktrace"
)

// AuthHandler handles authentication http requests
type AuthHandler struct {
	handler
	logger    telemetry.Logger
	tracer    telemetry.Tracer
	validator *validators.AuthHandlerValidator
	service   *services.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(
	logger telemetry.Logger,
	tracer telemetry.Tracer,
	validator *validators.AuthHandlerValidator,
	service *services.AuthService,
) *AuthHandler {
	return &AuthHandler{
		logger:    logger.WithService(fmt.Sprintf("%T", &AuthHandler{})),
		tracer:    tracer,
		validator: validator,
		service:   service,
	}
}

// RegisterRoutes registers the routes for the AuthHandler
func (h *AuthHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/v1/auth/signup", h.SignUp)
	router.Post("/v1/auth/signin", h.SignIn)
}

// SignUp handles user sign up
// @Summary      Sign up a new user
// @Description  Creates a new user account with email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload   	body 		requests.AuthSignUp  			true 	"Sign up payload"
// @Success      201 		{object}	responses.AuthResponse
// @Failure      400		{object}	responses.BadRequest
// @Failure      422		{object}	responses.UnprocessableEntity
// @Failure      500		{object}	responses.InternalServerError
// @Router       /auth/signup [post]
func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	ctx, span, ctxLogger := h.tracer.StartFromFiberCtxWithLogger(c, h.logger)
	defer span.End()

	var request requests.AuthSignUp
	if err := c.BodyParser(&request); err != nil {
		msg := fmt.Sprintf("cannot parse sign up request [%s]", c.OriginalURL())
		ctxLogger.Warn(stacktrace.Propagate(err, msg))
		return h.responseBadRequest(c, err)
	}

	sanitized := request.Sanitize()
	if errors := h.validator.ValidateSignUp(ctx, sanitized); len(errors) != 0 {
		msg := fmt.Sprintf("validation errors [%s] while signing up user [%+#v]", spew.Sdump(errors), sanitized)
		ctxLogger.Warn(stacktrace.NewError(msg))
		return h.responseUnprocessableEntity(c, errors, "validation errors while signing up")
	}

	result, err := h.service.SignUp(ctx, services.SignUpParams{
		Email:    sanitized.Email,
		Password: sanitized.Password,
	})
	if err != nil {
		// Check error code
		if stacktrace.GetCode(err) == services.ErrCodeUserExists {
			return h.responseUnprocessableEntity(c, url.Values{
				"email": []string{"Email đã được sử dụng"},
			}, "Email đã được sử dụng")
		}
		msg := fmt.Sprintf("cannot sign up user with email [%s]", sanitized.Email)
		ctxLogger.Error(stacktrace.Propagate(err, msg))
		return h.responseInternalServerError(c)
	}

	return h.responseCreated(c, "user signed up successfully", responses.AuthData{
		APIKey: result.APIKey,
		Email:  result.Email,
		UserID: result.UserID,
	})
}

// SignIn handles user sign in
// @Summary      Sign in a user
// @Description  Authenticates a user with email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload   	body 		requests.AuthSignIn  			true 	"Sign in payload"
// @Success      200 		{object}	responses.AuthResponse
// @Failure      400		{object}	responses.BadRequest
// @Failure      422		{object}	responses.UnprocessableEntity
// @Failure      500		{object}	responses.InternalServerError
// @Router       /auth/signin [post]
func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	ctx, span, ctxLogger := h.tracer.StartFromFiberCtxWithLogger(c, h.logger)
	defer span.End()

	var request requests.AuthSignIn
	if err := c.BodyParser(&request); err != nil {
		msg := fmt.Sprintf("cannot parse sign in request [%s]", c.OriginalURL())
		ctxLogger.Warn(stacktrace.Propagate(err, msg))
		return h.responseBadRequest(c, err)
	}

	sanitized := request.Sanitize()
	if errors := h.validator.ValidateSignIn(ctx, sanitized); len(errors) != 0 {
		msg := fmt.Sprintf("validation errors [%s] while signing in user [%+#v]", spew.Sdump(errors), sanitized)
		ctxLogger.Warn(stacktrace.NewError(msg))
		return h.responseUnprocessableEntity(c, errors, "validation errors while signing in")
	}

	result, err := h.service.SignIn(ctx, services.SignInParams{
		Email:    sanitized.Email,
		Password: sanitized.Password,
	})
	if err != nil {
		// Check error code
		if stacktrace.GetCode(err) == services.ErrCodeInvalidCredentials {
			return h.responseUnprocessableEntity(c, url.Values{
				"email": []string{"Email hoặc mật khẩu không đúng"},
			}, "Email hoặc mật khẩu không đúng")
		}
		msg := fmt.Sprintf("cannot sign in user with email [%s]", sanitized.Email)
		ctxLogger.Error(stacktrace.Propagate(err, msg))
		return h.responseInternalServerError(c)
	}

	return h.responseOK(c, "user signed in successfully", responses.AuthData{
		APIKey: result.APIKey,
		Email:  result.Email,
		UserID: result.UserID,
	})
}


package rest

import (
	"net/http"

	"forum.id/internal/domain"
	"forum.id/pkg/security"
	"forum.id/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	Router fiber.Router
	Repo   domain.UserRepository
	Log    *logrus.Logger
	Val    *validator.Validate
}

func NewAuthHandler(
	Router fiber.Router,
	Repo domain.UserRepository,
	Log *logrus.Logger,
	Val *validator.Validate,
) *authHandler {
	handler := &authHandler{Router, Repo, Log, Val}

	rest := Router.Group("/api/v1/auth")
	rest.Post("/sign-up", handler.SignUp)
	rest.Post("/sign-in", handler.SignIn)

	return handler
}

func (h *authHandler) SignUp(c *fiber.Ctx) error {
	var req NewUser

	if err := c.BodyParser(&req); err != nil {
		h.Log.WithError(err).Warn("AuthHandler.SignUp: failed to parse request body")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := h.Val.Struct(req); err != nil {
		h.Log.WithError(err).Warn("AuthHandler.SignUp: validation failed")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: err.Error(),
		})
	}

	existsUsername, err := h.Repo.ExistUsernameIgnoreCase(c.Context(), req.Username)
	if err != nil {
		h.Log.WithError(err).Error("AuthHandler.SignUp: failed to check username existence")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Internal server error",
		})
	}
	if existsUsername {
		return c.Status(http.StatusConflict).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Username already exists",
		})
	}

	existsEmail, err := h.Repo.ExistEmailIgnoreCase(c.Context(), req.Email)
	if err != nil {
		h.Log.WithError(err).Error("AuthHandler.SignUp: failed to check email existence")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Internal server error",
		})
	}

	if existsEmail {
		return c.Status(http.StatusConflict).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Email already exists",
		})
	}

	var user domain.User
	utils.Copy(&user, &req)

	savedUser, err := h.Repo.Save(c.Context(), &user)
	if err != nil {
		h.Log.WithError(err).Error("AuthHandler.SignUp: failed to save user")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to create user",
		})
	}

	var userProfile UserProfile
	utils.Copy(&userProfile, savedUser)

	return c.Status(http.StatusCreated).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "User successfully created",
		Data: struct {
			User interface{} `json:"user"`
		}{
			User: userProfile,
		},
	})
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	var req AuthUser

	if err := c.BodyParser(&req); err != nil {
		h.Log.WithError(err).Warn("AuthHandler.SignIn: failed to parse request body")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := h.Val.Struct(req); err != nil {
		h.Log.WithError(err).Warn("AuthHandler.SignIn: validation failed")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: err.Error(),
		})
	}

	user, err := h.Repo.FetchByEmail(c.Context(), req.Email)
	if err != nil {
		h.Log.WithError(err).Errorf("AuthHandler.SignIn: failed to fetch user by email %s", req.Email)
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Internal server error",
		})
	}

	if user == nil {
		h.Log.Warnf("AuthHandler.SignIn: user with email %s not found", req.Email)
		return c.Status(http.StatusUnauthorized).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		h.Log.Warn("AuthHandler.SignIn: password mismatch")
		return c.Status(http.StatusUnauthorized).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid credentials",
		})
	}

	tokenString := security.GenerateJWT(user.ID)

	var userCredential UserCredential
	userCredential.Token = tokenString
	utils.Copy(&userCredential.User, user)

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "User authentication successful",
		Data:    userCredential,
	})
}

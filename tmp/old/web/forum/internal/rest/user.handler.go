package rest

import (
	"net/http"

	"forum.id/internal/domain"
	"forum.id/internal/middleware"
	"forum.id/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type userHandler struct {
	Router fiber.Router
	Repo   domain.UserRepository
	Log    *logrus.Logger
	Val    *validator.Validate
}

func NewUserHandler(
	Router fiber.Router,
	Repo domain.UserRepository,
	Log *logrus.Logger,
	Val *validator.Validate,
) *userHandler {
	handler := &userHandler{Router, Repo, Log, Val}

	rest := Router.Group("/api/v1/users")
	rest.Get("/", handler.GetAll)
	rest.Get("/me", middleware.JWTAuthorization, handler.GetProfile)

	return handler
}

func (h *userHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.Repo.FetchAll(c.Context())
	if err != nil {
		h.Log.WithError(err).Error("UserHandler.GetAll: failed to fetch all users")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to retrieve users",
		})
	}

	var userProfiles []UserProfile
	utils.Copy(&userProfiles, &users)

	res := struct {
		Users interface{} `json:"users"`
	}{
		Users: userProfiles,
	}

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    res,
	})
}

func (h *userHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int64)

	user, err := h.Repo.FetchByID(c.Context(), userID)
	if err != nil {
		h.Log.WithError(err).Errorf("UserHandler.GetProfile: failed to fetch user with ID %d", userID)
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to retrieve user profile",
		})
	}

	if user == nil {
		h.Log.Warnf("UserHandler.GetProfile: user with ID %d not found", userID)
		return c.Status(http.StatusNotFound).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "User not found",
		})
	}

	var userProfile UserProfile
	utils.Copy(&userProfile, user)

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "User profile retrieved successfully",
		Data: struct {
			User interface{} `json:"user"`
		}{
			User: userProfile,
		},
	})
}

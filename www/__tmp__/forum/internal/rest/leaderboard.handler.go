package rest

import (
	"net/http"
	"sort"

	"forum.id/internal/domain"
	"forum.id/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type leaderBoardHandler struct {
	Router      fiber.Router
	RepoUser    domain.UserRepository
	RepoThread  domain.ThreadRepository
	RepoComment domain.CommentRepository
	Log         *logrus.Logger
}

func NewLeaderBoardHandler(
	Router fiber.Router,
	RepoUser domain.UserRepository,
	RepoThread domain.ThreadRepository,
	RepoComment domain.CommentRepository,
	Log *logrus.Logger,
) *leaderBoardHandler {
	handler := &leaderBoardHandler{Router, RepoUser, RepoThread, RepoComment, Log}

	rest := Router.Group("/api/v1/leaderBoards")
	rest.Get("/", handler.Get)

	return handler
}

func (h *leaderBoardHandler) Get(c *fiber.Ctx) error {
	// Fetch all users
	users, err := h.RepoUser.FetchAll(c.Context())
	if err != nil {
		h.Log.Warnf("LeaderBoardHandler.Get: failed to fetch users: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to fetch leaderBoards",
		})
	}

	// Initialize map to store scores per user
	userScores := make(map[int64]int)

	// Loop through each user and fetch threads and comments
	for _, user := range users {
		// Fetch all threads by user
		threads, err := h.RepoThread.FetchAllByUserID(c.Context(), user.ID)
		if err != nil {
			h.Log.Warnf("LeaderBoardHandler.Get: failed to fetch threads for user %d: %v", user.ID, err)
			return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
				Status:  "error",
				Message: "Failed to fetch threads",
			})
		}

		// Fetch all comments by user
		comments, err := h.RepoComment.FetchAllByUserID(c.Context(), user.ID)
		if err != nil {
			h.Log.Warnf("LeaderBoardHandler.Get: failed to fetch comments for user %d: %v", user.ID, err)
			return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
				Status:  "error",
				Message: "Failed to fetch comments",
			})
		}

		// Calculate score for threads
		for _, thread := range threads {
			upVotes, downVotes := 0, 0
			for _, vote := range thread.Votes {
				if vote.VoteType == string(domain.UpVote) {
					upVotes++
				} else if vote.VoteType == string(domain.DownVote) {
					downVotes++
				}
			}
			userScores[user.ID] += upVotes - downVotes
		}

		// Calculate score for comments
		for _, comment := range comments {
			upVotes, downVotes := 0, 0
			for _, vote := range comment.Votes {
				if vote.VoteType == string(domain.UpVote) {
					upVotes++
				} else if vote.VoteType == string(domain.DownVote) {
					downVotes++
				}
			}
			userScores[user.ID] += upVotes - downVotes
		}
	}

	// Create the leaderBoard result
	var leaderBoards []LeaderBoard

	for _, user := range users {
		score := userScores[user.ID]
		var userProfile UserProfile
		utils.Copy(&userProfile, &user)
		leaderBoards = append(leaderBoards, LeaderBoard{
			User:  userProfile,
			Score: score,
		})
	}

	// Sort leaderBoard by score from highest to lowest
	sort.Slice(leaderBoards, func(i, j int) bool {
		return leaderBoards[i].Score > leaderBoards[j].Score
	})

	// Return leaderBoard response
	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Fetch Leader Boards Successfully",
		Data: struct {
			LeaderBoards interface{} `json:"leaderboards"`
		}{
			LeaderBoards: leaderBoards,
		},
	})
}

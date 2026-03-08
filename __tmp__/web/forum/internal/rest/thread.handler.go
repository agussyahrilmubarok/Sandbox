package rest

import (
	"net/http"
	"strconv"
	"strings"

	"forum.id/internal/domain"
	"forum.id/internal/middleware"
	"forum.id/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type threadHandler struct {
	Router      fiber.Router
	RepoThread  domain.ThreadRepository
	RepoComment domain.CommentRepository
	RepoVote    domain.VoteRepository
	Log         *logrus.Logger
	Val         *validator.Validate
}

// NewThreadHandler sets up thread routes and returns a new threadHandler
func NewThreadHandler(
	Router fiber.Router,
	RepoThread domain.ThreadRepository,
	RepoComment domain.CommentRepository,
	RepoVote domain.VoteRepository,
	Log *logrus.Logger,
	Val *validator.Validate,
) *threadHandler {
	handler := &threadHandler{Router, RepoThread, RepoComment, RepoVote, Log, Val}

	rest := Router.Group("/api/v1/threads")
	rest.Get("/", handler.GetAllThread)
	rest.Get("/:id", handler.GetThreadDetail)
	rest.Post("/", middleware.JWTAuthorization, handler.CreateThread)
	rest.Post("/:id/up-vote", middleware.JWTAuthorization, handler.VoteThread)
	rest.Post("/:id/down-vote", middleware.JWTAuthorization, handler.VoteThread)
	rest.Post("/:id/neutral-vote", middleware.JWTAuthorization, handler.VoteThread)
	rest.Post("/:id/comments", middleware.JWTAuthorization, handler.CreateComment)
	rest.Post("/:id/comments/:commentId/up-vote", middleware.JWTAuthorization, handler.VoteComment)
	rest.Post("/:id/comments/:commentId/down-vote", middleware.JWTAuthorization, handler.VoteComment)
	rest.Post("/:id/comments/:commentId/neutral-vote", middleware.JWTAuthorization, handler.VoteComment)

	return handler
}

func (h *threadHandler) GetAllThread(c *fiber.Ctx) error {
	threads, err := h.RepoThread.FetchAll(c.Context())
	if err != nil {
		h.Log.WithError(err).Error("Failed to fetch all threads")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to retrieve threads",
		})

	}

	var threadTiles []ThreadTile
	for _, thread := range threads {
		var upVotesBy []string
		var downVotesBy []string
		if thread.Votes != nil {
			for _, vote := range thread.Votes {
				if vote.VoteType == string(domain.UpVote) {
					upVotesBy = append(upVotesBy, vote.User.Name)
				} else if vote.VoteType == string(domain.DownVote) {
					downVotesBy = append(downVotesBy, vote.User.Name)
				}
			}
		}

		threadTiles = append(threadTiles, ThreadTile{
			ID:            thread.ID,
			Title:         thread.Title,
			Body:          thread.Body,
			Category:      thread.Category,
			CreatedAt:     thread.CreatedAt,
			OwnerId:       thread.UserID,
			UpVotesBy:     utils.StructNonNilSlice(upVotesBy),
			DownVotesBy:   utils.StructNonNilSlice(downVotesBy),
			TotalComments: len(thread.Comments),
		})
	}

	res := struct {
		Threads []ThreadTile `json:"threads"`
	}{
		Threads: threadTiles,
	}

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Threads retrieved successfully",
		Data:    res,
	})
}

func (h *threadHandler) GetThreadDetail(c *fiber.Ctx) error {
	idParam := c.Params("id")
	threadID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.Log.WithError(err).Errorf("Invalid thread ID: %v", idParam)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid thread id",
		})
	}

	thread, err := h.RepoThread.FetchByID(c.Context(), threadID)
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to fetch thread with ID: %d", threadID)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to retrieve thread",
		})
	}
	if thread == nil {
		return c.Status(http.StatusNotFound).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Thread not found",
		})
	}

	var upVotesBy []string
	var downVotesBy []string
	if thread.Votes != nil {
		for _, vote := range thread.Votes {
			if vote.VoteType == string(domain.UpVote) {
				upVotesBy = append(upVotesBy, vote.User.Name)
			} else if vote.VoteType == string(domain.DownVote) {
				downVotesBy = append(downVotesBy, vote.User.Name)
			}
		}
	}

	var comments []CommentDetail
	for _, comment := range thread.Comments {
		var upVotesByComment []string
		var downVotesByComment []string
		for _, vote := range comment.Votes {
			if vote.VoteType == string(domain.UpVote) {
				upVotesByComment = append(upVotesByComment, vote.User.Name)
			} else if vote.VoteType == string(domain.DownVote) {
				downVotesByComment = append(downVotesByComment, vote.User.Name)
			}
		}

		comments = append(comments, CommentDetail{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			Owner: UserProfile{
				ID:     comment.User.ID,
				Name:   comment.User.Name,
				Email:  comment.User.Email,
				Avatar: comment.User.Avatar,
			},
			UpVotesBy:   utils.StructNonNilSlice(upVotesByComment),
			DownVotesBy: utils.StructNonNilSlice(downVotesByComment),
		})
	}

	threadDetail := ThreadDetail{
		ID:        thread.ID,
		Title:     thread.Title,
		Body:      thread.Body,
		Category:  thread.Category,
		CreatedAt: thread.CreatedAt,
		Owner: UserProfile{
			ID:     thread.User.ID,
			Name:   thread.User.Name,
			Email:  thread.User.Email,
			Avatar: thread.User.Avatar,
		},
		UpVotesBy:   utils.StructNonNilSlice(upVotesBy),
		DownVotesBy: utils.StructNonNilSlice(downVotesBy),
		Comments:    comments,
	}

	res := struct {
		ThreadDetail ThreadDetail `json:"detailThread"`
	}{
		ThreadDetail: threadDetail,
	}

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Thread detail retrieved successfully",
		Data:    res,
	})
}

func (h *threadHandler) CreateThread(c *fiber.Ctx) error {
	var req NewThread
	if err := c.BodyParser(&req); err != nil {
		h.Log.WithError(err).Error("Failed to parse thread request")
		return c.Status(http.StatusNotFound).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid request payload",
		})
	}

	if err := h.Val.Struct(req); err != nil {
		h.Log.WithError(err).Error("Validation failed for thread data")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Validation error",
		})
	}

	userID := c.Locals("userID").(int64)
	req.UserID = userID

	var thread domain.Thread
	utils.Copy(&thread, &req)

	createdThread, err := h.RepoThread.Save(c.Context(), &thread)
	if err != nil {
		h.Log.WithError(err).Error("Failed to create thread")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to create thread",
		})
	}

	var upVotesBy []string
	var downVotesBy []string
	if createdThread.Votes != nil {
		for _, vote := range createdThread.Votes {
			if vote.VoteType == string(domain.UpVote) {
				upVotesBy = append(upVotesBy, vote.User.Name)
			} else if vote.VoteType == string(domain.UpVote) {
				downVotesBy = append(downVotesBy, vote.User.Name)
			}
		}
	}

	threadTile := ThreadTile{
		ID:            createdThread.ID,
		Title:         createdThread.Title,
		Body:          createdThread.Body,
		Category:      createdThread.Category,
		CreatedAt:     createdThread.CreatedAt,
		OwnerId:       createdThread.UserID,
		UpVotesBy:     utils.StructNonNilSlice(upVotesBy),
		DownVotesBy:   utils.StructNonNilSlice(downVotesBy),
		TotalComments: len(createdThread.Comments),
	}

	res := struct {
		Thread ThreadTile `json:"thread"`
	}{
		Thread: threadTile,
	}

	return c.Status(http.StatusCreated).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Thread created successfully",
		Data:    res,
	})
}

func (h *threadHandler) VoteThread(c *fiber.Ctx) error {
	idParam := c.Params("id")
	threadID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.Log.WithError(err).Errorf("Invalid thread ID: %s", idParam)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid thread ID",
		})
	}

	voteType := ""
	switch {
	case strings.HasSuffix(c.Path(), "/up-vote"):
		voteType = string(domain.UpVote)
	case strings.HasSuffix(c.Path(), "/down-vote"):
		voteType = string(domain.DownVote)
	case strings.HasSuffix(c.Path(), "/neutral-vote"):
		voteType = string(domain.NeutralVote)
	default:
		h.Log.WithError(err).Errorf("Invalid vote type in path: %s", c.Path())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vote type",
		})
	}

	userID, ok := c.Locals("userID").(int64)
	if !ok {
		h.Log.Error("userID not found in context")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Unauthorized",
		})
	}

	exists, err := h.RepoThread.ExistByID(c.Context(), threadID)
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to check thread existence: %d", threadID)
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to check thread existence",
		})
	}
	if !exists {
		h.Log.Warnf("Thread with ID %d not found", threadID)
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Thread not found",
		})
	}

	vote := domain.VoteThread{
		VoteType: voteType,
		UserID:   userID,
		ThreadID: threadID,
	}

	voted, err := h.RepoVote.SetVoteThread(c.Context(), &vote)
	if err != nil {
		h.Log.WithError(err).Error("Failed to set vote")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to vote thread",
		})
	}

	var voteTypeInt int
	switch voteType {
	case string(domain.UpVote):
		voteTypeInt = 1
	case string(domain.DownVote):
		voteTypeInt = -1
	case string(domain.NeutralVote):
		voteTypeInt = 0
	}

	voteThread := VoteThread{
		ID:       voted.ID,
		UserID:   voted.UserID,
		ThreadID: voted.ThreadID,
		VoteType: voteTypeInt,
	}

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Vote cast successfully",
		Data: struct {
			Vote interface{} `json:"vote"`
		}{
			Vote: voteThread,
		},
	})
}

func (h *threadHandler) CreateComment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	threadID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.Log.WithError(err).Errorf("Invalid thread ID: %s", idParam)
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid thread ID",
		})
	}

	var req NewComment
	if err := c.BodyParser(&req); err != nil {
		h.Log.WithError(err).Error("Failed to parse comment request")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid request payload",
		})
	}

	if err := h.Val.Struct(req); err != nil {
		h.Log.WithError(err).Error("Validation failed for comment data")
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Validation error",
		})
	}

	userID := c.Locals("userID").(int64)
	req.UserID = userID
	req.ThreadID = threadID

	var comment domain.Comment
	utils.Copy(&comment, &req)

	createdComment, err := h.RepoComment.Save(c.Context(), &comment)
	if err != nil {
		h.Log.WithError(err).Error("Failed to create comment")
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to create comment",
		})
	}

	var upVotesBy []string
	var downVotesBy []string
	if createdComment.Votes != nil {
		for _, vote := range createdComment.Votes {
			if vote.VoteType == string(domain.UpVote) {
				upVotesBy = append(upVotesBy, vote.User.Name)
			} else if vote.VoteType == string(domain.DownVote) {
				downVotesBy = append(downVotesBy, vote.User.Name)
			}
		}
	}

	commentDetail := CommentDetail{
		ID:        createdComment.ID,
		Content:   createdComment.Content,
		CreatedAt: createdComment.CreatedAt,
		Owner: UserProfile{
			ID:     createdComment.User.ID,
			Name:   createdComment.User.Name,
			Email:  createdComment.User.Email,
			Avatar: createdComment.User.Avatar,
		},
		UpVotesBy:   utils.StructNonNilSlice(upVotesBy),
		DownVotesBy: utils.StructNonNilSlice(downVotesBy),
	}

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Comment created successfully",
		Data: struct {
			Comment interface{} `json:"comment"`
		}{
			Comment: commentDetail,
		},
	})
}

func (h *threadHandler) VoteComment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	threadID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.Log.WithError(err).Errorf("Invalid thread ID: %s", idParam)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid thread ID",
		})
	}

	idCommentParam := c.Params("commentId")
	commentID, err := strconv.ParseInt(idCommentParam, 10, 64)
	if err != nil {
		h.Log.WithError(err).Errorf("Invalid thread ID: %s", idParam)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid comment ID",
		})
	}

	voteType := ""
	switch {
	case strings.HasSuffix(c.Path(), "/up-vote"):
		voteType = string(domain.UpVote)
	case strings.HasSuffix(c.Path(), "/down-vote"):
		voteType = string(domain.DownVote)
	case strings.HasSuffix(c.Path(), "/neutral-vote"):
		voteType = string(domain.NeutralVote)
	default:
		h.Log.WithError(err).Errorf("Invalid vote type in path: %s", c.Path())
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Invalid vote type",
		})
	}

	userID, ok := c.Locals("userID").(int64)
	if !ok {
		h.Log.Error("userID not found in context")
		return c.Status(http.StatusUnauthorized).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Unauthorized",
		})
	}

	exists, err := h.RepoThread.ExistByID(c.Context(), threadID)
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to check thread existence: %d", threadID)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to check thread existence",
		})
	}
	if !exists {
		h.Log.Warnf("Thread with ID %d not found", threadID)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Thread not found",
		})
	}

	exists, err = h.RepoComment.ExistByID(c.Context(), commentID)
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to check thread existence: %d", threadID)
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to check thread existence",
		})
	}
	if !exists {
		h.Log.Warnf("Comment with ID %d not found", threadID)
		return c.Status(http.StatusNotFound).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Comment not found",
		})
	}

	vote := domain.VoteComment{
		VoteType:  voteType,
		UserID:    userID,
		CommentID: commentID,
	}

	voted, err := h.RepoVote.SetVoteComment(c.Context(), &vote)
	if err != nil {
		h.Log.WithError(err).Error("Failed to set vote")
		return c.Status(http.StatusNotFound).JSON(utils.ResponseFailure{
			Status:  "error",
			Message: "Failed to vote comment",
		})
	}

	var voteTypeInt int
	switch voteType {
	case string(domain.UpVote):
		voteTypeInt = 1
	case string(domain.DownVote):
		voteTypeInt = -1
	case string(domain.NeutralVote):
		voteTypeInt = 0
	}

	voteComment := VoteComment{
		ID:        voted.ID,
		UserID:    voted.UserID,
		CommentID: voted.CommentID,
		VoteType:  voteTypeInt,
	}

	return c.Status(http.StatusOK).JSON(utils.ResponseSuccess{
		Status:  "success",
		Message: "Vote comment successfully",
		Data: struct {
			Vote interface{} `json:"vote"`
		}{
			Vote: voteComment,
		},
	})
}

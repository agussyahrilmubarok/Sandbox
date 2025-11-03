package rest

import "time"

// New thread request payload
type NewThread struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	Category string `json:"category"` // optional: default is general
	UserID   int64
}

// Thread tile response payload
type ThreadTile struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	Category      string    `json:"category"`
	CreatedAt     time.Time `json:"createdAt"`
	OwnerId       int64     `json:"ownerId"`
	UpVotesBy     []string  `json:"upVotesBy"`
	DownVotesBy   []string  `json:"downVotesBy"`
	TotalComments int       `json:"totalComments"`
}

// Thread detail response payload
type ThreadDetail struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Body        string          `json:"body"`
	Category    string          `json:"category"`
	CreatedAt   time.Time       `json:"createdAt"`
	Owner       UserProfile     `json:"owner"`
	UpVotesBy   []string        `json:"upVotesBy"`
	DownVotesBy []string        `json:"downVotesBy"`
	Comments    []CommentDetail `json:"comments"`
}

// New comment request payload
type NewComment struct {
	Content  string `json:"content" validate:"required"`
	UserID   int64
	ThreadID int64
}

// Comment detail response payload
type CommentDetail struct {
	ID          int64       `json:"id"`
	Content     string      `json:"content"`
	CreatedAt   time.Time   `json:"createdAt"`
	Owner       UserProfile `json:"owner"`
	UpVotesBy   []string    `json:"upVotesBy"`
	DownVotesBy []string    `json:"downVotesBy"`
}

// Vote thread response
type VoteThread struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"userId"`
	ThreadID int64 `json:"threadId"`
	VoteType int   `json:"voteType"`
}

// Vote comment response
type VoteComment struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"userId"`
	CommentID int64 `json:"commentId"`
	VoteType  int   `json:"voteType"`
}

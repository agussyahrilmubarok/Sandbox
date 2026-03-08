package rest

// LeaderBoards response payload
type LeaderBoard struct {
	User  UserProfile `json:"user"`
	Score int         `json:"score"`
}

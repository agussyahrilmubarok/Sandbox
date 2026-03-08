package domain

import "context"

//go:generate mockery --name=VoteRepository
type VoteRepository interface {
	SetVoteThread(ctx context.Context, voteThread *VoteThread) (*VoteThread, error)
	SetVoteComment(ctx context.Context, voteComment *VoteComment) (*VoteComment, error)
}

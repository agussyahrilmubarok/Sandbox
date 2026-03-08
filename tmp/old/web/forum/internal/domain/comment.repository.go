package domain

import "context"

//go:generate mockery --name=CommentRepository
type CommentRepository interface {
	FetchAll(ctx context.Context) ([]*Comment, error)
	FetchAllByUserID(ctx context.Context, ID int64) ([]*Comment, error)
	FetchAllByThreadID(ctx context.Context, ID int64) ([]*Comment, error)
	FetchByID(ctx context.Context, ID int64) (*Comment, error)
	Save(ctx context.Context, comment *Comment) (*Comment, error)
	Update(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, comment *Comment) error
	ExistByID(ctx context.Context, ID int64) (bool, error)
}

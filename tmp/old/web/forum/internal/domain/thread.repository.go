package domain

import "context"

//go:generate mockery --name=ThreadRepository
type ThreadRepository interface {
	FetchAll(ctx context.Context) ([]*Thread, error)
	FetchAllByUserID(ctx context.Context, ID int64) ([]*Thread, error)
	FetchByID(ctx context.Context, ID int64) (*Thread, error)
	Save(ctx context.Context, thread *Thread) (*Thread, error)
	Update(ctx context.Context, thread *Thread) (*Thread, error)
	Delete(ctx context.Context, thread *Thread) error
	ExistByID(ctx context.Context, ID int64) (bool, error)
}

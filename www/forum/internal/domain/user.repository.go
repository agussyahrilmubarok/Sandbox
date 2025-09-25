package domain

import "context"

//go:generate mockery --name=UserRepository
type UserRepository interface {
	FetchAll(ctx context.Context) ([]*User, error)
	FetchByID(ctx context.Context, ID int64) (*User, error)
	FetchByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, user *User) error
	ExistUsernameIgnoreCase(ctx context.Context, username string) (bool, error)
	ExistEmailIgnoreCase(ctx context.Context, email string) (bool, error)
}

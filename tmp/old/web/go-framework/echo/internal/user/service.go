package user

import "context"

//go:generate mockery --name=IService
type IService interface {
	SignUp(ctx context.Context, param SignUpRequest) (*UserReponse, error)
	SignIn(ctx context.Context, param SignInRequest) (*UserWithTokenResponse, error)
	FindAll(ctx context.Context) ([]UserReponse, error)
	FindByID(ctx context.Context, userID string) (*UserReponse, error)
	FindByEmail(ctx context.Context, userEmail string) (*UserReponse, error)
}

package auth

import (
	"context"

	entity "github.com/fauzanmh/olp-user/entity/microservice"
)

type AuthAdapter interface {
	CreateUser(ctx context.Context, req *entity.CreateUserRequest) (err error)
}

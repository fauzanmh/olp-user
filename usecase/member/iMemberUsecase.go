package member

import (
	"context"

	"github.com/fauzanmh/olp-user/schema/member"
)

type Usecase interface {
	Register(ctx context.Context, req *member.RegisterRequest) (err error)
}

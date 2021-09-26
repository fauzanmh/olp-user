package statistic

import (
	"context"

	"github.com/fauzanmh/olp-user/schema/statistic"
)

type Usecase interface {
	Get(ctx context.Context) (res statistic.GetStatisticResponse, err error)
}

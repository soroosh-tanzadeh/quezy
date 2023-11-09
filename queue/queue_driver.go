package queue

import (
	"context"
	"github.com/soroosh-tanzadeh/quezy/worker/contract"
)

type Queue interface {
	Push(ctx context.Context, job string, args ...interface{})
	Pop(ctx context.Context) (contract.Job, error)
	Size(ctx context.Context) int
	GetConnectionName(ctx context.Context) string
	// IsAvailable Ping Check if queue is available
	IsAvailable(ctx context.Context) bool
}

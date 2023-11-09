package drivers

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/soroosh-tanzadeh/quezy/worker/contract"
)

type RedisQueue struct {
	rdb *redis.Client
}

func (r *RedisQueue) Push(ctx context.Context, job string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisQueue) Pop(ctx context.Context) (contract.Job, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisQueue) Size(ctx context.Context) int {
	//TODO implement me
	panic("implement me")
}

func (r *RedisQueue) GetConnectionName(_ context.Context) string {
	return "redis"
}

func (r *RedisQueue) IsAvailable(ctx context.Context) bool {
	return r.rdb.Ping(ctx).Err() == nil
}

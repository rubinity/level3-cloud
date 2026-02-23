package auth

import (
	"context"
	// "os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct{ Client *redis.Client }

func NewRedis(endpoint string) *Redis {
	println("endpr", endpoint)
	addr := endpoint //!!!!later change to env
	// addr := os.Getenv("REDIS_ADDR")
	// if addr == "" {
	// 	addr = "localhost:6379"
	// }
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return &Redis{Client: rdb}
}

// key - token value namespace
func (r *Redis) SetJTI(ctx context.Context, key, namespace string, exp time.Time) error {
	return r.Client.Set(ctx, key, namespace, time.Until(exp)).Err()
}

func (r *Redis) DelJTI(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *Redis) GetNamespaceByJTI(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}


package momento_redis

import (
	"context"
	"time"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Get(ctx context.Context, key string) *StringCmd {
	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Close() error {
	// do nothing as this doesn't apply to Momento
	return nil
}

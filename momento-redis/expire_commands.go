package momento_redis

import (
	"context"
	"time"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ExpireTime(ctx context.Context, key string) *DurationCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PExpireTime(ctx context.Context, key string) *DurationCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

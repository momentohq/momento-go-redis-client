package momento_redis

import (
	"context"
	"time"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Append(ctx context.Context, key, value string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Decr(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) DecrBy(ctx context.Context, key string, decrement int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetSet(ctx context.Context, key string, value interface{}) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetDel(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Incr(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) IncrBy(ctx context.Context, key string, value int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) MGet(ctx context.Context, keys ...string) *SliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) MSet(ctx context.Context, values ...interface{}) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) MSetNX(ctx context.Context, values ...interface{}) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) StrLen(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

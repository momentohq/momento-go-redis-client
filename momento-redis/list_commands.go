package momento_redis

import (
	"context"
	"time"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *KeyValuesCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LCS(ctx context.Context, q *LCSQuery) *LCSCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LIndex(ctx context.Context, key string, index int64) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LLen(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LMPop(ctx context.Context, direction string, count int64, keys ...string) *KeyValuesCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPop(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPopCount(ctx context.Context, key string, count int) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPush(ctx context.Context, key string, values ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPop(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPopCount(ctx context.Context, key string, count int) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPopLPush(ctx context.Context, source, destination string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPush(ctx context.Context, key string, values ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) GetBit(ctx context.Context, key string, offset int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitPosSpan(ctx context.Context, key string, bit int8, start, end int64, span string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

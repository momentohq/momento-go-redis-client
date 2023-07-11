package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SCard(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SInter(ctx context.Context, keys ...string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SMembers(ctx context.Context, key string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SMembersMap(ctx context.Context, key string) *StringStructMapCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SPop(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SRandMember(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

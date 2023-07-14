package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) ZAdd(ctx context.Context, key string, members ...Z) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZAddLT(ctx context.Context, key string, members ...Z) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZAddGT(ctx context.Context, key string, members ...Z) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZCard(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZCount(ctx context.Context, key, min, max string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZInter(ctx context.Context, store *ZStore) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZMPop(ctx context.Context, order string, count int64, keys ...string) *ZSliceWithKeyCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRank(ctx context.Context, key, member string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRank(ctx context.Context, key, member string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZScore(ctx context.Context, key, member string) *FloatCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRandMember(ctx context.Context, key string, count int) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRandMemberWithScores(ctx context.Context, key string, count int) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZUnion(ctx context.Context, store ZStore) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZDiff(ctx context.Context, keys ...string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

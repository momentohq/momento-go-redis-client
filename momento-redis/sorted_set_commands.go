package momento_redis

import (
	"context"
	"math"
	"strconv"
	"strings"

	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) ZAdd(ctx context.Context, key string, members ...Z) *IntCmd {

	elements := make([]momento.SortedSetElement, 0, len(members))
	marshaller := &Marshaller{}
	for i := 0; i < len(members); i++ {
		member := members[i]
		// the member itself could be a string, int or other types
		val, err := marshaller.MarshalRedisValue(member.Member)
		if err != nil {
			panic(UnsupportedOperationError("Member type not supported " + err.Error()))
		}
		momentoElement := &momento.SortedSetElement{
			Score: member.Score,
			Value: momento.String(val),
		}
		elements = append(elements, *momentoElement)
	}

	resp := &IntCmd{}

	sortedSetPutResponse, err := m.client.SortedSetPutElements(ctx, &momento.SortedSetPutElementsRequest{
		CacheName: m.cacheName,
		SetName:   key,
		Elements:  elements,
	})

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch sortedSetPutResponse.(type) {
	case *responses.SortedSetPutElementsSuccess:
		// redis returns the number of sorted set elements that were successfully stored
		resp.SetVal(int64(len(elements)))
	case error:
		resp.SetErr(RedisError(err.Error()))
	}

	return resp
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
	return m.zRangeByScoreWithOrder(ctx, key, opt, momento.ASCENDING)
}

func (m *MomentoRedisClient) ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	return m.zRangeByScoreWithScoresWithOrder(ctx, key, opt, momento.ASCENDING)
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
	return m.zRangeByScoreWithOrder(ctx, key, opt, momento.DESCENDING)
}

func (m *MomentoRedisClient) ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	return m.zRangeByScoreWithScoresWithOrder(ctx, key, opt, momento.DESCENDING)
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

func (m *MomentoRedisClient) zRangeByScoreWithOrder(ctx context.Context, key string, opt *ZRangeBy, dir momento.SortedSetOrder) *StringSliceCmd {
	resp := &StringSliceCmd{}

	mZrange := validateFetchByScoreOptions(opt)

	sortedSetFetchByScoreRequest := &momento.SortedSetFetchByScoreRequest{
		CacheName: m.cacheName,
		SetName:   key,
		Count:     mZrange.Count,
		Offset:    mZrange.Offset,
		MinScore:  mZrange.MinScore,
		MaxScore:  mZrange.MaxScore,
		Order:     dir,
	}

	sSetFetch, err := m.client.SortedSetFetchByScore(ctx, sortedSetFetchByScoreRequest)

	if err != nil {
		resp.SetErr(err)
		return resp
	}

	switch r := sSetFetch.(type) {
	case *responses.SortedSetFetchHit:
		elements := r.ValueStringElements()
		var val = make([]string, 0, len(elements))
		for i := 0; i < len(elements); i++ {
			val = append(val, elements[i].Value)
		}
		resp.SetVal(val)
	case *responses.SortedSetFetchMiss:
		// empty array is default
	}
	return resp
}

func (m *MomentoRedisClient) zRangeByScoreWithScoresWithOrder(ctx context.Context, key string, opt *ZRangeBy, dir momento.SortedSetOrder) *ZSliceCmd {
	resp := &ZSliceCmd{}

	mZrange := validateFetchByScoreOptions(opt)

	sortedSetFetchByScoreRequest := &momento.SortedSetFetchByScoreRequest{
		CacheName: m.cacheName,
		SetName:   key,
		Count:     mZrange.Count,
		Offset:    mZrange.Offset,
		MinScore:  mZrange.MinScore,
		MaxScore:  mZrange.MaxScore,
		Order:     dir,
	}

	sSetFetch, err := m.client.SortedSetFetchByScore(ctx, sortedSetFetchByScoreRequest)

	if err != nil {
		resp.SetErr(err)
		return resp
	}

	switch r := sSetFetch.(type) {
	case *responses.SortedSetFetchHit:
		elements := r.ValueStringElements()
		var val = make([]Z, 0, len(elements))
		for i := 0; i < len(elements); i++ {
			redisElement := Z{
				Score:  elements[i].Score,
				Member: elements[i].Value,
			}
			val = append(val, redisElement)
		}
		resp.SetVal(val)
	case *responses.SortedSetFetchMiss:
		// empty array is default
	}

	return resp
}

type MomentoZRangeBy struct {
	MinScore *float64
	MaxScore *float64
	Offset   *uint32
	Count    *uint32
}

func validateFetchByScoreOptions(opt *ZRangeBy) MomentoZRangeBy {
	mZRange := &MomentoZRangeBy{}
	if opt.Count < 0 {
		panic(UnsupportedOperationError("Negative count is not supported by Momento"))
	}
	if opt.Count > math.MaxUint32 {
		panic(UnsupportedOperationError("Count exceeds max supported integer value by Momento"))
	}
	if opt.Offset < 0 {
		panic(UnsupportedOperationError("Negative offset is not supported by Momento"))
	}
	if opt.Offset > math.MaxUint32 {
		panic(UnsupportedOperationError("Offset exceeds max supported integer value by Momento"))
	}

	// this is yet another special case. Go-Redis doesn't specify Count as a pointer, so we cannot reference
	// it for a null check. Go-lang by default treats integers as 0. So a Redis customer can not provide the
	// count and still have all the elements in the sorted-set even if Go-Lang treats it as 0. This is because
	// Redis simply streams the entire command as a string for the params provided by the client and excludes
	// the default 0 count. In case of Momento, we don't want to send this default value as it will always
	// result in returning an empty set of values.
	if opt.Count != 0 {
		count := uint32(opt.Count)
		mZRange.Count = &count
	}

	offset := uint32(opt.Offset)
	mZRange.Offset = &offset

	// if client provided +inf, we don't pass anything to Momento as that's the default UnboundedMax
	// the length check is present because an empty string is treated as 0 and will be assigned as such later
	if opt.Min != "-inf" && len(opt.Min) > 0 {
		if strings.HasPrefix(opt.Min, "(") {
			panic(UnsupportedOperationError("Momento currently does not support exclusion of scores. Please omit ( from your" +
				" request and retry. Reach out to us at Discord https://discord.com/invite/3HkAKjUZGq) or e-mail us at " +
				"support@momentohq.com if you want such capabilities!"))
		}
		val, err := strconv.ParseFloat(opt.Min, 64)
		if err != nil {
			panic("Cannot parse string MinScore into float, MinScore: " + opt.Min)
		}
		mZRange.MinScore = &val
	}

	// if client provided +inf, we don't pass anything to Momento as that's the default UnboundedMax
	// the length check is present because an empty string is treated as 0 and will be assigned as such later
	if opt.Max != "+inf" && len(opt.Max) > 0 {
		if strings.HasPrefix(opt.Max, "(") {
			panic(UnsupportedOperationError("Momento currently does not support exclusion of scores. Please omit ( from your" +
				" request and retry. Reach out to us at Discord https://discord.com/invite/3HkAKjUZGq) or e-mail us at " +
				"support@momentohq.com if you want such capabilities!"))
		}
		val, err := strconv.ParseFloat(opt.Max, 64)
		if err != nil {
			panic("Cannot parse string MaxScore into float, MaxScore: " + opt.Max)
		}
		mZRange.MaxScore = &val
	}

	// Redis considers empty string to be 0
	if len(opt.Min) == 0 {
		score := float64(0)
		mZRange.MinScore = &score
	}

	// Redis considers empty string to be 0
	if len(opt.Max) == 0 {
		score := float64(0)
		mZRange.MaxScore = &score
	}

	return *mZRange
}

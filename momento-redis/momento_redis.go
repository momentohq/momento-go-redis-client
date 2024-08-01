package momento_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/momentohq/client-sdk-go/momento"
)

// MomentoRedisClient wrapper over momento cache client that provides Redis compatible APIs
type MomentoRedisClient struct {
	client    momento.CacheClient
	cacheName string
}

func NewMomentoRedisClient(cacheClient momento.CacheClient, cacheName string) *MomentoRedisClient {
	client := MomentoRedisClient{
		cacheName: cacheName,
		client:    cacheClient,
	}
	return &client
}

func (m *MomentoRedisClient) String() string {
	return fmt.Sprintf("Momento< cache:%s>", m.cacheName)
}

type UnsupportedOperationError string

func (e UnsupportedOperationError) Error() string { return string(e) }

type RedisError string

func (e RedisError) Error() string { return string(e) }

// RedisError implementing this No-Op method of go-redis interface automatically makes this a Redis
// error. From Redis v9, only a non-existing key explicitly gives a Redis.Nil error, and
// all other errors are bubbled up as a RedisError
func (e RedisError) RedisError() {}

// MomentoRedisCmdable is a Momento flavored interface that only exposes the APIs this compatibility client implements.
// Using this will cause compile time errors incase you use an API that isn't already implemented.
type MomentoRedisCmdable interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	TTL(ctx context.Context, key string) *redis.DurationCmd

	// sorted set commands
	ZAdd(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd

	// dictionary commands
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HGet(ctx context.Context, key string, field string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
	HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
	HLen(ctx context.Context, key string) *redis.IntCmd

	// list commands
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LLen(ctx context.Context, key string) *redis.IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
}

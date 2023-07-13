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

func (c *MomentoRedisClient) String() string {
	return fmt.Sprintf("Momento< cache:%s>", c.cacheName)
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
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	TTL(ctx context.Context, key string) *redis.DurationCmd
}

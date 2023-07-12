package momento_redis

import (
	"fmt"
	"github.com/momentohq/client-sdk-go/momento"
)

// MomentoRedisClient wrapper over momento cache client that provides Redis compatible APIs
type MomentoRedisClient struct {
	client    momento.CacheClient
	cacheName string
}

func NewMomentoRedisClient(cacheClient momento.CacheClient, cacheName string) (*MomentoRedisClient, error) {
	client := MomentoRedisClient{
		cacheName: cacheName,
		client:    cacheClient,
	}
	return &client, nil
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

package momento_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
)

type MomentoRedisClient struct {
	client    momento.CacheClient
	cacheName string
}

const AuthTokenEnvVariable string = "MOMENTO_AUTH_TOKEN"

func NewMomentoRedisClientWithDefaultCacheClient(cacheName string) (*MomentoRedisClient, error) {
	credentials, envErr := auth.NewEnvMomentoTokenProvider(AuthTokenEnvVariable)
	// fail fast if we can't fetch the credentials from env variable
	if envErr != nil {
		return nil, envErr
	}

	// default client with INFO logging capabilities and a 60-second default TTL for all keys
	mClient, err := momento.NewCacheClient(config.LaptopLatest(), credentials, 60*time.Second)
	if err != nil {
		return nil, err
	}

	// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
	_, err = mClient.CreateCache(context.Background(), &momento.CreateCacheRequest{
		CacheName: cacheName,
	})
	if err != nil {
		return nil, err
	}

	client := MomentoRedisClient{
		cacheName: cacheName,
		client:    mClient,
	}

	return &client, nil
}

func NewMomentoRedisClient(cacheClient momento.CacheClient, cacheName string) (*MomentoRedisClient, error) {
	// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
	_, err := cacheClient.CreateCache(context.Background(), &momento.CreateCacheRequest{
		CacheName: cacheName,
	})
	if err != nil {
		return nil, err
	}

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

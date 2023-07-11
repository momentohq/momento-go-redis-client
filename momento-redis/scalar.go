package momento_redis

import (
	"context"
	"time"

	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	"github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	get, err := m.client.Get(ctx, &momento.GetRequest{
		CacheName: m.cacheName,
		Key:       momento.String(key),
	})

	resp := &redis.StringCmd{}

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch val := get.(type) {
	case *responses.GetMiss:
		// redis.Nil indicates "key doesn't exist"
		resp.SetErr(redis.Nil)
	case *responses.GetHit:
		resp.SetVal(val.ValueString())
	case error:
		// all other RedisErrors in go-redis v8+ are treated under the same RedisError umbrella
		resp.SetErr(RedisError(err.Error()))
	}
	return resp
}

func (m *MomentoRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {

	momentoValue := panicIfNotSupportedArgs(expiration, value)

	resp := &redis.StatusCmd{}

	set, err := m.client.Set(ctx, &momento.SetRequest{
		CacheName: m.cacheName,
		Key:       momento.String(key),
		Value:     momentoValue,
		Ttl:       expiration,
	})

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch set.(type) {
	case *responses.SetSuccess:
		// redis OK response for a set success
		resp.SetVal("OK")
	case error:
		// all other RedisErrors in go-redis v8+ are treated under the same RedisError umbrella
		resp.SetErr(RedisError(err.Error()))
	}

	return resp
}

func (m *MomentoRedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {

	momentoValue := panicIfNotSupportedArgs(expiration, value)

	resp := &redis.BoolCmd{}

	setNX, err := m.client.SetIfNotExists(ctx, &momento.SetIfNotExistsRequest{
		CacheName: m.cacheName,
		Key:       momento.String(key),
		Value:     momentoValue,
		Ttl:       expiration,
	})

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch setNX.(type) {
	case *responses.SetIfNotExistsStored:
		resp.SetVal(true)
	case *responses.SetIfNotExistsNotStored:
		resp.SetVal(false)
	case error:
		// all other RedisErrors in go-redis v8+ are treated under the same RedisError umbrella
		resp.SetErr(RedisError(err.Error()))
	}
	return resp
}

func panicIfNotSupportedArgs(expiration time.Duration, value interface{}) momento.Value {
	if expiration == redis.KeepTTL {
		panic(UnsupportedOperationError("Momento does not support KeepTTL; please specify a TTL"))
	}

	var momentoValue momento.Value
	switch v := value.(type) {
	case []byte:
		momentoValue = momento.Bytes(v)
	case string:
		momentoValue = momento.String(v)
	default:
		panic(UnsupportedOperationError("Momento supports bytes and string for the value of the key in Set operation"))
	}

	return momentoValue
}

func (m *MomentoRedisClient) Close() error {
	// do nothing as this doesn't apply to Momento
	return nil
}

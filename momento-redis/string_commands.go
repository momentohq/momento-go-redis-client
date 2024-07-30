package momento_redis

import (
	"context"
	"time"

	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	"github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Append(ctx context.Context, key, value string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Decr(ctx context.Context, key string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) GetDel(ctx context.Context, key string) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	resp := &redis.SliceCmd{}

	var momentoKeys []momento.Key
	for _, key := range keys {
		momentoKeys = append(momentoKeys, momento.String(key))
	}

	getBatchResp, err := m.client.GetBatch(ctx, &momento.GetBatchRequest{
		CacheName: m.cacheName,
		Keys:      momentoKeys,
	})
	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch r := getBatchResp.(type) {
	case responses.GetBatchSuccess:
		// redis returns list of key values or nil if missing
		resultsMap := r.ValueMap()
		batchGetValues := []interface{}{}

		for _, key := range keys {
			if val, ok := resultsMap[key]; ok {
				batchGetValues = append(batchGetValues, val)
			} else {
				batchGetValues = append(batchGetValues, nil)
			}
		}
		resp.SetVal(batchGetValues)
	}
	return resp
}

func (m *MomentoRedisClient) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) StrLen(ctx context.Context, key string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

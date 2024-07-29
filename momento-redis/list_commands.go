package momento_redis

import (
	"context"
	"time"

	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	"github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *redis.KeyValuesCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LCS(ctx context.Context, q *redis.LCSQuery) *redis.LCSCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	resp := &redis.IntCmd{}

	listLengthResponse, err := m.client.ListLength(ctx, &momento.ListLengthRequest{
		CacheName: m.cacheName,
		ListName:  key,
	})
	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch r := listLengthResponse.(type) {
	case *responses.ListLengthHit:
		resp.SetVal(int64(r.Length()))
	case *responses.ListLengthMiss:
		// redis interprets a non-existing key as a list of length 0
		resp.SetVal(0)
	}

	return resp
}

func (m *MomentoRedisClient) LMPop(ctx context.Context, direction string, count int64, keys ...string) *redis.KeyValuesCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPop(ctx context.Context, key string) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	resp := &redis.StringSliceCmd{}
	startIndex := int32(start)
	stopIndex := int32(stop + 1) // Momento uses exclusive end, Redis uses inclusive end
	var listLength int32

	// Get the length of the list
	lengthResp, err := m.client.ListLength(ctx, &momento.ListLengthRequest{
		CacheName: m.cacheName,
		ListName:  key,
	})
	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}
	switch r := lengthResp.(type) {
	case *responses.ListLengthHit:
		listLength = int32(r.Length())
	case *responses.ListLengthMiss:
		listLength = int32(0)
	}

	// If indices are negative, convert them to positive indices
	if start < 0 {
		startIndex = listLength + startIndex
	}
	if stop < 0 {
		stopIndex = listLength + stopIndex
	}

	listFetchResponse, err := m.client.ListFetch(ctx, &momento.ListFetchRequest{
		CacheName:  m.cacheName,
		ListName:   key,
		StartIndex: &startIndex,
		EndIndex:   &stopIndex,
	})
	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch r := listFetchResponse.(type) {
	case *responses.ListFetchHit:
		resp.SetVal(r.ValueList())
	case *responses.ListFetchMiss:
		// redis returns empty array when list doesn't exist
		resp.SetVal([]string{})
	}

	return resp
}

func (m *MomentoRedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPop(ctx context.Context, key string) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func rPushElementsFromStrings(values []interface{}) ([]momento.Value, error) {
	var elements []momento.Value
	for i := 0; i < len(values); i++ {
		value, ok := values[i].(string)
		if !ok {
			return nil, UnsupportedOperationError("RPush received a non-string element while processing elements")
		}
		elements = append(elements, momento.String(value))
	}
	return elements, nil
}

func (m *MomentoRedisClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	resp := &redis.IntCmd{}
	var elements []momento.Value
	var err error

	switch values[0].(type) {
	case string:
		elements, err = rPushElementsFromStrings(values)
	default:
		err = UnsupportedOperationError("RPush has not implemented a way to handle the passed in values. Please pass in a series of strings to represent the elements to append to the list.")
	}
	if err != nil {
		resp.SetErr(err)
		return resp
	}

	listConcatBackResponse, err := m.client.ListConcatenateBack(ctx, &momento.ListConcatenateBackRequest{
		CacheName: m.cacheName,
		ListName:  key,
		Values:    elements,
	})
	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch r := listConcatBackResponse.(type) {
	case *responses.ListConcatenateBackSuccess:
		// redis returns the length of the list after the operation
		resp.SetVal(int64(r.ListLength()))
	}

	return resp
}

func (m *MomentoRedisClient) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LMove(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

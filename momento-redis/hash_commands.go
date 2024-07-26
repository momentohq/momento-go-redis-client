package momento_redis

import (
	"context"

	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	"github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	resp := &redis.StringCmd{}

	dictionaryGetFieldResponse, err := m.client.DictionaryGetField(ctx, &momento.DictionaryGetFieldRequest{
		CacheName:      m.cacheName,
		DictionaryName: key,
		Field:          momento.String(field),
	})

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch r := dictionaryGetFieldResponse.(type) {
	case *responses.DictionaryGetFieldHit:
		resp.SetVal(r.ValueString())
	case *responses.DictionaryGetFieldMiss:
		resp.SetErr(redis.Nil)
	}

	return resp
}

func (m *MomentoRedisClient) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	resp := &redis.MapStringStringCmd{}

	dictionaryFetchResponse, err := m.client.DictionaryFetch(ctx, &momento.DictionaryFetchRequest{
		CacheName:      m.cacheName,
		DictionaryName: key,
	})

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	// Miss return values differ from HGET but it matches what Redis client expects
	switch r := dictionaryFetchResponse.(type) {
	case *responses.DictionaryFetchHit:
		resp.SetVal(r.ValueMapStringString())
	case *responses.DictionaryFetchMiss:
		resp.SetErr(nil)
		resp.SetVal(map[string]string{})
	}

	return resp
}

func (m *MomentoRedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func hSetElementsFromStrings(values []interface{}) ([]momento.DictionaryElement, error) {
	var elements []momento.DictionaryElement
	for i := 0; i < len(values); i += 2 {
		field, ok := values[i].(string)
		if !ok {
			return nil, UnsupportedOperationError("HSet received a non-string field while processing elements")
		}
		value, ok := values[i+1].(string)
		if !ok {
			return nil, UnsupportedOperationError("HSet received a non-string value while processing elements")
		}
		elements = append(elements, momento.DictionaryElement{
			Field: momento.String(field),
			Value: momento.String(value),
		})
	}
	return elements, nil
}

func hSetElementsFromStringSlices(values []interface{}) ([]momento.DictionaryElement, error) {
	var elements []momento.DictionaryElement
	for _, slice := range values {
		sliceValues, ok := slice.([]string)
		if !ok {
			return nil, UnsupportedOperationError("HSet received a non-string slice while processing elements")
		}
		for i := 0; i < len(sliceValues); i += 2 {
			field := sliceValues[i]
			value := sliceValues[i+1]
			elements = append(elements, momento.DictionaryElement{
				Field: momento.String(field),
				Value: momento.String(value),
			})
		}
	}
	return elements, nil
}

func hSetElementsFromStringMaps(values []interface{}) ([]momento.DictionaryElement, error) {
	var elements []momento.DictionaryElement
	for _, value := range values {
		valuesMap, ok := value.(map[string]string)
		if !ok {
			return nil, UnsupportedOperationError("HSet received a non string-interface{} map while processing elements")
		}
		for field, value := range valuesMap {
			elements = append(elements, momento.DictionaryElement{
				Field: momento.String(field),
				Value: momento.String(value),
			})
		}
	}
	return elements, nil
}

func (m *MomentoRedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	resp := &redis.IntCmd{}
	var elements []momento.DictionaryElement
	var err error

	// Assuming each of the variadic arguments are of the same type
	switch values[0].(type) {
	case string:
		elements, err = hSetElementsFromStrings(values)
	case []string:
		elements, err = hSetElementsFromStringSlices(values)
	case map[string]string:
		elements, err = hSetElementsFromStringMaps(values)
	default:
		err = UnsupportedOperationError("HSet has not implemented a way to handle the passed in values. Please pass in a series of strings, []string, or map[string]string to represent the elements to add to the hash map.")
	}
	if err != nil {
		resp.SetErr(err)
		return resp
	}

	dictionarySetFieldsResponse, err := m.client.DictionarySetFields(ctx, &momento.DictionarySetFieldsRequest{
		CacheName:      m.cacheName,
		DictionaryName: key,
		Elements:       elements,
	})

	if err != nil {
		resp.SetErr(RedisError(err.Error()))
		return resp
	}

	switch dictionarySetFieldsResponse.(type) {
	case responses.DictionarySetFieldsResponse:
		// redis returns the number of dictionary elements that were successfully stored
		resp.SetVal(int64(len(elements)))
	case error:
		resp.SetErr(RedisError(err.Error()))
	}

	return resp
}

func (m *MomentoRedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HVals(ctx context.Context, key string) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HRandField(ctx context.Context, key string, count int) *redis.StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HRandFieldWithValues(ctx context.Context, key string, count int) *redis.KeyValueSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

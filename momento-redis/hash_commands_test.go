package momento_redis_test

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	momentoredis "github.com/momentohq/momento-go-redis-client/momento-redis"

	. "github.com/momentohq/momento-go-redis-client/momento-redis/test_helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("Dictionary methods", func() {
	sContext := NewSharedContext()
	BeforeEach(func() {
		cacheName := fmt.Sprintf("golang-redis-%s", uuid.NewString())
		switch sContext.UseRedis {
		case true:
			host := "127.0.0.1"
			port := "6379"
			sContext.Client = redis.NewClient(&redis.Options{
				Addr: host + ":" + port,
			})
		case false:
			credential, _ := auth.NewEnvMomentoTokenProvider(AuthTokenEnvVariable)
			mClient, _ := momento.NewCacheClient(config.LaptopLatest(), credential, 60*time.Second)
			sContext.MomentoClient = mClient
			// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
			_, createErr := sContext.CreateCache(sContext.Ctx, mClient, cacheName)
			if createErr != nil {
				panic("Failed to create cache with cache name " + cacheName + "\n" + createErr.Error())
			}

			sContext.Client = momentoredis.NewMomentoRedisClient(mClient, cacheName)
		}
		DeferCleanup(func() {
			if sContext.UseRedis {
				sContext.Client.FlushDB(sContext.Ctx)
			} else {
				_, deleteErr := sContext.DeleteCache(sContext.Ctx, sContext.MomentoClient, cacheName)
				if deleteErr != nil {
					panic("Failed to delete cache with cache name " + cacheName + "\n" + deleteErr.Error())
				}
			}
		})
	})

	var _ = Describe("Dictionary set fields", func() {
		It("Adds to a dictionary", func() {
			// Accepts elements as a series of strings
			resp := sContext.Client.HSet(sContext.Ctx, "dictionary", "string-1", "value-1", "string-2", "value-2")
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as a string literal
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", []string{"string-slice-1", "value-1", "string-slice-2", "value-2"})
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as a slice
			var values []string
			values = append(values, "string-slice-1", "value-1", "string-slice-2", "value-2")
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", values)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as multiple slices
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", []string{"string-slice-1", "value-1"}, []string{"string-slice-2", "value-2"})
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as map[string]string
			var stringValuesMap = make(map[string]string)
			stringValuesMap["string-map-1"] = "value-1"
			stringValuesMap["string-map-2"] = "value-2"
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", stringValuesMap)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as multiple maps
			var stringMap1 = make(map[string]string)
			var stringMap2 = make(map[string]string)
			stringMap1["multiple-maps-1"] = "value-1"
			stringMap2["multiple-maps-2"] = "value-2"
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", stringMap1, stringMap2)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			momentoElement1 := momento.DictionaryElement{
				Field: momento.String("momento-element-1"),
				Value: momento.String("value-1"),
			}
			momentoElement2 := momento.DictionaryElement{
				Field: momento.String("momento-element-2"),
				Value: momento.String("value-2"),
			}

			// Accepts elements as individual momento.DictionaryElements
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary-elements", momentoElement1, momentoElement2)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as slice literal of momento.DictionaryElement
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary-elements-slice-literal", []momento.DictionaryElement{momentoElement1, momentoElement2})
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as slice of momento.DictionaryElement
			var momentoElements []momento.DictionaryElement
			momentoElements = append(momentoElements, momentoElement1, momentoElement2)
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary-elements-slice", momentoElements)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as slices of momento.DictionaryElement
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary-elements-slices", []momento.DictionaryElement{momentoElement1}, []momento.DictionaryElement{momentoElement2})
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))
		})

		It("Adds to dictionary unsupported type and panics", func() {
			if sContext.UseRedis {
				return
			}
			type random struct {
				Field string
				Value string
			}
			random1 := random{
				Field: "random-1",
				Value: "value-1",
			}
			random2 := random{
				Field: "random-2",
				Value: "value-2",
			}
			resp := sContext.Client.HSet(sContext.Ctx, "dictionary", random1, random2)
			Expect(resp.Err()).ToNot(BeNil())
			Expect(resp.Err().Error()).To(ContainSubstring("HSet has not implemented a way to handle the passed in values. Please pass in a series of strings, map[string]string, []string, or []momento.DictionaryElement to represent the elements to add to the hash map."))
		})
	})
})

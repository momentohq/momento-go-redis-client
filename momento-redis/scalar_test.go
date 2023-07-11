package momento_redis_test

import (
	"fmt"
	"reflect"
	"time"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	momentoredis "github.com/momentohq/go-redis-client/momento-redis"
	. "github.com/momentohq/go-redis-client/momento-redis/test_helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("Scalar methods", func() {
	sContext := NewSharedContext()
	BeforeEach(func() {
		cacheName := "default_cache"
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
			sContext.CreateCache(sContext.Ctx, mClient, cacheName)
			sContext.Client, _ = momentoredis.NewMomentoRedisClient(mClient, cacheName)
		}
		DeferCleanup(func() {
			if !sContext.UseRedis {
				sContext.DeleteCache(sContext.Ctx, sContext.MomentoClient, cacheName)
			}
		})
	})

	var _ = Describe("Get and Set", func() {
		It("Sets string and Gets it", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			getResp := sContext.Client.Get(sContext.Ctx, "key")
			Expect(getResp.Val()).To(Equal("value"))
			Expect(getResp.Err()).To(BeNil())
		})

		It("Sets bytes and Gets it", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", []byte("value"), 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			getResp := sContext.Client.Get(sContext.Ctx, "key")
			Expect(getResp.Val()).To(Equal("value"))
			Expect(getResp.Err()).To(BeNil())
		})

		It("Key doesn't exist", func() {
			getResp := sContext.Client.Get(sContext.Ctx, "idontexist")
			Expect(getResp.Val()).To(Equal(""))
			Expect(getResp.Err()).To(Equal(redis.Nil))
		})

		It("Cache doesn't exist", func() {
			// cache only applies to Momento
			if sContext.UseRedis {
				return
			}
			momentoRedisNonExistentCache, _ := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
			resp := momentoRedisNonExistentCache.Set(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("NotFoundError: Cache not found"))
		})

		It("Keep ttl results in panic", func() {
			// KeepTTL is valid in redis
			if sContext.UseRedis {
				return
			}
			defer assertUnsupportedOperationPanic("Momento does not support KeepTTL; please specify a TTL")
			sContext.Client.Set(sContext.Ctx, "key", "value", redis.KeepTTL)
		})

		It("Unsupported value type results in panic", func() {
			// negative test only applicable to Momento
			if sContext.UseRedis {
				return
			}
			type random struct{}
			defer assertUnsupportedOperationPanic("Momento supports bytes and string for the value of the key in Set operation")
			sContext.Client.Set(sContext.Ctx, "key", random{}, 60*time.Second)

		})

		It("Invalid ttl results in error", func() {
			// negative test only applicable to Momento
			if sContext.UseRedis {
				return
			}
			// minimum ttl is 1 second
			resp := sContext.Client.Set(sContext.Ctx, "key", "value", 60*time.Nanosecond)
			Expect(resp.Err()).NotTo(BeNil())
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("InvalidArgumentError"))
		})
	})

	var _ = Describe("Set if not exists", func() {
		It("Set succeeds first time and Gets it", func() {
			setResp := sContext.Client.SetNX(sContext.Ctx, "randomKey", "value", 2*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(true))

			getResp := sContext.Client.Get(sContext.Ctx, "randomKey")
			Expect(getResp.Val()).To(Equal("value"))
			Expect(getResp.Err()).To(BeNil())

			setResp = sContext.Client.SetNX(sContext.Ctx, "randomKey", "value", 2*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(false))
		})

		It("Cache doesn't exist", func() {
			// cache only applies to Momento
			if sContext.UseRedis {
				return
			}
			momentoRedisNonExistentCache, _ := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
			resp := momentoRedisNonExistentCache.SetNX(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("NotFoundError: Cache not found"))
		})

		It("Keep ttl results in panic", func() {
			// KeepTTL is valid in redis
			if sContext.UseRedis {
				return
			}
			defer assertUnsupportedOperationPanic("Momento does not support KeepTTL; please specify a TTL")
			sContext.Client.SetNX(sContext.Ctx, "key", "value", redis.KeepTTL)
		})

		It("Unsupported value type results in panic", func() {
			// negative test only applicable to Momento
			if sContext.UseRedis {
				return
			}
			type random struct{}
			defer assertUnsupportedOperationPanic("Momento supports bytes and string for the value of the key in Set operation")
			sContext.Client.SetNX(sContext.Ctx, "key", random{}, 60*time.Second)

		})

		It("Invalid ttl results in error", func() {
			// negative test only applicable to Momento
			if sContext.UseRedis {
				return
			}
			// minimum ttl is 1 second
			resp := sContext.Client.SetNX(sContext.Ctx, "key", "value", 60*time.Nanosecond)
			Expect(resp.Err()).NotTo(BeNil())
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("InvalidArgumentError"))
		})
	})

	var _ = Describe("Type check tests", func() {
		It("Momento Redis error is of type Go-Redis Redis error", func() {
			if sContext.UseRedis {
				return
			}
			err := fmt.Errorf("A Random Error")
			momentoRedisError := momentoredis.RedisError("MomentoRedis error")

			// get type of the Redis Error interface by go-lang
			t := reflect.TypeOf(new(redis.Error)).Elem()

			// check our custom Redis errors are of the type
			Expect(reflect.TypeOf(momentoRedisError).Implements(t)).To(Equal(true))
			// other errors are not of the go-lang Redis error type
			Expect(reflect.TypeOf(err).Implements(t)).To(Equal(false))
		})
	})
})

func assertUnsupportedOperationPanic(message string) {
	r := recover()
	if r == nil {
		Fail("The code did not panic")
	} else {
		errMsg, ok := r.(momentoredis.UnsupportedOperationError)
		if ok {
			Expect(errMsg.Error()).To(Equal(message))
		} else {
			Fail("Unknown panic occured")
		}
	}
}

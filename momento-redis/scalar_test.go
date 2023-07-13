package momento_redis_test

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	momentoredis "github.com/momentohq/momento-go-redis-client/momento-redis"
	. "github.com/momentohq/momento-go-redis-client/momento-redis/test_helpers"
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
			sContext.Client = momentoredis.NewMomentoRedisClient(mClient, cacheName)
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

		It("Sets int types and Gets it", func() {

			setResp := sContext.Client.Set(sContext.Ctx, "key", 4, 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			getResp := sContext.Client.Get(sContext.Ctx, "key")
			Expect(getResp.Val()).To(Equal(strconv.FormatInt(4, 10)))
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

		It("Sets time and Gets it", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", 10*time.Second, 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			getResp := sContext.Client.Get(sContext.Ctx, "key")
			Expect(getResp.Err()).To(BeNil())
			// RESP converts 10 seconds to nanoseconds which is 10 billion
			Expect(getResp.Val()).To(Equal("10000000000"))
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
			momentoRedisNonExistentCache := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
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
			momentoRedisNonExistentCache := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
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

	var _ = Describe("Delete key", func() {
		It("Deletes key successfully", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			getResp := sContext.Client.Get(sContext.Ctx, "key")
			Expect(getResp.Val()).To(Equal("value"))
			Expect(getResp.Err()).To(BeNil())

			delResp := sContext.Client.Del(sContext.Ctx, "key")
			Expect(delResp.Val()).To(Equal(int64(1)))
			Expect(delResp.Err()).To(BeNil())
		})

		It("Key doesn't exist", func() {

			delResp := sContext.Client.Del(sContext.Ctx, "idonatexist")
			if sContext.UseRedis {
				Expect(delResp.Val()).To(Equal(int64(0)))
			} else {
				Expect(delResp.Val()).To(Equal(int64(1)))

			}
			Expect(delResp.Err()).To(BeNil())
		})

		It("Cache doesn't exist", func() {
			// cache only applies to Momento
			if sContext.UseRedis {
				return
			}
			momentoRedisNonExistentCache := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
			resp := momentoRedisNonExistentCache.Del(sContext.Ctx, "key")
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("NotFoundError: Cache not found"))
		})

		It("More than one key not supported", func() {
			if sContext.UseRedis {
				// redis does support more than one key deletion at a time
				return
			}
			defer assertUnsupportedOperationPanic("Momento supports deletion of a single key at a time")
			sContext.Client.Del(sContext.Ctx, "idontexist", "anotherKey")
		})
	})

	var _ = Describe("Expire key", func() {
		It("expire or update ttl set successfully", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			expResp := sContext.Client.Expire(sContext.Ctx, "key", 10*time.Second)
			Expect(expResp.Err()).To(BeNil())
			Expect(expResp.Val()).To(Equal(true))

		})

		It("expire or update ttl key doesn't exist", func() {
			expResp := sContext.Client.Expire(sContext.Ctx, "IDontExist", 10*time.Second)
			Expect(expResp.Err()).To(BeNil())
			Expect(expResp.Val()).To(Equal(false))
		})

		It("expire or update ttl with 0 key exists and gets deleted", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			expResp := sContext.Client.Expire(sContext.Ctx, "key", 0*time.Second)
			Expect(expResp.Err()).To(BeNil())
			Expect(expResp.Val()).To(Equal(true))

			getResp := sContext.Client.Get(sContext.Ctx, "key")
			Expect(getResp.Val()).To(Equal(""))
		})

		It("expire or update ttl with 0 key doesnt exist and gets deleted", func() {

			expResp := sContext.Client.Expire(sContext.Ctx, "idontexist", 0*time.Second)
			Expect(expResp.Err()).To(BeNil())
			if sContext.UseRedis {
				// expire with 0 triggers delete and redis knows the key didn't existed
				// so says false
				Expect(expResp.Val()).To(Equal(false))
			} else {
				// expire with 0 triggers delete and momento doesn't know the key didn't existed
				// so says true
				Expect(expResp.Val()).To(Equal(true))
			}
		})

		It("expire or update ttl with <0 key doesnt exist and gets deleted", func() {

			expResp := sContext.Client.Expire(sContext.Ctx, "idontexist", -5*time.Second)
			Expect(expResp.Err()).To(BeNil())
			if sContext.UseRedis {
				// expire with 0 triggers delete and redis knows the key didn't existed
				// so says false
				Expect(expResp.Val()).To(Equal(false))
			} else {
				// expire with 0 triggers delete and momento doesn't know the key didn't existed
				// so says true
				Expect(expResp.Val()).To(Equal(true))
			}
		})

		It("Cache doesn't exist", func() {
			// cache only applies to Momento
			if sContext.UseRedis {
				return
			}
			momentoRedisNonExistentCache := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
			resp := momentoRedisNonExistentCache.Expire(sContext.Ctx, "key", 5*time.Second)
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("NotFoundError: Cache not found"))
		})
	})

	var _ = Describe("Get TTL", func() {

		It("fetches TTL successfully", func() {
			setResp := sContext.Client.Set(sContext.Ctx, "key", "value", 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal("OK"))

			// redis is connected locally so it's way too fast for the
			// ttl to change at all and falls within their expected error margin
			time.Sleep(1 * time.Second)
			ttl := sContext.Client.TTL(sContext.Ctx, "key")
			Expect(ttl.Err()).To(BeNil())
			Expect(ttl.Val().Seconds() < 60).To(BeTrue())
			// just adding enough buffer for tests to not be flaky
			Expect(ttl.Val().Seconds() > 50).To(BeTrue())
		})

		It("key doesn't exist", func() {
			ttl := sContext.Client.TTL(sContext.Ctx, "IDontExist")
			Expect(ttl.Err()).To(BeNil())
			Expect(ttl.Val()).To(Equal(time.Duration(-2)))
		})

		It("Cache doesn't exist", func() {
			// cache only applies to Momento
			if sContext.UseRedis {
				return
			}
			momentoRedisNonExistentCache := momentoredis.NewMomentoRedisClient(sContext.MomentoClient, "NonExistent")
			resp := momentoRedisNonExistentCache.TTL(sContext.Ctx, "key")
			Expect(resp.Err()).To(BeAssignableToTypeOf(momentoredis.RedisError("")))
			Expect(resp.Err().Error()).To(ContainSubstring("NotFoundError: Cache not found"))
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

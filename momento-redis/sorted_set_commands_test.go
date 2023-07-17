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

var _ = Describe("Sorted Set methods", func() {
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

	var _ = Describe("Sorted set add", func() {
		It("Adds to sorted set", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sset", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))
		})

		It("Adds to sorted set unsupported type and panics", func() {
			if sContext.UseRedis {
				return
			}
			type random struct{}
			member1 := &redis.Z{
				Member: "member1",
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: &random{},
				Score:  100,
			}
			defer assertUnsupportedOperationPanic("Member type not supported")
			sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)

		})
	})

	var _ = Describe("Sorted set fetch by score", func() {
		It("Fetch by score happy case", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))
			zRange := &redis.ZRangeBy{
				Min:   "0",
				Max:   "100",
				Count: int64(2),
			}
			zResp := sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(2))
			Expect(zResp.Val()[0]).To(Equal(member1.Member))
			Expect(zResp.Val()[1]).To(Equal(member2.Member))

			// try with offset
			zRange = &redis.ZRangeBy{
				Min:    "0",
				Max:    "100",
				Count:  int64(2),
				Offset: int64(1),
			}
			zResp = sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)
			Expect(len(zResp.Val())).To(Equal(1))
			Expect(zResp.Val()[0]).To(Equal(member2.Member))
		})

		It("Fetch by score with exclusion rule provided Min panics", func() {
			if sContext.UseRedis {
				// redis supports exclusion
				return
			}
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))

			zRange := &redis.ZRangeBy{
				Min: "(98",
				Max: "100",
			}
			defer assertUnsupportedOperationPanic("Momento currently does not support exclusion of scores")
			sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)
		})

		It("Fetch by score with exclusion rule provided Max panics", func() {
			if sContext.UseRedis {
				// redis supports exclusion
				return
			}
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))

			zRange := &redis.ZRangeBy{
				Min: "98",
				Max: "(100",
			}
			defer assertUnsupportedOperationPanic("Momento currently does not support exclusion of scores")
			sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

		})

		It("Fetch by score no Min", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			member3 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  -5,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2, *member3)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(3)))
			// Min will considered 0 since it's empty string and member3 won't be included
			zRange := &redis.ZRangeBy{
				Min: "",
				Max: "100",
			}
			zResp := sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(2))
			Expect(zResp.Val()[0]).To(Equal(member1.Member))
			Expect(zResp.Val()[1]).To(Equal(member2.Member))

			// let's also test without a Min altogether
			// // Min will considered 0 since it's not present and member3 won't be included
			zRange = &redis.ZRangeBy{
				Max: "99",
			}
			zResp = sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(1))
			Expect(zResp.Val()[0]).To(Equal(member1.Member))
		})

		It("Fetch by score no Max", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			member3 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  -6,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2, *member3)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(3)))

			zRange := &redis.ZRangeBy{
				Min: "5",
			}
			// since no Max is provided, it will be considered 0 as it's an empty string
			zResp := sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(0))

			// let's also test without a Max altogether.
			// since no Max is provided, it will be considered 0 as it's an empty string; therefore,
			// the member will score -6 will now be included
			zRange = &redis.ZRangeBy{
				Min: "-10",
			}
			zResp = sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(1))
			Expect(zResp.Val()[0]).To(Equal(member3.Member))
		})

		It("Fetch by score no min and no max empty result", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))
			zRange := &redis.ZRangeBy{}
			zResp := sContext.Client.ZRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(0))
		})

		It("Fetch by score reverse order", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))
			zRange := &redis.ZRangeBy{
				Min: "0",
				Max: "100",
			}
			zResp := sContext.Client.ZRevRangeByScore(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(2))
			// first element in result is member 2 with score 100
			Expect(zResp.Val()[0]).To(Equal(member2.Member))
			Expect(zResp.Val()[1]).To(Equal(member1.Member))
		})

		It("Fetch by score with score", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))
			zRange := &redis.ZRangeBy{
				Min:   "0",
				Max:   "100",
				Count: int64(2),
			}
			zResp := sContext.Client.ZRangeByScoreWithScores(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(2))
			Expect(zResp.Val()[0].Member).To(Equal(member1.Member))
			Expect(zResp.Val()[0].Score).To(Equal(member1.Score))
			Expect(zResp.Val()[1].Member).To(Equal(member2.Member))
			Expect(zResp.Val()[1].Score).To(Equal(member2.Score))
		})

		It("Fetch by score with score reverse order", func() {
			member1 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  98.2,
			}
			member2 := &redis.Z{
				Member: fmt.Sprintf("member-%s", uuid.NewString()),
				Score:  100,
			}
			setResp := sContext.Client.ZAdd(sContext.Ctx, "sortedSet", *member1, *member2)
			Expect(setResp.Err()).To(BeNil())
			Expect(setResp.Val()).To(Equal(int64(2)))
			zRange := &redis.ZRangeBy{
				Min:   "0",
				Max:   "100",
				Count: int64(2),
			}
			zResp := sContext.Client.ZRevRangeByScoreWithScores(sContext.Ctx, "sortedSet", zRange)

			Expect(zResp.Err()).To(BeNil())
			Expect(len(zResp.Val())).To(Equal(2))
			Expect(zResp.Val()[1].Member).To(Equal(member1.Member))
			Expect(zResp.Val()[1].Score).To(Equal(member1.Score))
			Expect(zResp.Val()[0].Member).To(Equal(member2.Member))
			Expect(zResp.Val()[0].Score).To(Equal(member2.Score))
		})
	})

})

package helpers

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/config/logger/momento_default_logger"
	"github.com/momentohq/client-sdk-go/responses"

	"github.com/momentohq/client-sdk-go/momento"
	momentoredis "github.com/momentohq/momento-go-redis-client/momento-redis"
	"github.com/redis/go-redis/v9"
)

type SharedContext struct {
	// If set to true using command line argument, the same set of tests will be run against Redis on port 6379
	UseRedis bool

	// Client the wrapper Client implements all methods of the Cmdable interface supporting all Redis commands.
	// Note that this doesn't include commands exposed by ClusterClient, which isn't applicable if you're
	// using Momento as we provide resource-level isolation and have no notion of a cluster exposed to our customers.
	// This type declaration here serves as a validation and compile-time errors will occur if we do not implement
	// any particular API exposed the go-redis Client
	Client        redis.Cmdable
	MomentoClient momento.CacheClient
	Ctx           context.Context
	CacheName     string
}

const AuthTokenEnvVariable string = "TEST_AUTH_TOKEN"

var useRedis bool

// This method is necessary so that all the Go related flags are initialized/parsed along with our custom
// defined ones. If we don't call testing.Init() at the start of the test suite and try to parse our flags,
// Ginkgo will complain about it saying our custom flag was defined but not provided even if we do provide it.
var _ = func() bool {
	flag.BoolVar(&useRedis, "UseRedis", false, "Whether we want to run the tests using Momento or Redis")
	testing.Init()
	return true
}()

func NewSharedContext() SharedContext {
	flag.Parse()
	shared := SharedContext{}
	shared.UseRedis = useRedis
	shared.Ctx = context.Background()

	shared.CacheName = fmt.Sprintf("golang-redis-%s", uuid.NewString())
	switch shared.UseRedis {
	case true:
		host := "127.0.0.1"
		port := "6379"
		shared.Client = redis.NewClient(&redis.Options{
			Addr: host + ":" + port,
		})
	case false:
		momentoLoggerFactory := momento_default_logger.NewDefaultMomentoLoggerFactory(momento_default_logger.WARN)
		credential, err := auth.NewEnvMomentoTokenProvider(AuthTokenEnvVariable)
		if err != nil {
			panic("Failed to create testing momento credential provider\n" + err.Error())
		}
		mClient, err := momento.NewCacheClient(config.LaptopLatestWithLogger(momentoLoggerFactory), credential, 60*time.Second)
		if err != nil {
			panic("Failed to create testing momento client\n" + err.Error())
		}
		shared.MomentoClient = mClient
		// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
		_, createErr := shared.CreateCache(shared.Ctx, mClient, shared.CacheName)
		if createErr != nil {
			panic("Failed to create cache with cache name " + shared.CacheName + "\n" + createErr.Error())
		}
		shared.Client = momentoredis.NewMomentoRedisClient(mClient, shared.CacheName)
	}
	return shared
}

func (SharedContext) CreateCache(ctx context.Context, client momento.CacheClient, cacheName string) (responses.CreateCacheResponse, error) {
	return client.CreateCache(ctx, &momento.CreateCacheRequest{
		CacheName: cacheName,
	})
}

func (SharedContext) DeleteCache(ctx context.Context, client momento.CacheClient, cacheName string) (responses.DeleteCacheResponse, error) {
	return client.DeleteCache(ctx, &momento.DeleteCacheRequest{
		CacheName: cacheName,
	})
}

func (shared SharedContext) Close() {
	if shared.UseRedis {
		shared.Client.FlushDB(shared.Ctx)
	} else {
		_, deleteErr := shared.DeleteCache(shared.Ctx, shared.MomentoClient, shared.CacheName)
		if deleteErr != nil {
			panic("Failed to delete cache with cache name " + shared.CacheName + "\n" + deleteErr.Error())
		}
	}
}

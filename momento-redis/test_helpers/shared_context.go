package helpers

import (
	"context"
	"flag"
	"testing"

	"github.com/momentohq/client-sdk-go/momento"
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
	return shared
}

func (SharedContext) CreateCache(ctx context.Context, client momento.CacheClient, cacheName string) {
	client.CreateCache(ctx, &momento.CreateCacheRequest{
		CacheName: cacheName,
	})
}

func (SharedContext) DeleteCache(ctx context.Context, client momento.CacheClient, cacheName string) {
	client.DeleteCache(ctx, &momento.DeleteCacheRequest{CacheName: cacheName})
}

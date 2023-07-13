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

// required so that some Flags are autopopulated by testing without which Ginkgo complains
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

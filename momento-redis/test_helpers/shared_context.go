package helpers

import (
	"context"
	"flag"
	"testing"

	. "github.com/onsi/ginkgo/v2"

	momentoredis "github.com/momento-redis/go-redis-client/momento-redis"
	"github.com/redis/go-redis/v9"
)

type SharedContext struct {
	// If set to true using command line argument, the same set of tests will be run against Redis on port 6379
	UseRedis bool

	// Client the wrapper Client implements all methods of the Cmdable interface supporting all Redis commands.
	// Note that this doesn't include commands exposed by ClusterClient, which isn't applicable if you're
	// using Momento as we provide resource-level isolation and have no notion of a cluster exposed to our customers.
	// This type declaration here serves as a validation and compile-time errors will occur if we do not implement
	// any particular AP exposed the go-redis Client
	Client redis.Cmdable

	Ctx context.Context
}

// required so that some Flags are autopopulated by testing without which Ginkgo complains
var _ = func() bool {
	testing.Init()
	return true
}()

var SContext = NewSharedContext()

func NewSharedContext() SharedContext {
	shared := SharedContext{}
	setupFlags(&shared)
	shared.Ctx = context.Background()
	return shared
}

func setupFlags(shared *SharedContext) {
	flag.BoolVar(&shared.UseRedis, "UseRedis", false, "Whether we want to run the tests using Momento or Redis")
	flag.Parse()
}

var _ = BeforeSuite(func() {
	switch SContext.UseRedis {
	case true:
		SContext.Client = redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		})
	case false:
		SContext.Client, _ = momentoredis.NewMomentoRedisClientWithDefaultCacheClient("default_cache")
	}
})

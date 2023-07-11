package momento_redis_test

import (
	"fmt"

	. "github.com/momento-redis/go-redis-client/momento-redis/test_helpers"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Unsupported operations", func() {
	It("errors on get for not supported", func() {
		fmt.Printf("Value for use redis is %t\n", SContext.UseRedis)
		defer assertPanic()
		SContext.Client.Get(SContext.Ctx, "key")
	})
	It("errors on set for not supported", func() {
		defer assertPanic()
		SContext.Client.Set(SContext.Ctx, "key", "value", 1)
	})
})

func assertPanic() {
	if r := recover(); r == nil {
		Fail("The code did not panic")
	}
}

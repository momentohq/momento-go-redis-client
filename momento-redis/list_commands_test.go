package momento_redis_test

import (
	. "github.com/momentohq/momento-go-redis-client/momento-redis/test_helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dictionary methods", func() {
	var sContext SharedContext
	BeforeEach(func() {
		sContext = NewSharedContext()
		DeferCleanup(func() {
			sContext.Close()
		})
	})

	Describe("List concatenate back", func() {
		It("Adds strings to a list", func() {
			resp := sContext.Client.RPush(sContext.Ctx, "list", "value-1", "value-2")
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))
		})

		It("Adding non-strings to a list returns unsupported error", func() {
			if sContext.UseRedis {
				return
			}
			type random struct {
				Value string
			}
			random1 := random{
				Value: "value-1",
			}
			random2 := random{
				Value: "value-2",
			}
			resp := sContext.Client.RPush(sContext.Ctx, "list", random1, random2)
			Expect(resp.Err()).ToNot(BeNil())
			Expect(resp.Err().Error()).To(ContainSubstring("RPush has not implemented a way to handle the passed in values. Please pass in a series of strings to represent the elements to append to the list."))
		})
	})
})

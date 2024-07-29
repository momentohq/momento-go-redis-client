package momento_redis_test

import (
	"fmt"

	"github.com/google/uuid"
	. "github.com/momentohq/momento-go-redis-client/momento-redis/test_helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func newListName() string {
	return fmt.Sprintf("list-%s", uuid.NewString())
}

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
			listName := newListName()
			resp := sContext.Client.RPush(sContext.Ctx, listName, "value-1", "value-2")
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
			listName := newListName()
			resp := sContext.Client.RPush(sContext.Ctx, listName, random1, random2)
			Expect(resp.Err()).ToNot(BeNil())
			Expect(resp.Err().Error()).To(ContainSubstring("RPush has not implemented a way to handle the passed in values. Please pass in a series of strings to represent the elements to append to the list."))
		})
	})

	Describe("List length", func() {
		It("Gets length of a list", func() {
			listName := newListName()

			// Length of empty list should be zero
			resp := sContext.Client.LLen(sContext.Ctx, listName)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(0)))

			// Add some values to the list
			resp = sContext.Client.RPush(sContext.Ctx, listName, "value-1", "value-2")
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Length of existing list should be 2
			resp = sContext.Client.LLen(sContext.Ctx, listName)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))
		})

		It("Returns error when trying to get length of non-list item", func() {
			listName := newListName()

			// Set scalar item
			setResp := sContext.Client.Set(sContext.Ctx, listName, "value", 0)
			Expect(setResp.Err()).To(BeNil())

			// Get list length of scalar item should fail
			resp := sContext.Client.LLen(sContext.Ctx, listName)
			Expect(resp.Err()).ToNot(BeNil())
			if sContext.UseRedis {
				Expect(resp.Err().Error()).To(ContainSubstring("WRONGTYPE Operation against a key holding the wrong kind of value"))
			} else {
				Expect(resp.Err().Error()).To(ContainSubstring("FailedPreconditionError: The stored type for this key is not compatible with this operation"))
			}
		})
	})

	Describe("List fetch", func() {
		It("Gets elements given various ranges", func() {
			listName := newListName()

			// Empty list should return empty array
			resp := sContext.Client.LRange(sContext.Ctx, listName, 0, 10)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]string{}))

			// Add some values to the list
			pushResp := sContext.Client.RPush(sContext.Ctx, listName, "value-1", "value-2", "value-3", "value-4")
			Expect(pushResp.Err()).To(BeNil())
			Expect(pushResp.Val()).To(Equal(int64(4)))

			// Fetch all elements using exact range
			resp = sContext.Client.LRange(sContext.Ctx, listName, 0, 3)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]string{"value-1", "value-2", "value-3", "value-4"}))

			// Fetch using stop index > end of list
			resp = sContext.Client.LRange(sContext.Ctx, listName, 0, 10)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]string{"value-1", "value-2", "value-3", "value-4"}))

			// Fetch using start index > end of list
			resp = sContext.Client.LRange(sContext.Ctx, listName, 10, 20)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]string{}))

			// Fetch using start index > stop index
			resp = sContext.Client.LRange(sContext.Ctx, listName, 2, 1)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]string{}))

			// Fetch using negative range
			resp = sContext.Client.LRange(sContext.Ctx, listName, -2, -1)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]string{"value-3", "value-4"}))
		})
	})
})

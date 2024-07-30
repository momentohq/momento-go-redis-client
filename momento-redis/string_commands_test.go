package momento_redis_test

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	. "github.com/momentohq/momento-go-redis-client/momento-redis/test_helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func newKey() string {
	return fmt.Sprintf("key-%s", uuid.NewString())
}

func newValue() string {
	return fmt.Sprintf("value-%s", uuid.NewString())
}

var _ = Describe("Dictionary methods", func() {
	var sContext SharedContext
	BeforeEach(func() {
		sContext = NewSharedContext()
		DeferCleanup(func() {
			sContext.Close()
		})
	})

	Describe("Batch get", func() {
		It("Gets multiple keys", func() {
			key1 := newKey()
			key2 := newKey()
			value1 := newValue()
			value2 := newValue()

			// Add the values
			setResp := sContext.Client.Set(sContext.Ctx, key1, value1, 60*time.Second)
			Expect(setResp.Err()).To(BeNil())
			setResp = sContext.Client.Set(sContext.Ctx, key2, value2, 60*time.Second)
			Expect(setResp.Err()).To(BeNil())

			// Get all existing values
			resp := sContext.Client.MGet(sContext.Ctx, key1, key2)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]interface{}{value1, value2}))

			// Get some nonexistent values mixed with existing values
			resp = sContext.Client.MGet(sContext.Ctx, key1, "nonexistent", key2)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal([]interface{}{value1, nil, value2}))
		})
	})
})

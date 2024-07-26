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

	Describe("Dictionary set fields", func() {
		It("Adds to a dictionary", func() {
			// Accepts elements as a series of strings
			resp := sContext.Client.HSet(sContext.Ctx, "dictionary", "string-1", "value-1", "string-2", "value-2")
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as a string literal
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", []string{"string-slice-literal-1", "value-1", "string-slice-literal-2", "value-2"})
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as a slice
			var values []string
			values = append(values, "string-slice-1", "value-1", "string-slice-2", "value-2")
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", values)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Accepts elements as map[string]string
			var stringValuesMap = make(map[string]string)
			stringValuesMap["string-map-1"] = "value-1"
			stringValuesMap["string-map-2"] = "value-2"
			resp = sContext.Client.HSet(sContext.Ctx, "dictionary", stringValuesMap)
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))
		})

		It("Adds to dictionary unsupported type and returns error", func() {
			if sContext.UseRedis {
				return
			}
			type random struct {
				Field string
				Value string
			}
			random1 := random{
				Field: "random-1",
				Value: "value-1",
			}
			random2 := random{
				Field: "random-2",
				Value: "value-2",
			}
			resp := sContext.Client.HSet(sContext.Ctx, "dictionary", random1, random2)
			Expect(resp.Err()).ToNot(BeNil())
			Expect(resp.Err().Error()).To(ContainSubstring("HSet has not implemented a way to handle the passed in values. Please pass in a series of strings, []string, or map[string]string to represent the elements to add to the hash map."))
		})
	})

	Describe("Dictionary get fields", func() {
		It("Gets a single value from a dictionary", func() {
			// Add some elements to the dictionary
			resp := sContext.Client.HSet(sContext.Ctx, "dictionary", "string-1", "value-1", "string-2", "value-2")
			Expect(resp.Err()).To(BeNil())
			Expect(resp.Val()).To(Equal(int64(2)))

			// Get existing elements -> cache hit
			getResp := sContext.Client.HGet(sContext.Ctx, "dictionary", "string-1")
			Expect(getResp.Err()).To(BeNil())
			Expect(getResp.Val()).To(Equal("value-1"))

			getResp = sContext.Client.HGet(sContext.Ctx, "dictionary", "string-2")
			Expect(getResp.Err()).To(BeNil())
			Expect(getResp.Val()).To(Equal("value-2"))

			// Get nonexistent element -> cache miss
			getResp = sContext.Client.HGet(sContext.Ctx, "dictionary", "string-3")
			Expect(getResp.Err()).To(Not(BeNil()))
		})
	})
})

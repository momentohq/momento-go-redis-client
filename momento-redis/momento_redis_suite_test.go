package momento_redis_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMomentoRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MomentoRedis Suite")
}

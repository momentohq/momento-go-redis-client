package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Pipeline() Pipeliner {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) TxPipeline() Pipeliner {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

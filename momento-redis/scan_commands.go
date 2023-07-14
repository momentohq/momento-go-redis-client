package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) ClusterMyShardID(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterSlots(ctx context.Context) *ClusterSlotsCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterShards(ctx context.Context) *ClusterShardsCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterLinks(ctx context.Context) *ClusterLinksCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterNodes(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterMeet(ctx context.Context, host, port string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterForget(ctx context.Context, nodeID string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterResetSoft(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterResetHard(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterInfo(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterKeySlot(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterCountKeysInSlot(ctx context.Context, slot int) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterDelSlots(ctx context.Context, slots ...int) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterDelSlotsRange(ctx context.Context, min, max int) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterSaveConfig(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterFailover(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterAddSlots(ctx context.Context, slots ...int) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClusterAddSlotsRange(ctx context.Context, min, max int) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

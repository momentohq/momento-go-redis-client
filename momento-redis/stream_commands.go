package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) XAdd(ctx context.Context, a *XAddArgs) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XDel(ctx context.Context, stream string, ids ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XLen(ctx context.Context, stream string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XGroupDestroy(ctx context.Context, stream, group string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XPending(ctx context.Context, stream, group string) *XPendingCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XTrimMinID(ctx context.Context, key string, minID string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XInfoStream(ctx context.Context, key string) *XInfoStreamCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

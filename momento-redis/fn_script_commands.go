package momento_redis

import (
	"context"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ScriptFlush(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ScriptKill(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ScriptLoad(ctx context.Context, script string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionLoad(ctx context.Context, code string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionLoadReplace(ctx context.Context, code string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionDelete(ctx context.Context, libName string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionFlush(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionKill(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionFlushAsync(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionList(ctx context.Context, q FunctionListQuery) *FunctionListCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionDump(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionRestore(ctx context.Context, libDump string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FunctionStats(ctx context.Context) *FunctionStatsCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FCall(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FCallRo(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FCallRO(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Publish(ctx context.Context, channel string, message interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SPublish(ctx context.Context, channel string, message interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

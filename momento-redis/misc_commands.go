package momento_redis

import (
	"context"
	"time"

	. "github.com/redis/go-redis/v9"
)

func (m *MomentoRedisClient) Command(ctx context.Context) *CommandsInfoCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) CommandList(ctx context.Context, filter *FilterBy) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) CommandGetKeys(ctx context.Context, commands ...interface{}) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) CommandGetKeysAndFlags(ctx context.Context, commands ...interface{}) *KeyFlagsCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientGetName(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Echo(ctx context.Context, message interface{}) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Ping(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Quit(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Del(ctx context.Context, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Unlink(ctx context.Context, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Dump(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Exists(ctx context.Context, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Keys(ctx context.Context, pattern string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Move(ctx context.Context, key string, db int) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ObjectRefCount(ctx context.Context, key string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ObjectEncoding(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ObjectIdleTime(ctx context.Context, key string) *DurationCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Persist(ctx context.Context, key string) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PTTL(ctx context.Context, key string) *DurationCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RandomKey(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Rename(ctx context.Context, key, newkey string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RenameNX(ctx context.Context, key, newkey string) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Sort(ctx context.Context, key string, sort *Sort) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SortRO(ctx context.Context, key string, sort *Sort) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SortStore(ctx context.Context, key, store string, sort *Sort) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SortInterfaces(ctx context.Context, key string, sort *Sort) *SliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Touch(ctx context.Context, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) TTL(ctx context.Context, key string) *DurationCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Type(ctx context.Context, key string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *ZSliceWithKeyCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PFCount(ctx context.Context, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BgRewriteAOF(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) BgSave(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientKill(ctx context.Context, ipPort string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientList(ctx context.Context) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientInfo(ctx context.Context) *ClientInfoCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientPause(ctx context.Context, dur time.Duration) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientUnpause(ctx context.Context) *BoolCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientID(ctx context.Context) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientUnblock(ctx context.Context, id int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ClientUnblockWithError(ctx context.Context, id int64) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ConfigGet(ctx context.Context, parameter string) *MapStringStringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ConfigResetStat(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ConfigSet(ctx context.Context, parameter, value string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ConfigRewrite(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) DBSize(ctx context.Context) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FlushAll(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FlushAllAsync(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FlushDB(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) FlushDBAsync(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Info(ctx context.Context, section ...string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) LastSave(ctx context.Context) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Save(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Shutdown(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ShutdownSave(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ShutdownNoSave(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SlaveOf(ctx context.Context, host, port string) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) SlowLogGet(ctx context.Context, num int64) *SlowLogCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) Time(ctx context.Context) *TimeCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) DebugObject(ctx context.Context, key string) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ReadOnly(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ReadWrite(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) MemoryUsage(ctx context.Context, key string, samples ...int) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PubSubNumSub(ctx context.Context, channels ...string) *MapStringIntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PubSubNumPat(ctx context.Context) *IntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PubSubShardChannels(ctx context.Context, pattern string) *StringSliceCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) PubSubShardNumSub(ctx context.Context, channels ...string) *MapStringIntCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ACLDryRun(ctx context.Context, username string, command ...interface{}) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ACLLog(ctx context.Context, count int64) *ACLLogCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ACLLogReset(ctx context.Context) *StatusCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

func (m *MomentoRedisClient) ModuleLoadex(ctx context.Context, conf *ModuleLoadexConfig) *StringCmd {

	panic(UnsupportedOperationError("This operation has not been implemented yet"))
}

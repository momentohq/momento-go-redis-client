<img src="https://docs.momentohq.com/img/logo.svg" alt="logo" width="400"/>

[![project status](https://momentohq.github.io/standards-and-practices/badges/project-status-incubating.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)
[![project stability](https://momentohq.github.io/standards-and-practices/badges/project-stability-experimental.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)


# Momento Go-Redis Compatibility Client

## What and Why?

This project provides a Momento-backed implementation of [go-redis](hhttps://github.com/redis/go-redis).
The goal is to provide a drop-in replacement for [go-redis](hhttps://github.com/redis/go-redis) so that you can
use the same code with either a Redis server or with the Momento Cache service!

You can use Momento as your cache engine for any Go project that support a redis-backed cache.

## Usage

To switch your existing `go-redis` application to use Momento, you only need to change the code where you construct your client object:

<table>
<tr>
 <td width="50%">With go-redis client</td>
 <td width="50%">With Momento's go-redis compatibility client</td>
</tr>
<tr>
 <td width="50%" valign="top">

```go
package redis
import (
	"github.com/redis/go-redis/v9"
)
// Replace these values with your Redis server's details
REDIS_HOST := "my.redis-server.com"
REDIS_PORT := 6379
// Create a Redis client
redisClient := redis.NewClient(&redis.Options{Addr: REDIS_HOST + ":" + REDIS_PORT,})
```
</td>
<td width="50%">

```go
package redis

import (
	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/momento-go-redis-client/momento-redis"
)

credential, _ := auth.NewEnvMomentoTokenProvider("MOMENTO_AUTH_TOKEN")
cacheClient, _ := momento.NewCacheClient(config.LaptopLatest(), credential, 60*time.Second)
// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
cacheClient.CreateCache(sContext.Ctx, mClient, cacheName)
redisClient, _ := momento_redis.NewMomentoRedisClient(mClient, cacheName)
```
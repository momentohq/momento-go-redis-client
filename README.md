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
	"context"
	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/momento-go-redis-client/momento-redis"
)

credential, _ := auth.NewEnvMomentoTokenProvider("MOMENTO_AUTH_TOKEN")
cacheClient, _ := momento.NewCacheClient(config.LaptopLatest(), credential, 60*time.Second)
// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
cacheClient.CreateCache(context.Background(), &momento.CreateCacheRequest {CacheName : "default_cache"})
redisClient, _ := momento_redis.NewMomentoRedisClient(cacheClient, "default_cache")
```

</td>
</tr>
</table>

**NOTE**: The Momento `momento-redis` implementation currently supports simple key/value pairs (`GET`, `SET`, `SETNX`, `DEL`, `EXPIRE`, `TTL`). 
We will continue to add support for additional Redis APIs in the future; for more information see the [Current Redis API Support](#current-redis-api-support) section later in this doc.

## Current Redis API Support

This library supports the most popular Redis APIs, but does not yet support all Redis APIs. We currently support the most
common APIs related to string values (GET, SET, etc.). We will be adding support for additional
APIs in the future. If there is a particular API that you need support for, please drop by our [Discord](https://discord.com/invite/3HkAKjUZGq)
or e-mail us at [support@momentohq.com](mailto:support@momentohq.com) and let us know!

In the meantime, if you call a method from the `momento-redis` API that we do not yet support, you will get a panic for 
`UnsupportedOperationError`; letting you know that the method is not implemented yet.
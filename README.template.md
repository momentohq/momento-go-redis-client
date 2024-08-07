
{{ ossHeader }}

# Momento Go-Redis Compatibility Client

## What and Why?

This project provides a Momento-backed implementation of [go-redis](hhttps://github.com/redis/go-redis).
The goal is to provide a drop-in replacement for [go-redis](hhttps://github.com/redis/go-redis) so that you can
use the same code with either a Redis server or with the Momento Cache service!

You can use Momento as your cache engine for any Go project that support a redis-backed cache.

## Usage

To switch your existing `go-redis` application to use Momento, you only need to change the code where you construct your client object:

### With go-redis client

```go
package redis

import (
  "github.com/redis/go-redis/v9"
)

func initRedisClient() redis.Cmdable {
  // Replace these values with your Redis server's details
  REDIS_HOST := "my.redis-server.com"
  REDIS_PORT := "6379"
  // Create a Redis client
  redisClient := redis.NewClient(&redis.Options{Addr: REDIS_HOST + ":" + REDIS_PORT})
  return redisClient
}
```

### With Momento's go-redis compatibility client

This snippet assumes that a Momento Cache named "default_cache" already exists.
You can create a cache using the `CreateCache` method or using the [Momento Console](https://console.gomomento.com/caches).

```go
package redis

import (
  "context"
  "github.com/momentohq/client-sdk-go/auth"
  "github.com/momentohq/client-sdk-go/config"
  "github.com/momentohq/client-sdk-go/momento"
  momentoredis "github.com/momentohq/momento-go-redis-client/momento-redis"
  "github.com/redis/go-redis/v9"
  "time"
)

func initRedisClient() redis.Cmdable {
  credential, eErr := auth.NewEnvMomentoTokenProvider("MOMENTO_API_KEY")
  if eErr != nil {
    panic("Failed to initialize credentials through API key " + eErr.Error())
  }
  cacheClient, cErr := momento.NewCacheClient(config.LaptopLatest(), credential, 60*time.Second)
  if cErr != nil {
    panic("Failed to initialize Momento cache client " + cErr.Error())
  }
  redisClient := momentoredis.NewMomentoRedisClient(cacheClient, "default_cache")
  return redisClient
}
```

**NOTE**: The Momento `momento-redis` implementation currently supports a subset of Redis commands.
We will continue to add support for additional Redis APIs in the future;
for more information see the [Current Redis API Support](#current-redis-api-support) section later in this doc.

## Installation

```bash
go get github.com/momentohq/momento-go-redis-client
```

## Examples

### Prerequisites

To run these examples, you will need a Momento API key. You can generate one using the [Momento Console](https://console.gomomento.com)

The examples will utilize your API key via the environment variable `MOMENTO_API_KEY` you set.

### Basic Example

In the [`examples/basic`](./examples/basic) directory, you will find a simple CLI app that does some basic sets and gets
on string values. You can also run the tests against your Redis server by providing ```-useRedis``` flag along
with ```-host x.x.x.x``` and ```-port xxxx```.

You can run the example via `go run`.

Here's an example run against Momento Cache:

```bash
cd examples/basic
export MOMENTO_API_KEY=<your momento API key goes here>
go run main.go -cacheName cache
```

And the output should look something like this:

```bash
INFO (CacheClient): Creating cache with name: cache
INFO (CacheClient): Cache 'cache' created successfully

Using Momento as a backend for go-redis with cache name "cache"

-------------------------------------------------
-----------------------SET-----------------------
Successfully set key "Momento" with response "OK"
-----------------------SET-----------------------

-----------------------GET-----------------------
Got response value as "cache" for key "Momento"
-----------------------GET-----------------------

-----------------------DEL-----------------------
Delete successful for key "Momento" with response "1"
-----------------------DEL-----------------------

-----------------------SETNX-----------------------
Successfully set key "Momento" with response "true"
-----------------------SETNX-----------------------

-----------------------EXPIRE-----------------------
Successfully set expiration for key "Momento" with response "true"
-----------------------EXPIRE-----------------------

-----------------------TTL-----------------------
Successfully received remaining ttl 1m39.969s for key "Momento"
-----------------------TTL-----------------------

INFO (CacheClient): Deleting cache with name: cache
INFO (CacheClient): Cache 'cache' deleted successfully
```

To run against Redis, the command will look like:

```bash
 go run main.go -useRedis -host 127.0.0.1 -port 6379
```

## Current Redis API Support

This library supports the most popular Redis APIs, but does not yet support all Redis APIs.

<table>
  <tr>
    <td>
      Scalar Operations
    </td>
    <td>
      SET, GET, DEL, SETNX, EXPIRE, TTL
    </td>
  </tr>
  <tr>
    <td>
      Sorted Set Operations
    </td>
    <td>
      ZADD, ZRANGEBYSCORE, ZRANGEBYSCOREWITHSCORES, ZREVRANGEBYSCORE, ZREVRANGEBYSCOREWITHSCORES
    </td>
  </tr>
  <tr>
    <td>
      Dictionary Operations
    </td>
    <td>
      HSET
    </td>
  </tr>
</table>

We will be adding support for additional APIs in the future. 
If there is a particular API that you need support for, please drop by our [Discord](https://discord.com/invite/3HkAKjUZGq) or e-mail us at [support@momentohq.com](mailto:support@momentohq.com) and let us know!

In the meantime, if you call a method from the `momento-redis` API that we do not yet support, you will get a panic for `UnsupportedOperationError` letting you know that the method is not yet implemented.

### Go-Lang Compile-Time API Checking

If you'd like compile-time checking to tell you if you are using any APIs that we don't yet
support, we provide our own `MomentoRedisCmdable` interface, which is a fully compatible subset of the official `go-redis`
interface `Cmdable`, but explicitly lists out the APIs that we currently support.

With a one-line change to your initialization call, you get back an instance of this interface instead of the
default `redis.Cmdable` interface. Then the go-lang compiler will catch any calls your code is making to Redis
API methods that we don't yet support, so you'll know before you even try to run the code.

All you need to do is type the `MomentoRedisClient` object we instantiated above as
`MomentoRedisCmdable`. Here's what it looks like:

```go
package redis

import (
  "context"
  "github.com/momentohq/client-sdk-go/auth"
  "github.com/momentohq/client-sdk-go/config"
  "github.com/momentohq/client-sdk-go/momento"
  momentoredis "github.com/momentohq/momento-go-redis-client/momento-redis"
  "time"
)

// only change in the function definition from before and the body remains the same
func initRedisClient() momentoredis.MomentoRedisCmdable {
  credential, eErr := auth.NewEnvMomentoTokenProvider("MOMENTO_API_KEY")
  if eErr != nil {
    panic("Failed to initialize credentials through API key " + eErr.Error())
  }
  cacheClient, cErr := momento.NewCacheClient(config.LaptopLatest(), credential, 60*time.Second)
  if cErr != nil {
    panic("Failed to initialize Momento cache client " + cErr.Error())
  }
  // create cache; it resumes execution normally incase the cache already exists
  _, createErr := cacheClient.CreateCache(context.Background(), 
                    &momento.CreateCacheRequest{CacheName: "default_cache"})
  if createErr != nil {
    panic("Failed to create cache with cache name default cache \n" + createErr.Error())
  }
  redisClient := momentoredis.NewMomentoRedisClient(cacheClient, "default_cache")
  return redisClient
}
```

Exactly the same initialization call as before other than the `momento_redis.MomentoRedisCmdable` type, and now you get compile-time compatibility checking!\*

If you try this, and your code doesn't compile because we are missing APIs that you need, please do reach out to us!

\* Note that some flags are not supported. You may get a runtime error (`UnsupportedOperationError`) for those.

{{ ossFooter }}

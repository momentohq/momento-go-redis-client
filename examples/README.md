## Basic Example

### Prerequisites

To run these examples, you will need a Momento API key. You can generate one using the [Momento Console](https://console.gomomento.com)

The examples will utilize your API key via the environment variable `MOMENTO_API_KEY` you set.

## Installation

```bash
go get github.com/momentohq/momento-go-redis-client
```

### Basic Example

In the [`basic`](./basic) directory, you will find a simple CLI app that does some basic sets and gets
on string values. You can also run the tests against your Redis server by providing ```-useRedis=true``` flag

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

#### Running examples with compile-time checking and Momento flavor interface

If you'd like compile-time checking to tell you if you are using any APIs that we don't yet
support, we provide our own `MomentoRedisCmdable` interface, which is a fully compatible subset of the official `go-redis`
interface `Cmdable`, but explicitly lists out the APIs that we currently support.

If you want the examples to be run using the Momento flavor interface,
change the type from ```redis.Cmdable``` to ```momentoredis.MomentoRedisCmdable``` where the client variable is declared near the 
top of main(). With this one line of code change, you get compile time checking and it also runs against both 
Momento and Redis!

###### From:

```go
    // change this to the type momentoredis.MomentoRedisCmdable for compile-time checking. This interface only
    // has the Redis Commands that this compatibility client supports.
    // var client momentoredis.MomentoRedisCmdable
    var client redis.Cmdable
```

###### To:

```go
    // change this to the type momentoredis.MomentoRedisCmdable for compile-time checking. This interface only
    // has the Redis Commands that this compatibility client supports.
    var client momentoredis.MomentoRedisCmdable
    // var client redis.Cmdable
```
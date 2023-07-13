### Basic Example

In the [`basic`](./basic) directory, you will find a simple CLI app that does some basic sets and gets
on string values. You can also run the tests against your Redis server by providing ```-useRedis=true``` flag

You can run the example via `go run`.

Here's an example run against Momento Cache:

```bash
cd examples/basic
export MOMENTO_AUTH_TOKEN=<your momento auth token goes here>
go run main.go -cacheName cache -authToken $MOMENTO_AUTH_TOKEN
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
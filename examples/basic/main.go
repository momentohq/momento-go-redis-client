package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	momento_redis "github.com/momentohq/momento-go-redis-client/momento-redis"
	"github.com/redis/go-redis/v9"
)

var useRedis bool

type RedisOptions struct {
	host string
	port string
}

type MomentoOptions struct {
	cacheName         string
	authToken         string
	defaultTTlSeconds int
}

var options interface{}
var momentoCacheClient momento.CacheClient
var momentoCacheName string

func main() {
	parseFlags()

	// Based on the -useRedis command line flag, we initialize either the
	// Redis client or the Momento backed Redis client
	var client redis.Cmdable

	switch o := options.(type) {
	case *RedisOptions:
		client = initGoRedisClient(*o)
		fmt.Println("\nUsing Redis as a backend for go-redis with host " + o.host + " and port " + o.port)

	case *MomentoOptions:
		client = initMomentoRedisClient(*o)
		fmt.Println("\nUsing Momento as a backend for go-redis with cache name \"" + o.cacheName + "\"")

	default:
		panic("nothing found")
	}
	fmt.Println("\n-------------------------------------------------")

	key := "Momento"
	sResp := client.Set(context.Background(), key, "cache", 60*time.Second)
	if sResp.Val() != "OK" {
		panic("Set response should be OK")
	} else {
		fmt.Println("-----------------------SET-----------------------")
		fmt.Println("Successfully set key \"" + key + "\" with response \"" + sResp.Val() + "\"")
		fmt.Println("-----------------------SET-----------------------")
	}
	fmt.Println()
	gResp := client.Get(context.Background(), key)
	if gResp.Val() == "" {
		panic("Get response should have returned value")
	} else {
		fmt.Println("-----------------------GET-----------------------")
		fmt.Println("Got response value as \"" + gResp.Val() + "\" for key \"" + key + "\"")
		fmt.Println("-----------------------GET-----------------------")
	}
	fmt.Println()
	dResp := client.Del(context.Background(), key)
	if dResp.Val() == 0 {
		panic("Should have successfully deleted key")
	} else {
		fmt.Println("-----------------------DEL-----------------------")
		fmt.Println("Delete successful for key \"" + key + "\" with response \"" + strconv.FormatInt(dResp.Val(), 10) + "\"")
		fmt.Println("-----------------------DEL-----------------------")

	}
	fmt.Println()
	sNXResp := client.SetNX(context.Background(), key, "cache", 60*time.Second)
	if sNXResp.Val() != true {
		panic("Set Not exists response should be true/successful")
	} else {
		fmt.Println("-----------------------SETNX-----------------------")
		fmt.Println("Successfully set key \"" + key + "\" with response \"" + strconv.FormatBool(sNXResp.Val()) + "\"")
		fmt.Println("-----------------------SETNX-----------------------")

	}
	fmt.Println()
	expResp := client.Expire(context.Background(), key, 100*time.Second)
	if expResp.Val() != true {
		panic("Expire response should be true/successful")
	} else {
		fmt.Println("-----------------------EXPIRE-----------------------")
		fmt.Println("Successfully set expiration for key \"" + key + "\" with response \"" + strconv.FormatBool(expResp.Val()) + "\"")
		fmt.Println("-----------------------EXPIRE-----------------------")
	}
	fmt.Println()
	ttlResp := client.TTL(context.Background(), key)
	if ttlResp.Val() == -2 {
		panic("Received no key exists response for existing key \"" + key + " while fetching TTL")
	} else {
		fmt.Println("-----------------------TTL-----------------------")
		fmt.Println("Successfully received remaining ttl " + ttlResp.Val().String() + " for key \"" + key + "\"")
		fmt.Println("-----------------------TTL-----------------------\n")
	}

	if !useRedis {
		momentoCacheClient.DeleteCache(context.Background(), &momento.DeleteCacheRequest{
			CacheName: momentoCacheName,
		})
	}

}

func initMomentoRedisClient(options MomentoOptions) *momento_redis.MomentoRedisClient {
	momentoCacheName = options.cacheName
	credential, _ := auth.NewStringMomentoTokenProvider(options.authToken)
	cacheClient, _ := momento.NewCacheClient(config.LaptopLatest(), credential, time.Duration(options.defaultTTlSeconds)*time.Second)
	// create cache; it resumes execution normally incase the cache already exists and isn't exceptional
	cacheClient.CreateCache(context.Background(), &momento.CreateCacheRequest{
		CacheName: options.cacheName,
	})

	redisClient := momento_redis.NewMomentoRedisClient(cacheClient, options.cacheName)
	momentoCacheClient = cacheClient
	return redisClient
}

func initGoRedisClient(options RedisOptions) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: options.host + ":" + options.port})
}

func parseFlags() {
	flag.BoolVar(&useRedis, "useRedis", false, "Specifies whether to use "+
		"the Redis or Momento cache engine to run operations")
	host := flag.String("host", "", "Hostname for the Redis server")
	port := flag.String("port", "", "Hostname for the Redis server")
	cacheName := flag.String("cacheName", "", "Cache name for the Momento service")
	authToken := flag.String("authToken", "", "Auth token for the Momento service")
	defaultTTL := flag.Int("defaultTTLSeconds", 60, "The default TTL for your Momento cache")
	flag.Parse()

	if useRedis {

		if *host == "" || *port == "" {
			panic("Running in Redis mode: Redis host (-host) and port (-port) should be provided through command line arguments")
		}
		options = &RedisOptions{
			host: *host,
			port: *port,
		}
	} else {

		if *cacheName == "" || *authToken == "" {
			panic("Running in Momento mode: Momento cacheName (-cacheName) and authToken (-authToken) should be provided through command line arguments." +
				"For Redis more, use flag -useRedis along with -host and -port")
		}
		options = &MomentoOptions{
			cacheName:         *cacheName,
			authToken:         *authToken,
			defaultTTlSeconds: *defaultTTL,
		}
	}

}

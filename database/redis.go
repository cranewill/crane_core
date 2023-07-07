package database

import (
	"github.com/go-redis/redis"
	"runtime"
)

var clusterClient redis.UniversalClient

func StartRedis() {
	poolSize := runtime.NumCPU()
	clusterClient = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       1,
		PoolSize: poolSize,
	})
	//if gosconf.Ins.ServerName == "beta" {
	//	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
	//		Addrs:    []string{gosconf.Ins.Redis.Addr},
	//		Password: gosconf.Ins.Redis.Password,
	//		PoolSize: poolSize,
	//	})
	//} else {
	//	clusterClient = redis.NewClient(&redis.Options{
	//		Addr:     gosconf.Ins.Redis.Addr,
	//		Password: gosconf.Ins.Redis.Password,
	//		DB:       gosconf.Ins.Redis.DB,
	//		PoolSize: poolSize,
	//	})
	//}
}

func RedisIns() redis.UniversalClient {
	return clusterClient
}

func SetRedisIns(client redis.UniversalClient) {
	clusterClient = client
}

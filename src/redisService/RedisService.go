/**
 * @Author: David Ma
 * @Date: 2021/2/5
 */
package redisService

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

var clusterClient *redis.ClusterClient

var DEFAULTREDISADDRS = []string{"192.168.211.128:7000", "192.168.211.128:7001", "192.168.211.128:7002"}

func init() {
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:          DEFAULTREDISADDRS,
		ReadOnly:       true,         //允许slave节点执行读命令，也就是可以实现读写分离
		RouteByLatency: true,         //从负责管理该slot的master节点和slave节点中选择一个来执行读命令，选取规则为：优先选取ping延迟小的那个节点
		RouteRandomly:  false,        //随机从负责管理该slot的master节点和slave节点中选择一个来执行读命令，选取规则为：随机选取
		Password:       "davidmaqq0", //相当于-a选项，对应server节点requirePass字段的值
	})
}

func changeDefaultRedisCluster(addrs []string) {

}

func IsRedisDefaultClusterClientPoolAlive(ctx context.Context) bool {
	if _, err := clusterClient.Ping(ctx).Result(); err != nil {
		zap.L().Error("redisClusterConnectionError:" + err.Error())
		poolStats := clusterClient.PoolStats()
		zap.L().Info("redisClusterConnectionPoolStates:" + fmt.Sprintf("总连接数=%d,空闲连接数=%d,已经移除的连接数=%d", poolStats.TotalConns, poolStats.IdleConns, poolStats.StaleConns))
		return false
	}
	return true
}

func Set(ctx context.Context, key, value string, expiration time.Duration) bool {
	if _, err := clusterClient.Set(ctx, key, value, expiration).Result(); err != nil {
		zap.L().Error(err.Error())
		return false
	}
	return true
}

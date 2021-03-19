/**
 * @Author: David Ma
 * @Date: 2021/2/5
 */
package redisService

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedisClusterOption(Addrs []string, ReadOnly, RouteByLatency, RouteRandomly bool,
	AuthPWD string, PoolSize, MinIdleConns int) *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:          Addrs,
		ReadOnly:       ReadOnly,       //允许slave节点执行读命令，也就是可以实现读写分离
		RouteByLatency: RouteByLatency, //从负责管理该slot的master节点和slave节点中选择一个来执行读命令，选取规则为：优先选取ping延迟小的那个节点
		RouteRandomly:  RouteRandomly,  //随机从负责管理该slot的master节点和slave节点中选择一个来执行读命令，选取规则为：随机选取
		Password:       AuthPWD,        //相当于-a选项，对应server节点requirePass字段的值
		PoolSize:       PoolSize,
		MinIdleConns:   MinIdleConns,
	}
}
func GetRedisClusterConnectionPoll(ctx context.Context, options *redis.ClusterOptions) *redis.ClusterClient {
	redisConnectionPoll := redis.NewClusterClient(options)
	if _, err := redisConnectionPoll.Ping(ctx).Result(); err != nil {
		zap.L().Error("redisClusterConnectionError:" + err.Error())
		poolStats := redisConnectionPoll.PoolStats()
		zap.L().Info("redisClusterConnectionPoolStates:" + fmt.Sprintf("总连接数=%d,空闲连接数=%d,已经移除的连接数=%d", poolStats.TotalConns, poolStats.IdleConns, poolStats.StaleConns))
		return nil
	}
	return redisConnectionPoll
}

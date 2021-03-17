/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package util

import (
	"reflect"
	"sync"
	"time"
)

/*
twitter 雪花算法 非线程安全版本
把时间戳,工作机器ID, 序列号组合成一个 64位 int
第一位置零(作为符号位，ID必然是一个非负值，所以第一位恒为0), [2,42]这41位存放时间戳,[43,52]这10位存放机器id,[53,64]最后12位存放序列号

【问题】：仅用DCL或者CAS来保证对sn的操作是线程安全的但也不能保证整个接口是线程安全性，因为对变量lastTimeStamp、curTimeStamp、rightBinValue、id、machineID的操作并不是线程安全的，
【解决思路】：1、这里考虑把变量：curTimeStamp、rightBinValue、id、machineID当做每个goroutine私有的来处理。可考虑借鉴JAVA的ThreadLoacl方式来实现，或者利用Context来实现？
2、可以采用DCL方式或者CAS操作来实现对变量lastTimeStamp的写操作是线程安全的
*/
var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]，指的不一定是mac地址
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ],同一毫秒最多产生4095个序列号
	lastTimeStamp int64 // 记录该接口上次被调用时的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒(nanoseconds)
	lock          sync.RWMutex
)

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1000000 //这里要除以10^6是因为要保证时间戳最多占41位(2^64 / 2^20)
	SetMachineId()
}

func SetMachineId() { //设置一个机器标识，如IP编码,防止分布式机器生成重复码,也可以是IP地址加上一个随机值
	mid := IPStr2Int(IPV4())
	machineID = mid << 12 //// 把机器 id 左移 12 位,让出 12 位空间给序列号使用
}

func GetSnowflakeId() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000000 //记录该接口本次被调用时的时间戳(毫秒级)

	// 如果两次调用在同一毫秒
	if curTimeStamp == lastTimeStamp {
		sn++
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sn > 4095 { //这里采用DCL保证对sn的操作是线程安全的，否则有可能当每毫秒的并发数超过4095的时候导致出现重复的ID(有可能出现这么一种情况：有两个在同一毫秒发起的请求且此时都判定sn>4095,如果没有施加DCL则有可能会导致将sn重复归0，从而产生重复ID)
			lock.Lock()
			if sn > 4095 {
				time.Sleep(time.Millisecond)
				curTimeStamp = time.Now().UnixNano() / 1000000 //这些都是临界资源
				lastTimeStamp = curTimeStamp
				sn = 0
			}
			lock.Unlock()
		}

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}

	if curTimeStamp > lastTimeStamp {
		if sn != 0 { //这里也是采用DCL保证对sn的操作是线程安全的
			lock.Lock()
			if sn != 0 {
				sn = 0
			}
			lock.Unlock()
		}

		lastTimeStamp = curTimeStamp
		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行与操作
		// 结果的( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}

	return 0
}

//去重
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

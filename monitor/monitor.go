/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: watch.go
 * @time: 2019/8/22 8:52
 */
package monitor

import (
	`sync/atomic`
	`time`

	`github.com/astaxie/beego/logs`

	`github.com/generalzgd/grpc-svr-frame/prewarn`
)

const (
	Stat_Tps = 0 + iota
	Stat_Qps
	Stat_Mem
	Stat_Goroutine
	Stat_Analyse
)

var (
	tickerCount int32

	list []IMonitor

	warnHandler func(string)

	// intervalCfg = map[int]int{
	// 	Stat_Tps: 1,   // 1 秒
	// 	Stat_Qps: 300, // 5分钟
	// }
	// thresholdCfg = map[int]int{
	// 	Stat_Tps: 100000, // 每秒 1000,000次
	// 	Stat_Qps: 1000000000,
	// }
)

func Register(tar IMonitor) {
	if tar == nil {
		return
	}
	list = append(list, tar)
}

func SetWarnHandler(f func(string)) {
	warnHandler = f
}

// 触发事件的方法，外部使用要埋点的地方调用该方法
func NewRecord(typ int, args ...interface{}) {
	for _, item := range list {
		if item.GetType() == typ {
			item.NewRecord(typ, args...)
		}
	}
}

func init() {
	// 超过阈值需要发起预警，设置预警接收方法
	SetWarnHandler(prewarn.NewWarn)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer func() {
			ticker.Stop()
			if r := recover(); r != nil {
			}
		}()

		for {
			select {
			case <-ticker.C:
				onTicker()
			}
		}
	}()
}

func onTicker() {
	cnt := int(atomic.AddInt32(&tickerCount, 1))
	for _, item := range list {
		if cnt%item.GetInterval() != 0 {
			continue
		}
		// 预警
		value := item.GetCount()
		thres := item.GetThreshold()
		if warnHandler != nil {
			if str, ok := item.MakeWarnStr(value, thres); ok {
				warnHandler(str)
			}
		}
		if str, ok := item.Debug(value, thres); ok {
			logs.Debug(str)
		}
	}
}

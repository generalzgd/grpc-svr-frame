/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: calculate.go
 * @time: 2019/8/21 19:49
 */
package tps

import (
	`fmt`
	`sync/atomic`

	`github.com/generalzgd/grpc-svr-frame/statistic`
)

func init() {
	monitor.Register(&StatTps{
		threshold: 10000,
	})
}

type StatTps struct {
	procNum   int32
	threshold int // 阈值
}

func (p *StatTps) NewRecord(typ int, args ...interface{}) {
	if typ != monitor.Stat_Tps {
		return
	}
	atomic.AddInt32(&p.procNum, 1)
}

func (p *StatTps) GetType() int {
	return monitor.Stat_Tps
}

func (p *StatTps) GetCount() int {
	num := atomic.SwapInt32(&p.procNum, 0)
	return int(num)
}

func (p *StatTps) GetInterval() int {
	return 1 // 1秒间隔
}

func (p *StatTps) GetThreshold() int {
	return p.threshold
}

func (p *StatTps) MakeWarnStr(tps int, threshold int) (string, bool) {
	if tps > 0 && threshold > 0 && tps > threshold {
		return fmt.Sprintf("告警 Tps:%d, threshold:%d", tps, threshold), true
	}
	return "", false
}

func (p *StatTps) Debug(tps int, threshold int) (string, bool) {
	if tps > 0 {
		return fmt.Sprintf("Tps:%d, threshold:%d", tps, threshold), true
	}
	return "", false
}

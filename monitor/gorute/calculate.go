/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: calculate.go
 * @time: 2019/8/22 11:26
 */
package gorute

import (
	`fmt`
	`runtime`

	`github.com/generalzgd/grpc-svr-frame/monitor`
)

func init() {
	monitor.Register(&StatGorutine{
		threshold: 50000,
	})
}

type StatGorutine struct {
	threshold int
}

func (p *StatGorutine) NewRecord(typ int, args ...interface{}) {
	if typ != monitor.Stat_Goroutine {
		return
	}
}

func (p *StatGorutine) GetType() int {
	return monitor.Stat_Goroutine
}

func (p *StatGorutine) GetCount() int {
	return runtime.NumGoroutine()
}

func (p *StatGorutine) GetInterval() int {
	return 60
}

func (p *StatGorutine) GetThreshold() int {
	return p.threshold
}

func (p *StatGorutine) MakeWarnStr(num int, threshold int) (string, bool) {
	if num > 0 && threshold > 0 && num > threshold {
		return fmt.Sprintf("告警 goroutine:%d threshold:%d", num, threshold), true
	}
	return "", false
}

func (p *StatGorutine) Debug(num int, threshold int) (string, bool) {
	return "", false
}

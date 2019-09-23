/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: calculate.go
 * @time: 2019/8/21 20:40
 */
package qps

import (
	`fmt`
	`time`

	`github.com/generalzgd/grpc-svr-frame/statistic`
)

func init() {
	monitor.Register(&StatQps{
		threshold:100000,
	})
}

type StatQps struct {
	procNum int32
	procDur time.Duration
	threshold int
}

func (p *StatQps) NewRecord(typ int, args ...interface{}) {
	if typ != monitor.Stat_Qps {
		return
	}
	p.procNum++
	for _, arg := range args {
		if d, ok := arg.(time.Duration); ok {
			p.procDur += d
		}
		break
	}
}

func (p *StatQps) GetType() int {
	return monitor.Stat_Qps
}

func (p *StatQps) GetCount() int {
	qpsamt := float64(p.procDur) / 1000000000
	if qpsamt > 0 {
		return int(float64(p.procNum) / qpsamt)
	}
	return 0
}

func (p *StatQps) GetInterval() int {
	return 30 // 30秒
}

func (p *StatQps) GetThreshold() int {
	return p.threshold
}

func (p *StatQps) MakeWarnStr(num int, threshold int) (string, bool) {
	if num >0 && threshold>0 && num > threshold {
		return fmt.Sprintf("告警 Qps:%d, threshold:%d", num, threshold), true
	}
	return "", false
}

func (p *StatQps) Debug(qps int, threshold int) (string, bool) {
	if qps > 0 {
		return fmt.Sprintf("Qps:%d, threshold:%d", qps, threshold), true
	}
	return "", false
}

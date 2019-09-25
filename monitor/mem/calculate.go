/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: calculate.go
 * @time: 2019/8/22 11:04
 */
package mem

import (
	`fmt`
	`runtime`
	`strings`

	`github.com/generalzgd/grpc-svr-frame/monitor`
)

func init() {
	monitor.Register(&StatMem{
		threshold: 500 * 1024 * 1024,
	})
}

type StatMem struct {
	threshold int // 内存上线
}

func (p *StatMem) NewRecord(typ int, args ...interface{}) {
	if typ != monitor.Stat_Mem {
		return
	}
}

func (p *StatMem) GetType() int {
	return monitor.Stat_Mem
}

func (p *StatMem) GetCount() int {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)

	return int(m.Mallocs)
}

func (p *StatMem) GetInterval() int {
	return 60
}

func (p *StatMem) GetThreshold() int {
	return p.threshold
}

func (p *StatMem) MakeWarnStr(num int, threshold int) (string, bool) {
	if num > 0 && threshold > 0 && num > threshold {
		return fmt.Sprintf("告警 memory: %s threshold: %s", formateMem(num), formateMem(threshold)), true
	}
	return "", false
}

func (p *StatMem) Debug(int, int) (string, bool) {
	return "", false
}

func formateMem(num int) string {
	var parts []int
	var i int
	for num > 0 && i <= 4 {
		i++
		yu := num % 1024
		num = num / 1024
		parts = append(parts, yu)
	}
	if num > 0 {
		parts = append(parts, num)
	}

	units := []string{"B", "K", "M", "G", "T"}
	list := make([]string, 0, len(parts))
	for i, val := range parts {
		if val > 0 {
			list = append(list, fmt.Sprintf("%d%s", val, units[i]))
		}
	}
	ll := len(list)
	for i := 0; i < ll/2; i++ {
		list[i], list[ll-1-i] = list[ll-1-i], list[i]
	}
	if ll > 2 {
		list = list[:2]
	}
	return strings.Join(list, "")
}

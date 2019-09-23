/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: istatistic.go
 * @time: 2019/8/22 8:53
 */
package monitor

type IMonitor interface {
	NewRecord(int, ...interface{})
	GetType() int
	GetCount() int
	GetInterval() int // 秒
	GetThreshold() int
	MakeWarnStr(int, int) (string, bool)
	Debug(int, int)(string,bool)
}

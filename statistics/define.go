/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: define.go
 * @time: 2019-07-30 15:42
 */

package statistics

import "github.com/toolkits/slice"

const (
	STS_SUM = "sum"
	STS_AVG = "avg"
	STS_TPS = "tps"
	STS_CNT = "count"
)

var (
	StsList = []string{STS_SUM, STS_AVG, STS_TPS, STS_CNT}
)

func ValidateStatisticType(in string) bool {
	return slice.ContainsString(StsList, in)
}
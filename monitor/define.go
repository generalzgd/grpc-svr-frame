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

package monitor

import "github.com/toolkits/slice"

const (
	// 分析类型
	ANALYSE_SUM = "sum"
	ANALYSE_AVG = "avg"
	ANALYSE_CNT = "count"
)

var (
	StsList = []string{ANALYSE_SUM, ANALYSE_AVG, ANALYSE_CNT}
)

func ValidateAnalyseType(in string) bool {
	return slice.ContainsString(StsList, in)
}
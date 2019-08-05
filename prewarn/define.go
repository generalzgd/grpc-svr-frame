/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: define.go
 * @time: 2019-07-30 15:48
 */

package prewarn

import "github.com/toolkits/slice"

const (
	WARN_SMS = "sms"
	WARN_EMAIL = "email"
)

var (
	WarnList = []string{WARN_SMS, WARN_EMAIL}
)

func ValidateWarnType(in string) bool {
	return slice.ContainsString(WarnList, in)
}
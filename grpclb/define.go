/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: policy.go
 * @time: 2019-07-18 14:01
 */

package grpclb

import "github.com/toolkits/slice"

const (
	POLICY_RANDOM    = "random"
	POLICY_HOLD      = "hold"
	POLICY_PERCENT   = "percent"
	POLICY_LOW_FIRST = "lowfirst"
)

var (
	PolicyList = []string{POLICY_RANDOM, POLICY_HOLD, POLICY_PERCENT, POLICY_LOW_FIRST}
)

func ValidatePolicyType(in string) bool {
	if slice.ContainsString(PolicyList, in) {
		return true
	}
	return false
}

type PolicyRandomConfig struct {
}

type PolicyHoldConfig struct {
}

type PolicyPercentConfig struct {
}

type PolicyLowFirstConfig struct {
}

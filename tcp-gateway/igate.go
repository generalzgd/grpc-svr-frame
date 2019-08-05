/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: igate.go
 * @time: 2019-07-17 22:18
 */

package tcp_gateway

import (
	"github.com/astaxie/beego"
	"github.com/funny/slab"
)

type IGate interface {
}


func makePool(t string, minChunk, maxChunk, factor, pageSize int) slab.Pool {
	switch t {
	case "sync":
		return slab.NewSyncPool(minChunk, maxChunk, factor)
	case "atom":
		return slab.NewAtomPool(minChunk, maxChunk, factor, pageSize)
	case "chan":
		return slab.NewChanPool(minChunk, maxChunk, factor, pageSize)
	default:
		beego.Error(`unsupported memory pool type, must be "sync", "atom" or "chan"`)
	}
	return nil
}
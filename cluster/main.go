/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: main.go
 * @time: 2019-05-20 19:44
 */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logger := logs.GetBeeLogger()
	logger.SetLevel(beego.LevelInformational)
	_ = logger.SetLogger(logs.AdapterConsole)
	_ = logger.SetLogger(logs.AdapterFile, `{"filename":"logs/file.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":7}`)
	logger.Async()
}

func main() {
	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}

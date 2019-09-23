/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: warn.go
 * @time: 2019/8/21 21:32
 */
package prewarn

import (
	`fmt`
	`strings`
	`sync`
	`time`
)

type Warn struct {
	lock        sync.Mutex
	contentList []string
}

func (p *Warn) Run() {
	ticker := time.NewTicker(time.Minute * 5)
	defer func() {
		ticker.Stop()
		if r := recover(); r != nil {
		}
	}()
	for {
		select {
		case <-ticker.C:
			p.onMinuteTick()
		}
	}
}

func (p *Warn) newWarn(str string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.contentList = append(p.contentList, str)
}

func (p *Warn) clean() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.contentList = make([]string, 0, 500)
}

func (p *Warn) onMinuteTick() {
	if len(p.contentList) == 0 || sendMailCallback == nil {
		return
	}
	tmp := make(map[string]int, len(p.contentList))
	for _, it := range p.contentList {
		if v, ok := tmp[it]; ok {
			tmp[it] = v + 1
		} else {
			tmp[it] = 1
		}
	}
	p.clean()
	li := make([]string, 0, len(tmp))
	for k, v := range tmp {
		li = append(li, fmt.Sprintf("%s (%d次)", k, v))
	}
	body := strings.Join(li, "\n")

	sendMailCallback(body)
}

var (
	inst *Warn

	sendMailCallback func(string)
)

func init() {
	inst = &Warn{
		contentList: make([]string, 0, 500),
	}

	go inst.Run()
}

func SetSendMailCallback(f func(string)) {
	sendMailCallback = f
}

func NewWarn(str string) {
	inst.newWarn(str)
}

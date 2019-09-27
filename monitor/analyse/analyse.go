/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: analyse.go
 * @time: 2019/9/25 14:19
 */
package analyse

import (
	`fmt`
	`math`
	`strconv`
	`sync`
	`sync/atomic`

	`github.com/generalzgd/comm-libs/pickjson`

	`github.com/generalzgd/grpc-svr-frame/monitor`
)

func NewAnalyse(threshold, analyseNum uint, analyseType, fieldName string) *Analyse {
	if threshold < 1 || analyseNum < 1 || !monitor.ValidateAnalyseType(analyseType) || len(fieldName) < 1 {
		return nil
	}
	p := &Analyse{
		threshold:   int(threshold),
		analyseType: analyseType,
		analyseNum:  int(analyseNum),
		fieldName:   fieldName,
		inFlow:      make(chan string, 1000),
	}
	go p.run()
	return p
}

// json分析统计，需要手动注册。如果有多个分析，请注册多次
type Analyse struct {
	threshold   int
	analyseNum  int    // 分析师数量
	analyseType string // 要分析的类型
	fieldName   string // 要分析的字段名
	inFlow      chan string
	// processFlow chan int // 统计数据流水线
	analyseStore int64 // 其他类型要累积数值
	analyseCnt   int64 // avg 类型需要记录
	lock         sync.Mutex
}

func (p *Analyse) run() {
	wg := sync.WaitGroup{}
	for i := 0; i < p.analyseNum; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer func() {
				wg.Done()
			}()
			for data := range p.inFlow {
				p.analyseData(data)
			}
		}(&wg)
	}
	wg.Wait()
}

// 解析出对应的字段值
func (p *Analyse) analyseData(data string) {

	res := pickjson.PickBytes([]byte(data), []byte(p.fieldName))
	if res == nil {
		return
	}
	num, err := strconv.Atoi(string(res))
	if err != nil {
		return
	}
	if p.analyseType == monitor.ANALYSE_AVG {
		p.lock.Lock()
		atomic.AddInt64(&p.analyseStore, int64(num))
		atomic.AddInt64(&p.analyseCnt, 1)
		p.lock.Unlock()
	} else {
		atomic.AddInt64(&p.analyseStore, int64(num))
		atomic.AddInt64(&p.analyseCnt, 1)
	}
}

// 定义参数
// typ int 监控类型
// data string json字符串
func (p *Analyse) NewRecord(typ int, args ...interface{}) {
	if typ != monitor.Stat_Analyse {
		return
	}

	if len(args) > 0 {
		if str, ok := args[0].(string); ok {
			p.inFlow <- str
		}
	}
}

func (p *Analyse) GetType() int {
	return monitor.Stat_Analyse
}

func (p *Analyse) GetCount() int {
	if p.analyseType == monitor.ANALYSE_AVG {
		p.lock.Lock()
		sum := atomic.SwapInt64(&p.analyseStore, 0)
		cnt := atomic.SwapInt64(&p.analyseCnt, 0)
		p.lock.Unlock()
		if cnt > 0 {
			return int(sum / cnt)
		} else if sum > 0 {
			return math.MaxInt64
		}
		return 0
	} else {
		sum := atomic.SwapInt64(&p.analyseStore, 0)
		cnt := atomic.SwapInt64(&p.analyseCnt, 0)
		if p.analyseType == monitor.ANALYSE_CNT {
			return int(cnt)
		} else if p.analyseType == monitor.ANALYSE_SUM {
			return int(sum)
		}
	}
	return 0
}

func (p *Analyse) GetInterval() int { // 60秒计算一次
	return 60
}

func (p *Analyse) GetThreshold() int {
	return p.threshold
}

func (p *Analyse) MakeWarnStr(num int, threshold int) (string, bool) {
	out := ""
	if num >= threshold {

		switch p.analyseType {
		case monitor.ANALYSE_SUM:
			out = fmt.Sprintf("Got sum:%d of %s field dur %d's", num, p.fieldName, p.GetInterval())
		case monitor.ANALYSE_CNT:
			out = fmt.Sprintf("Got count:%d of %s field dur %d's", num, p.fieldName, p.GetInterval())
		case monitor.ANALYSE_AVG:
			out = fmt.Sprintf("Got average:%d of %s field dur %d's", num, p.fieldName, p.GetInterval())
		}
	}
	return out, false
}

func (p *Analyse) Debug(num int, threshold int) (string, bool) {
	// if num > 0 {
	// 	return fmt.Sprintf("num:%d, threshold:%d", num, threshold), true
	// }
	return "", false
}

package tqueue

import "sync"

//环形队列
type CircleQueue struct {
	Total int
	lock  sync.RWMutex //读写锁
	valus []CircleQueueContent
	Maxl  int
}

type CircleQueueContent struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

//控制 Lines 长度 不超过 maxl 下追加新line
func (cir *CircleQueue) Append(line string) {
	cir.lock.Lock()
	if len(cir.valus) >= cir.Maxl {
		cir.valus = cir.valus[1:]
	}
	cir.Total++
	cir.valus = append(cir.valus, CircleQueueContent{
		Id:      cir.Total,
		Content: line,
	})
	cir.lock.Unlock()
}

//取内容 指定起始id
func (cir *CircleQueue) Getlines(strid int) []CircleQueueContent {
	var lines []CircleQueueContent
	cir.lock.Lock()
	if strid < cir.Total {
		lines = cir.valus
	} else {

		for _, v := range cir.valus {
			if v.Id > strid {
				lines = append(lines, v)
			}
		}

	}
	cir.lock.Unlock()
	return lines
}

//创建一个环形队列 指定 最大容量
func NewCircleQueue(Maxl int) *CircleQueue {

	var cir = &CircleQueue{
		Maxl: 100,
	}
	return cir
}

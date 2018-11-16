package workshop

import (
	"sync"
)

type jobs struct {
	m    *sync.Mutex
	list []job
}

type job interface {
	Do() error
	CallBack(func() error)
}

func (js *jobs) pop() job {
	js.m.Lock()
	defer js.m.Unlock()
	var job job
	if len(js.list) > 0 {
		job = js.list[0]
		js.list = js.list[1:]
	}
	return job
}

func (js *jobs) push(jobs ...job) {
	js.m.Lock()
	defer js.m.Unlock()
	if len(js.list) > 0 {
		js.list = js.list[:len(js.list)-1]
	}
	js.list = append(js.list, jobs...)
	js.list = append(js.list, _EOF)
}

func (js *jobs) len() int {
	js.m.Lock()
	defer js.m.Unlock()
	return len(js.list)

}

type e int

const (
	_EOF e = iota
)

func (e e) Do() error               { return nil }
func (e e) CallBack(f func() error) {}

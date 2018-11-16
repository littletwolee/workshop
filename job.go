package workshop

import (
	"sync"
)

type jobs struct {
	m    *sync.Mutex
	list []Job
}

type Job interface {
	Do(obj interface{}) error
	CallBack(interface{}, func(obj interface{}) error)
}

func (js *jobs) pop() Job {
	js.m.Lock()
	defer js.m.Unlock()
	var job Job
	if len(js.list) > 0 {
		job = js.list[0]
		js.list = js.list[1:]
	}
	return job
}

func (js *jobs) push(jobs ...Job) {
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

func (e e) Do(obj interface{}) error                                { return nil }
func (e e) CallBack(obj interface{}, f func(obj interface{}) error) {}

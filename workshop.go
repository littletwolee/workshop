package workshop

import (
	"sync"
)

type workShop struct {
	pip  *pip
	stop chan bool
	jobs *jobs
	wg   *sync.WaitGroup
}

func NewWorkShop(workerNum int) *workShop {
	return &workShop{
		pip:  newPip(workerNum),
		stop: make(chan bool),
		jobs: &jobs{
			m: new(sync.Mutex),
		},
		wg: new(sync.WaitGroup),
	}
}

func (w *workShop) AddJobs(jobs ...job) {
	w.jobs.push(jobs...)
	w.wg.Add(len(jobs))
}

func (w *workShop) Start() {
	go func(w *workShop) {
		for {
			select {
			case j := <-w.pip.in:
				go func(j job, w *workShop) {
					j.CallBack(j.Do)
					w.wg.Done()
					<-w.pip.out
				}(j, w)
			case <-w.stop:
				return
			default:
			}
		}
	}(w)
	for {
		j := w.jobs.pop()
		if j == nil {
			w.wg.Done()
			continue
		}
		if _, ok := j.(E); ok {
			break
		}
		w.pip.out <- true
		w.pip.in <- j
	}
}
func (w *workShop) Wait() {
	w.wg.Wait()
	w.stop <- true
}

// func (w *workShop) Stop() {
// 	w.inPuts.close <- true
// }

type pip struct {
	in  chan job
	out chan bool
}

func newPip(n int) *pip {
	return &pip{
		in:  make(chan job, n),
		out: make(chan bool, n),
	}
}

// func newInput(n int) *inPuts {
// 	return &inPuts{
// 		ch:  make(chan job, n),
// 		pip: newPip(),
// 	}
// }

// type inPuts struct {
// 	ch chan job
// 	pip
// }

// func newOutPut(n int) *outPuts {
// 	return &outPuts{
// 		Ch: make(chan error, n),
// 	}
// }

// type outPuts struct {
// 	Ch chan error
// }

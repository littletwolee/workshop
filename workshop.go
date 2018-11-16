package workshop

import (
	"sync"
)

type workShop struct {
	chNum int
	pip   *pip
	stop  chan bool
	jobs  *jobs
	wg    *sync.WaitGroup
}

func NewWorkShop(chNum int) *workShop {
	return &workShop{
		chNum: chNum,
		pip:   &pip{},
		stop:  make(chan bool),
		jobs: &jobs{
			m: new(sync.Mutex),
		},
		wg: new(sync.WaitGroup),
	}
}

func (w *workShop) AddJobs(jobs ...Job) {
	w.jobs.push(jobs...)
	w.wg.Add(len(jobs))
}

func (w *workShop) Start() {
	if w.pip == nil {
		return
	}
	w.pip.refresh(w.chNum)
	go func(w *workShop) {
		for {
			select {
			case j := <-w.pip.in:
				go func(j Job, w *workShop) {
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
		if _, ok := j.(e); ok {
			break
		}
		w.pip.out <- true
		w.pip.in <- j
	}
}
func (w *workShop) Wait() {
	w.wg.Wait()
	w.stop <- true
	w.pip.close()
}

type pip struct {
	in  chan Job
	out chan bool
}

// func newPip(chNum int) *pip {
// 	return &pip{
// 		in:  make(chan job, chNum),
// 		out: make(chan bool, chNum),
// 	}
// }

func (p *pip) refresh(chNum int) {
	p.in = make(chan Job, chNum)
	p.out = make(chan bool, chNum)
}

func (p *pip) close() {
	defer func() {
		if recover() != nil {

		}
	}()
	close(p.in)
	close(p.out)
}

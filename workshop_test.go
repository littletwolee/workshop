package workshop

import (
	"fmt"
	"testing"
	"time"
)

func Test_WorkShop(t *testing.T) {
	num := 100000
	chNum := 2
	ws := NewWorkShop(chNum)
	var list []Job
	for index := 0; index < num; index++ {
		list = append(list, newTestJob(index))
	}
	ws.AddJobs(list...)
	para := &para{sex: "male"}
	go ws.Start(para)
	ws.Wait()
	fmt.Println("completed")
}

type para struct {
	sex string
}

type testJob struct {
	id int
}

func newTestJob(id int) Job {
	return &testJob{
		id: id,
	}
}

func (t *testJob) Do(obj interface{}) error {
	time.Sleep(time.Second * 1)
	if t.id%2 == 0 {
		fmt.Printf("mod 2==0: %d, sex: %s\n", t.id, obj.(*para).sex)
		return nil
	}
	return fmt.Errorf("error %d", t.id)
}
func (t *testJob) CallBack(obj interface{}, f func(obj interface{}) error) {
	if err := f(obj); err != nil {
		fmt.Println(err)
	}
}

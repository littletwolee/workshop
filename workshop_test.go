package workshop

import (
	"fmt"
	"testing"
	"time"
)

func Test_WorkShop(t *testing.T) {
	num := 1000
	chNum := 10
	ws := NewWorkShop(chNum)
	var list []job
	for index := 0; index < num; index++ {
		list = append(list, newTestJob(index))
	}
	ws.AddJobs(list...)
	go ws.Start()
	ws.Wait()
	fmt.Println("completed")
}

type testJob struct {
	id int
}

func newTestJob(id int) job {
	return &testJob{
		id: id,
	}
}

func (t *testJob) Do() error {
	time.Sleep(1 * time.Second)
	if t.id%2 == 0 {
		fmt.Printf("mod 2==0: %d\n", t.id)
		return nil
	}

	return fmt.Errorf("error %d", t.id)
}
func (t *testJob) CallBack(f func() error) {
	if err := f(); err != nil {
		fmt.Println(err)
	}
}

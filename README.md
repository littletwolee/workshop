# workshop

## Summary

Golang pipline with channal & waitgroup

## Usage

```
// in ur shell
go get github.com/littletwolee/workshop
```

```
// import it
import (
	"github.com/littletwolee/workshop"
)
// get new workshop, chNum is a num that how much worker can working simultaneously
ws := NewWorkShop(chNum) 
// u need implementation job interface
// type job interface {
//      Do() error
//      CallBack(func() error)
// }
// like these:
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
// prepare your data
var list []job
for index := 0; index < num; index++ {
    list = append(list, newTestJob(index))
}
// add all jobs in workshop
ws.AddJobs(list...)
// create a goroutine to start it
go ws.Start()
// and wait it completed
ws.Wait()

....
ur next step code
....

```

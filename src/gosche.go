package main

import (
	"fmt"
	"sort"
)

type runner interface {
	Run(x interface {}) (ret int, err error)
}


type echoRunner struct {
}

func (er echoRunner) Run(x interface {}) (ret int, err error) {
	fmt.Println(x)
	ret, err = 0, nil
	return ret, err
}

///////////

type Task struct {
	hour, minute int
	task runner
}

type TaskSlice []Task

func (t TaskSlice) Len() int		{ return len(t) }
func (t TaskSlice) Less(i, j int) bool { return t[i].hour < t[j].hour || (t[i].hour == t[j].hour && t[i].minute < t[j].minute) }
func (t TaskSlice) Swap(i, j int)	{ t[i], t[j] = t[j], t[i] }

func (t TaskSlice) Print() {
	for _, e := range(t) {
		fmt.Println(e)
	}
}

func Schedule() {
	//ch := make(chan int)
	er := new(echoRunner)
	tasks := TaskSlice{{1, 5, *er}, {0, 2, *er}}
	tasks.Print()
	sort.Sort(tasks)
	tasks.Print()
}

func main() {
	Schedule()
}

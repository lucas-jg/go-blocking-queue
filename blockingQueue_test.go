package queue

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := New()

	go printMemStats()

	cnt := 102817
	c := make(chan int)

	go enqueueLoop(cnt, q, c)
	time.Sleep(100 * time.Millisecond)
	go dequeueLoop(cnt, q, c)

	e, d := <-c, <-c

	fmt.Printf("Result : %d %d", e, d)
	time.Sleep(6000 * time.Second)
}

func enqueueLoop(cnt int, q *Queue, c chan int) {
	for i := 0; i < cnt; i++ {
		q.Enqueue("Queue Test !!! n")
	}

	c <- cnt
}

func dequeueLoop(cnt int, q *Queue, c chan int) {
	for i := 0; i < cnt; i++ {
		q.Dequeue()
		//fmt.Printf("cnt %d ,, [%v]\n", i, q.Dequeue())
	}

	c <- cnt
}

func printMemStats() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v   TotalAlloc = %v   Sys = %v   NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
		time.Sleep(5 * time.Second)
	}
}

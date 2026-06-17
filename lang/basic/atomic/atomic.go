package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type atomicIntMutex struct {
	value int
	lock  sync.Mutex
}

func (a *atomicIntMutex) increment() {
	fmt.Println("mutex increment")
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicIntMutex) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

type atomicInt struct {
	value int64
}

func (a *atomicInt) increment() {
	fmt.Println("atomic increment")
	atomic.AddInt64(&a.value, 1)
}

func (a *atomicInt) get() int64 {
	return atomic.LoadInt64(&a.value)
}

func main() {
	var a atomicIntMutex
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("mutex result:", a.get())

	var b atomicInt
	b.increment()
	go func() {
		b.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("atomic result:", b.get())
}

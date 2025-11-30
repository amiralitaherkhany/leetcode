package main

import (
	"fmt"
	"sync"
)

func main() {
	foo := NewFoo()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		foo.Second(func() { fmt.Print("second") })
	}()
	go func() {
		defer wg.Done()
		foo.First(func() { fmt.Print("first") })
	}()
	go func() {
		defer wg.Done()
		foo.Third(func() { fmt.Print("third") })
	}()
	wg.Wait()
}

type Foo struct {
	ch1 chan any
	ch2 chan any
}

func NewFoo() *Foo {
	return &Foo{
		make(chan any),
		make(chan any),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	close(f.ch1)
}

func (f *Foo) Second(printSecond func()) {
	<-f.ch1
	/// Do not change this line
	printSecond()
	close(f.ch2)
}

func (f *Foo) Third(printThird func()) {
	<-f.ch2
	// Do not change this line
	printThird()
}

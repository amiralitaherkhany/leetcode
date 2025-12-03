package main

import (
	"fmt"
	"sync"
)

func main() {
	fooBar := NewFooBar(3)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fooBar.Foo(func() { fmt.Print("foo") })
	}()
	go func() {
		defer wg.Done()
		fooBar.Bar(func() { fmt.Print("bar") })
	}()

	wg.Wait()
}

type FooBar struct {
	n          int
	fooChannel chan bool
	barChannel chan bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:          n,
		fooChannel: make(chan bool),
		barChannel: make(chan bool),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		if i == 0 {
			printFoo()
			fb.fooChannel <- true
			continue
		}
		<-fb.barChannel
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.fooChannel <- true
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooChannel
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		if i == fb.n-1 {
			close(fb.barChannel)
			return
		}
		fb.barChannel <- true
	}
}

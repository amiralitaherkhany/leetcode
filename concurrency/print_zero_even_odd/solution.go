package main

import (
	"fmt"
	"sync"
)

func main() {
	obj := NewZeroEvenOdd(3)

	wg := &sync.WaitGroup{}

	wg.Go(func() {
		obj.Zero(func(i int) { fmt.Print(i) })
	})
	wg.Go(func() {
		obj.Even(func(i int) { fmt.Print(i) })
	})
	wg.Go(func() {
		obj.Odd(func(i int) { fmt.Print(i) })
	})

	wg.Wait()
}

type ZeroEvenOdd struct {
	n        int
	zeroChan chan int
	oddChan  chan int
	evenChan chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:        n,
		zeroChan: make(chan int),
		oddChan:  make(chan int),
		evenChan: make(chan int),
	}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		printNumber(0)
		if i%2 == 0 {
			z.evenChan <- i
		} else {
			z.oddChan <- i
		}
		<-z.zeroChan
	}
	close(z.evenChan)
	close(z.oddChan)
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := range z.evenChan {
		printNumber(i)
		z.zeroChan <- 0
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := range z.oddChan {
		printNumber(i)
		z.zeroChan <- 0
	}
}

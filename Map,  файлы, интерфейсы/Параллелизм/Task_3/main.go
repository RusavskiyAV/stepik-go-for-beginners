package main

import "sync"

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		var count int
		wg := new(sync.WaitGroup)

		for i := 0; i < n; i++ {
			a := <-in1
			b := <-in2

			wg.Add(1)

			go func(wg *sync.WaitGroup, a, b, i int) {
				defer wg.Done()

				var c, d int
				wg2 := new(sync.WaitGroup)
				wg2.Add(2)

				go func() {
					defer wg2.Done()
					c = fn(a)
				}()

				go func() {
					defer wg2.Done()
					d = fn(b)
				}()

				wg2.Wait()

				for count != i {
				}

				out <- c + d
				count++
			}(wg, a, b, i)
		}
	}()
}

func main() {
}

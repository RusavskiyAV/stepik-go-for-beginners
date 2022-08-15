package main

func work() {
}

func main() {
	done := make(chan struct{})

	go func() {
		work()
		close(done)
	}()

	<-done
}

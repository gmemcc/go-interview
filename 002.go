package main

import "time"

func main() {
	nums := []int{1, 2, 3}
	continueCh := make(chan struct{}, 3)
	for _, num := range nums {
		go func() {
			<-continueCh
			print(num)
		}()
		continueCh <- struct{}{}
	}
	time.Sleep(time.Second)
}

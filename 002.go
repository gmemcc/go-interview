package main

func main() {
	nums := []int{1, 2, 3}
	continueCh := make(chan struct{}, 3)
	for _, num := range nums {
		go func() {
			continueCh <- struct{}{}
			print(num)
		}()
		<-continueCh
	}
}

package main

import "time"

func main() {
	go func() {
		time.Sleep(time.Second)
		print(1)
	}()
}

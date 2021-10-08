package main

func main() {
	oldNums := []int{1, 2, 3}
	var newNums []*int
	for _, num := range oldNums {
		newNums = append(newNums, &num)
	}
	for _, num := range newNums {
		print(*num)
	}
}

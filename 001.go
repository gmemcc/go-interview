package main

func main() {
	oldNums := []int{1, 2, 3}
	newNums := toPtrs(oldNums)
	for _, num := range newNums {
		print(*num)
	}
}

func toPtrs(nums []int) []*int {
	var newNums []*int
	for _, num := range nums {
		newNums = append(newNums, &num)
	}
	return newNums
}

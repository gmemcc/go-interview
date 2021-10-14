package main

func main() {
	defer newDuck().quack()
	print(3)
}

type duck int

func (d duck) quack() {
	print(d)
}

func newDuck() duck {
	print(1)
	return duck(2)
}

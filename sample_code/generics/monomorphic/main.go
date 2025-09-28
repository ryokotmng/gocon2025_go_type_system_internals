package main

func Add[T int | int64](a, b T) T {
	return a + b
}

func main() {
	_ = Add(1, 2) // int
	_ = Add[int64](3, 4)
}

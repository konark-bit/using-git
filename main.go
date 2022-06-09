package main

// let us try to write a function that returns the sum of height of two students
//int compute(int a ,int b)
//return a+b
func heightSum(heights []int) int {
	sum := 0
	for _, value := range heights {
		sum += value
	}
	return sum
}

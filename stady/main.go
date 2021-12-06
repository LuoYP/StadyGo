package main

import (
	"fmt"
)

func main() {
	const name = "name"
	fmt.Println("hello world " + name)
	hello := "hello"
	arr := []byte(hello)
	arr[0] = 'p'
	result := string(arr)
	fmt.Println(result)
	myMap := make(map[string]int)
	myMap["one"] = 1
	fmt.Println(myMap["one"])
	fmt.Println(len(myMap))
	delete(myMap, "one")
	i, ok := myMap["one"]
	if ok {
		fmt.Println("ok", i)
	} else {
		fmt.Println("err")
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("numbers = ", numbers)
	odd := addFilter(numbers, isOdd)
	fmt.Println("numbers = ", odd)
	even := addFilter(numbers, isEven)
	fmt.Println("numbers = ", even)
}

type filterFunc func(int) bool

func addFilter(numbers []int, fun filterFunc) []int {
	var result []int
	for _, value := range numbers {
		if fun(value) {
			result = append(result, value)
		}
	}
	return result
}

func isOdd(number int) bool {
	return number%2 == 0
}

func isEven(number int) bool {
	return number%2 != 0
}

package main

import "math/rand"
import "time"
import "fmt"

type fn func([]int)

func make_int_array(number int) []int {
	rand.Seed(int64(time.Now().Nanosecond()))
	result := make([]int, number, number)
	for i, _ := range result {
		result[i] = rand.Intn(100)
	}
	return result
}

func result_helper(f fn, name string, input []int) {
	fmt.Printf("Name: %s\nInput: %v\n", name, input)
	now := time.Now()
	f(input)
	fmt.Printf("Result %v\nTook: %s\n", input, time.Since(now))
}

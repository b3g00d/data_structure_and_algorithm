package main

import "math/rand"
import "time"
import "fmt"
import "sync"

type fn_p func([]int, *sync.WaitGroup)
type fn func([]int)

func make_int_array(number int) []int {
	rand.Seed(int64(time.Now().Nanosecond()))
	result := make([]int, number, number)
	for i, _ := range result {
		result[i] = rand.Intn(100000000)
	}
	return result
}

func result_helper(f fn, name string, input []int) {
	fmt.Printf("Name: %s\n", name)
	now := time.Now()
	f(input)
	fmt.Printf("Took: %s\n", time.Since(now))
}

func result_helper_parallel(f fn_p, name string, input []int, wg *sync.WaitGroup) {
	fmt.Printf("Name: %s\nInput: %v\n", name, input)
	now := time.Now()
	f(input, wg)
	fmt.Printf("Result %v\nTook: %s\n", input, time.Since(now))
}

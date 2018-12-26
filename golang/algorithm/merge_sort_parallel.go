package main

import "sync"

func merge_sort_parallel(arr []int) {
	total := len(arr)
	start := 0
	if len(arr) > 1 {
		mid := int(total / 2)

		l_half := make([]int, len(arr[start:mid]))
		copy(l_half, arr[start:mid])

		r_half := make([]int, len(arr[mid:]))
		copy(r_half, arr[mid:])
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()
			merge_sort_parallel(l_half)
		}()
		go func() {
			defer wg.Done()
			merge_sort_parallel(r_half)
		}()

		wg.Wait()
		merge_parallel(arr, l_half, r_half)


	}
}
func merge_parallel(arr []int, l_half []int, r_half []int) {
	i, j, k := 0, 0, 0
	for ; i < len(l_half) && j < len(r_half); k++ {
		if l_half[i] > r_half[j] {
			arr[k] = r_half[j]
			j++
		} else {
			arr[k] = l_half[i]
			i++
		}
	}

	for ; i < len(l_half); i++ {
		arr[k] = l_half[i]
		k++
	}
	for ; j < len(r_half); j++ {
		arr[k] = r_half[j]
		k++
	}
}

func main() {
	new_arr := make_int_array(10000000)
	result_helper(merge_sort_parallel, "Merge Sort Parallel", new_arr)

}

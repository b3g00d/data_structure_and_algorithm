package main

func selection_sort(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			min := i
			for j := i + 1; j < len(arr); j++ {
				if arr[j] < arr[min] {
					min = j
				}
			}
			if min == i {
				continue
			}
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

func main() {
	new_arr := make_int_array(20)
	result_helper(selection_sort, "Selection Sort", new_arr)
}

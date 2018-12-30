package main

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	last_min := low
	for i := low + 1; i < high + 1; i++ {
		if arr[i] <= pivot {
			last_min++
			arr[last_min], arr[i] = arr[i], arr[last_min]

		}
	}
	arr[last_min], arr[low] = arr[low], arr[last_min]
	return last_min
}

func quick_sort(arr []int, low int, high int) {
	if len(arr) > 1 {
		if low < high {
			pi := partition(arr, low, high)
				
			quick_sort(arr, low, pi-1)
			quick_sort(arr, pi+1, high)

		}
	}
}

func quick_sort_helper(arr []int) {
	high := len(arr) - 1
	low := 0
	quick_sort(arr, low, high)
}

func main() {
	new_arr := make_int_array(10000000)
	result_helper(quick_sort_helper, "Quick Sort", new_arr)
}

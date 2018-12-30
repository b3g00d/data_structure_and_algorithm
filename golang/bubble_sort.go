package main

func buble_sort(arr []int) []int {
	if len(arr) > 1 {

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}

		}
	}
}

func main() {
	new_arr := make_int_array(20)
	result_helper(buble_sort, "Buble Sort", new_arr)

}

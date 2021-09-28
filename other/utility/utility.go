package utility

//快速排序  时间复杂度为 O(nlgn).
func SortFastInt(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	low := []int{}
	height := []int{}
	mid := []int{}
	l := len(arr)
	mid = append(mid, arr[0])

	for i := 1; i < l; i++ {
		if arr[i] < arr[0] {
			low = append(low, arr[i])
		}
		if arr[i] > arr[0] {
			height = append(height, arr[i])
		}
		if arr[i] == arr[0] {
			mid = append(mid, arr[i])
		}
	}

	low = SortFastInt(low)
	height = SortFastInt(height)

	myarr := append(append(low, mid...), height...)
	return myarr
}

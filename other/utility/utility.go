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

//选择  如果 a =零值 返回 f 否则返回t   //aabb
func Select[B any](a interface{}, t, f B) B {

	/* if !a {
		return f
	} else {
		return t
	}

	return t */

	switch value := a.(type) {
	case int:
		return inttf(value, t, f)
	case int8:
		return inttf(int(value), t, f)
	case int16:
		return inttf(int(value), t, f)
	case int32:
		return inttf(int(value), t, f)
	case int64:
		return inttf(int(value), t, f)
	case uint:
		return inttf(int(value), t, f)
	case string:
		if value == "" {
			return f
		} else {
			return t
		}
	case uint8:
		return inttf(int(value), t, f)
	case uint16:
		return inttf(int(value), t, f)
	case uint32:
		return inttf(int(value), t, f)
	case uint64:
		return inttf(int(value), t, f)
	case float32:
		if value == 0 {
			return f
		} else {
			return t
		}
	case float64:
		if value == 0 {
			return f
		} else {
			return t
		}
	case bool:
		if !value {
			return f
		} else {
			return t
		}
	default:
		if value == nil {
			return f
		} else {
			return t
		}
	}
}

//cc
func inttf[B any](value int, t, f B) B {
	if value == 0 {
		return f
	} else {
		return t
	}
}

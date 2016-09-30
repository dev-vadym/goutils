package utilx


func EnsureIntArrayNotNull(slice []int) []int{
	if slice == nil{
		slice = make([]int, 0)
	}
	return slice
}

func AppendIntArray(slice []int, elems ...int) []int {
	slice = EnsureIntArrayNotNull(slice)
	slice = append(slice, elems...)
	return slice
}

func EnsureStringArrayNotNull(slice []string) []string{
	if slice == nil{
		slice = make([]string, 0)
	}
	return slice
}

func AppendStringArray(slice []string, elems ...string) []string {
	slice = EnsureStringArrayNotNull(slice)
	slice = append(slice, elems...)
	return slice
}

func PageSliceIndex(arrLen int, page int, pageSize int) (int, int) {
	if page < 1{
		page = 1
	}
	if arrLen == 0{
		return 0, 0
	}
	max := arrLen - 1
	x := (page - 1) * pageSize
	y := x + pageSize
	if x > max{
		x = max
		y = max
	}
	if y > arrLen{
		y = arrLen
	}
	return x, y
}
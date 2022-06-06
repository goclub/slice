package xslice

// IndexOfFormIndex
// IndexOfFormIndex([]string{"a","b","c"}, "b", 0) // 1, true
// IndexOfFormIndex([]string{"a","b","c"}, "b", 2) // 0, false
// IndexOfFormIndex([]string{"a","b","c"}, "d", 0) // 0, false
func IndexOfFormIndex[T comparable](slice []T, value T, fromIndex uint64) (index uint64, found bool) {
	sliceLen := uint64(len(slice))
	i := fromIndex
	for ; i < sliceLen; i++ {
		item := slice[i]
		if item == value {
			found = true
			index = i
			return
		}
	}
	return
}
func IndexOf[T comparable](slice []T, value T) (index uint64, found bool) {
	return IndexOfFormIndex(slice, value, 0)
}

// ContainsFormIndex
// ContainsFormIndex([]string{"a","b","c"}, "b", 0) // true
// ContainsFormIndex([]string{"a","b","c"}, "b", 2) // false
// ContainsFormIndex([]string{"a","b","c"}, "d", 0) // false
func ContainsFormIndex[T comparable](slice []T, value T, fromIndex uint64) (found bool) {
	_, found = IndexOfFormIndex(slice, value, fromIndex)
	return
}
func Contains[T comparable](slice []T, value T) (found bool) {
	return ContainsFormIndex(slice, value, 0)
}

// Intersection
// Intersection([]int{1,0,8, 2}, []int{3,0,8, 4}) // []int{0,8}
func Intersection[T comparable](slices ...[]T) (unorderedResult []T) {
	slicesLen := uint64(len(slices))
	if slicesLen == 1 {
		return slices[0]
	}
	count := map[T]uint64{}
	for _, slice := range slices {
		for _, item := range slice {
			count[item]++
		}
	}
	for key, value := range count {
		if value == slicesLen {
			unorderedResult = append(unorderedResult, key)
		}
	}
	return
}

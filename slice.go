package xslice

// IndexOfFormIndex
// IndexOfFormIndex([]string{"a","b","c"}, "b", 0) // true, 1
// IndexOfFormIndex([]string{"a","b","c"}, "b", 2) // false, 0
// IndexOfFormIndex([]string{"a","b","c"}, "d", 0) // false, 0
func IndexOfFormIndex[T comparable](slice []T, value T, fromIndex uint64) (found bool, index uint64) {
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
func IndexOf[T comparable](slice []T, value T) (found bool, index uint64) {
	return IndexOfFormIndex(slice, value, 0)
}

// ContainsFormIndex
// ContainsFormIndex([]string{"a","b","c"}, "b", 0) // true
// ContainsFormIndex([]string{"a","b","c"}, "b", 2) // false
// ContainsFormIndex([]string{"a","b","c"}, "d", 0) // false
func ContainsFormIndex[T comparable](slice []T, value T, fromIndex uint64) (found bool) {
	found, _ = IndexOfFormIndex(slice, value, fromIndex)
	return
}
func Contains[T comparable](slice []T, value T) (found bool) {
	return ContainsFormIndex(slice, value, 0)
}

// Intersection
// Intersection([]int{1,0,8, 2}, []int{3,0,8, 4}) // []int{0,8}
func Intersection[T comparable](slices ...[]T) (result []T) {
	slicesLen := uint64(len(slices))
	if slicesLen == 0 {
		return nil
	}
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
			result = append(result, key)
		}
	}
	return
}

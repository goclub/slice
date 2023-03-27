package xslice

// IndexOfFormIndex 函数接收一个字符串切片 slice、一个字符串 s 和一个起始索引 startIndex，
// 返回 s 在 slice 中第一次出现的索引值和一个布尔值。如果找到了 s，则返回它的索引和 true。
// 如果没有找到 s，则返回 0 和 false。
//
// 参数：
//   - slice: 一个字符串切片，包含要搜索的元素。
//   - s: 要查找的字符串。
//   - startIndex: 搜索的起始索引。
//
// 返回值：
//   - 如果找到了 s，则返回它在 slice 中第一次出现的索引和 true。
//   - 如果没有找到 s，则返回 0 和 false。
// 示例：
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

// IndexOf 函数接收一个泛型类型的切片 slice 和一个泛型类型的值 value，
// 返回 value 在 slice 中第一次出现的索引值和一个布尔值。如果找到了 value，
// 则返回它的索引和 true。如果没有找到 value，则返回 0 和 false。
//
// 参数：
//   - slice: 一个泛型类型的切片，包含要搜索的元素。
//   - value: 要查找的值。
//
// 返回值：
//   - 如果找到了 value，则返回它在 slice 中第一次出现的索引和 true。
//   - 如果没有找到 value，则返回 0 和 false。
//
// 注意：
// IndexOf 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
// 示例：
//   i, found := IndexOf([]string{"foo", "bar", "baz"}, "bar")
//   // i == 1, found == true
//   i, found := IndexOf([]int{1, 2, 3, 4, 5}, 6)
//   // i == 0, found == false
func IndexOf[T comparable](slice []T, value T) (index uint64, found bool) {
	return IndexOfFormIndex(slice, value, 0)
}

// ContainsFormIndex 函数接收一个泛型类型的切片 slice、一个泛型类型的值 value，
// 和一个起始索引 fromIndex，返回 value 是否在 slice 中从 fromIndex 开始的位置中出现。
//
// 参数：
//   - slice: 一个泛型类型的切片，包含要搜索的元素。
//   - value: 要查找的值。
//   - fromIndex: 搜索的起始索引。
//
// 返回值：
//   - 如果 value 在从 fromIndex 开始的位置中出现，则返回 true。
//   - 如果 value 没有在从 fromIndex 开始的位置中出现，则返回 false。
//
// 注意：
// ContainsFormIndex 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
//
// 示例：
//   found := ContainsFormIndex([]string{"foo", "bar", "baz"}, "bar", 0)
//   // found == true
//   found := ContainsFormIndex([]int{1, 2, 3, 4, 5}, 6, 0)
//   // found == false
func ContainsFormIndex[T comparable](slice []T, value T, fromIndex uint64) (found bool) {
	_, found = IndexOfFormIndex(slice, value, fromIndex)
	return
}

// Contains 函数接收一个泛型类型的切片 slice 和一个泛型类型的值 value，
// 返回 value 是否在 slice 中出现。
//
// 参数：
//   - slice: 一个泛型类型的切片，包含要搜索的元素。
//   - value: 要查找的值。
//
// 返回值：
//   - 如果 value 在 slice 中出现，则返回 true。
//   - 如果 value 没有在 slice 中出现，则返回 false。
//
// 注意：
// Contains 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
//
// 示例：
//   found := Contains([]string{"foo", "bar", "baz"}, "bar")
//   // found == true
//   found := Contains([]int{1, 2, 3, 4, 5}, 6)
//   // found == false
func Contains[T comparable](slice []T, value T) (found bool) {
	return ContainsFormIndex(slice, value, 0)
}

// Union 函数接收一个或多个泛型类型的切片 slices，返回这些切片的并集（即去重后的所有元素）。
// 返回的结果无序。
//
// 参数：
//   - slices: 一个或多个泛型类型的切片，包含要合并的元素。
//
// 返回值：
//   - 返回一个新的泛型类型的切片 unorderedResult，包含所有 slices 中的去重后的元素。
//
// 注意：
// Union 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
//
// 示例：
//   unorderedResult := Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
//   // unorderedResult == []int{1, 2, 3, 4, 5}
func Union[T comparable](slices ...[]T) (unorderedResult []T) {
	unorderedResult = []T{}
	hash := make(map[T]struct{})
	for _, slice := range slices {
		for _, elem := range slice {
			if _, ok := hash[elem]; !ok {
				unorderedResult = append(unorderedResult, elem)
				hash[elem] = struct{}{}
			}
		}
	}
	return
}

// Intersection 函数接收一个或多个泛型类型的切片 slices，返回这些切片的交集（即所有切片中都存在的元素）。
// 返回的结果无序。
//
// 参数：
//   - slices: 一个或多个泛型类型的切片，包含要计算交集的元素。
//
// 返回值：
//   - 返回一个新的泛型类型的切片 unorderedResult，包含所有 slices 中的交集元素。
//
// 注意：
// Intersection 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
//
// 示例：
//   unorderedResult := Intersection([]int{1, 0, 8, 2}, []int{3, 0, 8, 4})
//   // unorderedResult == []int{0, 8}
func Intersection[T comparable](slices ...[]T) (unorderedResult []T) {
	unorderedResult = []T{}
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

// Difference
// 函数计算输入切片 slices 的差集，即在第一个切片中出现，
// 但在其余所有切片中都没有出现的元素。如果 slices 为空，则返回一个空切片。
//
// 参数：
// slices：一个或多个切片，用于计算差集。
//
// 返回值：
// result：一个切片，包含差集元素。如果差集为空，则返回一个空切片。
//
// 示例：
// Difference([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5}) // []int{1}
func Difference[T comparable](slices ...[]T) (result []T) {
	result = []T{}
	if len(slices) == 0 {
		return
	}
	hash := make(map[T]struct{})
	for _, slice := range slices[1:] {
		for _, elem := range slice {
			hash[elem] = struct{}{}
		}
	}
	for _, elem := range slices[0] {
		if _, ok := hash[elem]; !ok {
			result = append(result, elem)
		}
	}
	return
}

// Unique 函数接收一个泛型类型的切片 slice，返回一个去重后的新切片 unique，
// 该切片中只包含 slice 中的唯一元素。
//
// 参数：
//   - slice: 一个泛型类型的切片，包含要去重的元素。
//
// 返回值：
//   - 返回一个新的泛型类型的切片 unique，包含 slice 中的唯一元素。
//
// 注意：
// Unique 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
//
// 示例：
//   unique := Unique([]int{1, 2, 2, 3})
//   // unique == []int{1, 2, 3}
func Unique[T comparable](slice []T) (unique []T) {
	unique = []T{}
	hash := map[T]struct{}{}
	for _, v := range slice {
		if _, has := hash[v]; has == false {
			unique = append(unique, v)
			hash[v] = struct{}{}
		}
	}
	return
}

// First
// First([]string{"a","b","c"}) // "a", true;
// First([]string{}) // "", false
// var list []string
// First(list) // "", false
func First[T any](slice []T) (v T, has bool) {
	if len(slice) == 0 {
		return v, false
	}
	return slice[0], true
}

// Last 函数接收一个泛型类型的切片 slice，返回 slice 中的最后一个元素和一个布尔值。
// 如果 slice 不为空，则返回最后一个元素和 true。如果 slice 为空，则返回 T 类型的零值和 false。
//
// 参数：
//   - slice: 一个泛型类型的切片，包含要获取最后一个元素的元素。
//
// 返回值：
//   - 如果 slice 不为空，则返回 slice 中的最后一个元素和 true。
//   - 如果 slice 为空，则返回 T 类型的零值和 false。
//
// 注意：
// Last 函数使用了泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，
// 这是因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。
// 如果 T 类型不满足 comparable 接口，则编译时会产生错误。
//
// 示例：
//   last, found := Last([]string{"a", "b", "c"})
//   // last == "c", found == true
//   last, found := Last([]string{})
//   // last == "", found == false
//   var list []string
//   last, found := Last(list)
//   // last == "", found == false
func Last[T any](slice []T) (v T, has bool) {
	n := len(slice)
	if n == 0 {
		return v, false
	}
	return slice[n-1], true
}

package sl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import (
	"reflect"
	"sort"
)

func TestIndexOfFromIndex(t *testing.T) {
	source := []string{"a", "b", "c"}
	{
		index, found := IndexOf(source, "a")
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(0))
	}
	{
		index, found := IndexOfFromIndex(source, "a", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(0))
	}
	{
		index, found := IndexOfFromIndex(source, "a", 1)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
	{
		index, found := IndexOfFromIndex(source, "a", 2)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
	{
		index, found := IndexOfFromIndex(source, "b", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(1))
	}
	{
		index, found := IndexOfFromIndex(source, "b", 1)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(1))
	}
	{
		index, found := IndexOfFromIndex(source, "b", 2)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
	{
		index, found := IndexOfFromIndex(source, "c", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(2))
	}
	{
		index, found := IndexOfFromIndex(source, "c", 1)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(2))
	}
	{
		index, found := IndexOfFromIndex(source, "c", 2)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(2))
	}
	{
		index, found := IndexOfFromIndex(source, "c", 3)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
}
func TestIntersection(t *testing.T) {
	tt := []struct {
		name string
		ins  [][]string
		out  []string
	}{
		{
			"empty",
			[][]string{
				{},
			},
			[]string{},
		},
		{
			"1 1",
			[][]string{
				{"1"},
			},
			[]string{"1"},
		},
		{
			"1 2 | 1 2",
			[][]string{
				{"1", "2"},
			},
			[]string{"1", "2"},
		},
		{
			"1 2 | 1",
			[][]string{
				{"1", "2"},
				{"1"},
			},
			[]string{"1"},
		},
		{
			"1 2 | 1 | 1",
			[][]string{
				{"1", "2"},
				{"1"},
				{"1"},
			},
			[]string{"1"},
		},
		{
			"1 2 | 1 | 1 3",
			[][]string{
				{"1", "2"},
				{"1"},
				{"1", "3"},
			},
			[]string{"1"},
		},
		{
			"1 2 | 1 | 1 2 3 ",
			[][]string{
				{"1", "2"},
				{"1"},
				{"1", "2", "3"},
			},
			[]string{"1"},
		},
		{

			"1 2 | 1 2 | 1 2 3",
			[][]string{
				{"1", "2"},
				{"1", "2"},
				{"1", "2", "3"},
			},
			[]string{"1", "2"},
		},
		{
			"1 1 1 1 1 | 1 1 1 1 1 | 1 1 1 1 1",
			[][]string{
				{"1", "1", "1", "1", "1"},
				{"1", "1", "1", "1", "1"},
				{"1", "1", "1", "1", "1"},
			},
			[]string{"1"},
		},
		{
			"more repeat",
			[][]string{
				{"1", "1", "1", "1", "1", "2"},
				{"1", "1", "1", "1", "1", "2", "2"},
				{"1", "1", "1", "1", "1", "2", "2"},
			},
			[]string{"1", "2"},
		},
		{
			"shanghai",
			[][]string{
				{"310000", "310000", "310113"},
				{"110000"},
			},
			[]string(nil),
		},
	}

	for _, item := range tt {
		result := Intersection(item.ins...)
		sort.Strings(result)
		assert.Equal(t, result, item.out, item.name)
	}
}

func TestUniqueStrings(t *testing.T) {
	// Test case 1: slice with duplicate strings
	input1 := []string{"apple", "banana", "orange", "banana", "peach", "orange"}
	expected1 := []string{"apple", "banana", "orange", "peach"}
	output1 := Unique(input1)
	if !reflect.DeepEqual(expected1, output1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: empty slice of strings
	input2 := []string{}
	expected2 := []string{}
	output2 := Unique(input2)
	if !reflect.DeepEqual(expected2, output2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: slice with all Unique strings
	input3 := []string{"hello", "world", "this", "is", "a", "test"}
	expected3 := []string{"hello", "world", "this", "is", "a", "test"}
	output3 := Unique(input3)
	if !reflect.DeepEqual(expected3, output3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}
}

func TestUniqueInts(t *testing.T) {
	// Test case 1: slice with duplicate ints
	input1 := []int{1, 2, 3, 2, 4, 3, 5}
	expected1 := []int{1, 2, 3, 4, 5}
	output1 := Unique(input1)
	if !reflect.DeepEqual(expected1, output1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: empty slice of ints
	input2 := []int{}
	expected2 := []int{}
	output2 := Unique(input2)
	if !reflect.DeepEqual(expected2, output2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: slice with all unique ints
	input3 := []int{1, 2, 3, 4, 5, 0}
	expected3 := []int{1, 2, 3, 4, 5, 0}
	output3 := Unique(input3)
	if !reflect.DeepEqual(expected3, output3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: slice with negative and positive ints
	input4 := []int{-1, 2, 3, 2, -4, 3, 5}
	expected4 := []int{-1, 2, 3, -4, 5}
	output4 := Unique(input4)
	if !reflect.DeepEqual(expected4, output4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}

	// Test case 5: slice with all negative ints
	input5 := []int{-5, -4, -3, -2, -1}
	expected5 := []int{-5, -4, -3, -2, -1}
	output5 := Unique(input5)
	if !reflect.DeepEqual(expected5, output5) {
		t.Errorf("Test case 5 failed: expected %v, but got %v", expected5, output5)
	}
}
func TestUniqueFloats(t *testing.T) {
	// Test case 1: slice with duplicate floats
	input1 := []float64{1.1, 2.2, 3.3, 2.2, 4.4, 3.3, 5.5}
	expected1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	output1 := Unique(input1)
	if !reflect.DeepEqual(expected1, output1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: empty slice of floats
	input2 := []float64{}
	expected2 := []float64{}
	output2 := Unique(input2)
	if !reflect.DeepEqual(expected2, output2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: slice with all unique floats
	input3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expected3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	output3 := Unique(input3)
	if !reflect.DeepEqual(expected3, output3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: slice with negative and positive floats
	input4 := []float64{-1.1, 2.2, 3.3, 2.2, -4.4, 3.3, 5.5}
	expected4 := []float64{-1.1, 2.2, 3.3, -4.4, 5.5}
	output4 := Unique(input4)
	if !reflect.DeepEqual(expected4, output4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}

	// Test case 5: slice with all negative floats
	input5 := []float64{-5.5, -4.4, -3.3, -2.2, -1.1}
	expected5 := []float64{-5.5, -4.4, -3.3, -2.2, -1.1}
	output5 := Unique(input5)
	if !reflect.DeepEqual(expected5, output5) {
		t.Errorf("Test case 5 failed: expected %v, but got %v", expected5, output5)
	}
}

func TestUniqueAlias(t *testing.T) {
	type ID int
	ids := []ID{1, 2, 2, 3}
	assert.Equal(t, []ID{1, 2, 3}, Unique(ids))
}

func TestFirstStringSlice(t *testing.T) {
	// Test case 1: slice with one element
	input1 := []string{"a"}
	expected1 := "a"
	output1, has1 := First(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []string{"a", "b", "c"}
	expected2 := "a"
	output2, has2 := First(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []string{}
	expected3 := ""
	output3, has3 := First(input3)
	if has3 {
		t.Errorf("Test case 3 failed: expected has=false, but got has=true")
	}
	if output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}
	// Test case 4: empty string slice
	var s []string
	expected4 := ""
	output4, has4 := First(s)
	if has4 {
		t.Errorf("Test case 4 failed: expected has=false, but got has=true")
	}
	if output4 != expected4 {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestFirstIntSlice(t *testing.T) {
	// Test case 1: slice with one element
	input1 := []int{1}
	expected1 := 1
	output1, has1 := First(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []int{1, 2, 3}
	expected2 := 1
	output2, has2 := First(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []int{}
	expected3 := 0
	output3, has3 := First(input3)
	if has3 {
		t.Errorf("Test case 3 failed: expected has=false, but got has=true")
	}
	if output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	var s []int
	expected4 := 0
	output4, has4 := First(s)
	if has4 {
		t.Errorf("Test case 4 failed: expected has=false, but got has=true")
	}
	if output4 != expected4 {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestLastStringSlice(t *testing.T) {
	// Test case 1: slice with one element
	input1 := []string{"a"}
	expected1 := "a"
	output1, has1 := Last(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []string{"a", "b", "c"}
	expected2 := "c"
	output2, has2 := Last(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []string{}
	expected3 := ""
	output3, has3 := Last(input3)
	if has3 {
		t.Errorf("Test case 3 failed: expected has=false, but got has=true")
	}
	if output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}
}

func TestLastIntSlice(t *testing.T) {
	// Test case 1: slice with one element
	input1 := []int{1}
	expected1 := 1
	output1, has1 := Last(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []int{1, 2, 3}
	expected2 := 3
	output2, has2 := Last(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []int{}
	expected3 := 0
	output3, has3 := Last(input3)
	if has3 {
		t.Errorf("Test case 3 failed: expected has=false, but got has=true")
	}
	if output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}
}

func TestLastEmptyStringSlice(t *testing.T) {
	// Test case 4: empty string slice
	var s []string
	expected4 := ""
	output4, has4 := Last(s)
	if has4 {
		t.Errorf("Test case 4 failed: expected has=false, but got has=true")
	}
	if output4 != expected4 {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestLastEmptyIntSlice(t *testing.T) {
	// Test case 5: empty int slice
	var s []int
	expected5 := 0
	output5, has5 := Last(s)
	if has5 {
		t.Errorf("Test case 5 failed: expected has=false, but got has=true")
	}
	if output5 != expected5 {
		t.Errorf("Test case 5 failed: expected %v, but got %v", expected5, output5)
	}
}

func TestUnionStringSlice(t *testing.T) {
	// Test case 1: multiple non-empty slices
	input1 := [][]string{
		{"a", "b", "c"},
		{"b", "c", "d"},
		{"c", "d", "e"},
	}
	expected1 := []string{"a", "b", "c", "d", "e"}
	output1 := Union(input1...)
	sort.Strings(output1)
	if !reflect.DeepEqual(output1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: multiple empty slices
	input2 := [][]string{
		{},
		{},
		{},
	}
	expected2 := []string{}
	output2 := Union(input2...)
	sort.Strings(output2)
	if !reflect.DeepEqual(output2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: one non-empty slice and two empty slices
	input3 := [][]string{
		{"a", "b", "c"},
		{},
		{},
	}
	expected3 := []string{"a", "b", "c"}
	output3 := Union(input3...)
	sort.Strings(output3)
	if !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	var s []string
	expected4 := []string{}
	output4 := Union(s)
	sort.Strings(output4)
	if !reflect.DeepEqual(output4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestDifferenceIntSlice(t *testing.T) {
	// Test case 1: multiple non-empty slices
	input1 := [][]int{
		{1, 1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	expected1 := []int{1}
	output1 := Difference(input1...)
	sort.Ints(output1)
	if !reflect.DeepEqual(output1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: multiple empty slices
	assert.Equal(t, Difference([][]int{{}, {}, {}}...), []int(nil))

	// Test case 3: one non-empty slice and two empty slices
	input3 := [][]int{
		{1, 2, 3},
		{},
		{},
	}
	expected3 := []int{1, 2, 3}
	output3 := Difference(input3...)
	sort.Ints(output3)
	if !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	assert.Equal(t, Difference([]int(nil)), []int(nil))
	// Test case 5
	{
		input := [][]int{
			{1, 1, 2, 3},
			{3},
		}
		expected := []int{1, 2}
		output := Difference(input...)
		sort.Ints(output)
		assert.Equal(t, expected, output)

	}
}

func TestFilter(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	evenFilter := func(a int) bool {
		return a%2 == 0
	}
	evenSlice := Filter(intSlice, evenFilter)
	expectedEvenSlice := []int{2, 4}

	if !reflect.DeepEqual(evenSlice, expectedEvenSlice) {
		t.Errorf("Filter failed, expected %v but got %v", expectedEvenSlice, evenSlice)
	}

	stringSlice := []string{"Hello", "World", "Go"}
	lengthFilter := func(a string) bool {
		return len(a) > 3
	}
	longStringSlice := Filter(stringSlice, lengthFilter)
	expectedLongStringSlice := []string{"Hello", "World"}

	if !reflect.DeepEqual(longStringSlice, expectedLongStringSlice) {
		t.Errorf("Filter failed, expected %v but got %v", expectedLongStringSlice, longStringSlice)
	}
}

func TestFind(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	// Test finding an even number
	isEven := func(a int) bool { return a%2 == 0 }
	v, has := Find(intSlice, isEven)
	if !has || v != 2 {
		t.Errorf("Find failed, expected 2 but got %d", v)
	}

	// Test finding an odd number
	isOdd := func(a int) bool { return a%2 == 1 }
	v, has = Find(intSlice, isOdd)
	if !has || v != 1 {
		t.Errorf("Find failed, expected 1 but got %d", v)
	}

	// Test finding a number not in the slice
	isNegative := func(a int) bool { return a < 0 }
	v, has = Find(intSlice, isNegative)
	if has {
		t.Errorf("Find failed, expected false but got true")
	}
}

func TestEvery(t *testing.T) {
	// Test all elements are even
	isEven := func(a int) bool { return a%2 == 0 }
	allEven := Every([]int{2, 4, 6, 8, 10}, isEven)
	if !allEven {
		t.Errorf("Every failed, expected true but got false")
	}

	// Test some elements are odd
	allEven = Every([]int{2, 4, 6, 7, 10}, isEven)
	if allEven {
		t.Errorf("Every failed, expected false but got true")
	}

	// Test an empty slice
	allEven = Every([]int{}, isEven)
	if !allEven {
		t.Errorf("Every failed, expected true but got false")
	}
}

func TestSome(t *testing.T) {
	// 测试用例1：空切片
	slice1 := []int{}
	fn1 := func(v int) bool { return v > 0 }
	assert.Equal(t, Some(slice1, fn1), false)

	// 测试用例2：切片中有满足条件的元素
	slice2 := []string{"hello", "world", "golang"}
	fn2 := func(v string) bool { return len(v) == 5 }
	assert.Equal(t, Some(slice2, fn2), true)

	// 测试用例3：切片中没有满足条件的元素
	slice3 := []int{1, 2, 3, 4}
	fn3 := func(v int) bool { return v > 5 }
	assert.Equal(t, Some(slice3, fn3), false)
}
func TestShuffle(t *testing.T) {
	// 创建一个包含 10 个元素的整数切片。
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 1000; i++ {
		// 记录初始顺序。
		original := Copy(slice)

		// 打乱切片顺序。
		Shuffle(slice)

		// 检查切片是否被打乱。
		if reflect.DeepEqual(slice, original) {
			t.Errorf("Slice not shuffled: %v", slice)
		}

		// 检查切片中的元素是否保持不变。
		for _, elem := range original {
			if !Contains(slice, elem) {
				t.Errorf("Element missing from shuffled slice: %d", elem)
			}
		}
	}
}
func TestShuffle_2(t *testing.T) {
	// 创建一个包含 10 个元素的整数切片。
	{
		slice := []int{1}
		Shuffle(slice)
		assert.Equal(t, slice, []int{1})
	}
	{
		slice := []int{1, 2}
		Shuffle(slice)
	}
}

func TestPluck(t *testing.T) {
	// 测试用例1：空切片
	slice1 := []int{}
	fn1 := func(v int) int { return v * 2 }
	result1 := Pluck(slice1, fn1)
	assert.Equal(t, []int{}, result1)

	// 测试用例2：非空切片
	slice2 := []string{"hello", "world", "golang"}
	fn2 := func(v string) int { return len(v) }
	result2 := Pluck(slice2, fn2)
	expected2 := []int{5, 5, 6}
	assert.Equal(t, expected2, result2)

	// 测试用例3：结构体
	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Cathy", 22},
	}
	// 提取人名的函数
	extractName := func(p Person) string { return p.Name }
	// 使用 Pluck 函数提取 persons 中所有人的名字
	assert.Equal(t, Pluck(persons, extractName), []string{"Alice", "Bob", "Cathy"})
}

func TestCopy(t *testing.T) {
	// 测试用例1：空切片
	slice1 := []int{}
	result1 := Copy(slice1)
	if len(result1) != 0 {
		t.Errorf("Copy(%v) = %v, want %v", slice1, result1, []int{})
	}

	// 测试用例2：非空切片
	slice2 := []string{"hello", "world", "golang"}
	result2 := Copy(slice2)
	expected2 := []string{"hello", "world", "golang"}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Copy(%v) = %v, want %v", slice2, result2, expected2)
	}
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{2, 3, 4}
	output := Map(input, func(v int) int { return v + 1 })
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Map(%v, fn) = %v; want %v", input, output, expected)
	}
}

func TestSort(t *testing.T) {
	// 测试用例 1：对整数类型的切片进行排序
	input1 := []int{5, 2, 7, 1, 9}
	expected1 := []int{1, 2, 5, 7, 9}
	Sort(input1, func(a, b int) bool {
		return a < b
	})
	if !reflect.DeepEqual(input1, expected1) {
		t.Errorf("Sort(%v, less) = %v; want %v", input1, input1, expected1)
	}

	// 测试用例 2：对字符串类型的切片进行排序
	input2 := []string{"apple", "pear", "banana", "orange"}
	expected2 := []string{"apple", "banana", "orange", "pear"}
	Sort(input2, func(a, b string) bool {
		return a < b
	})
	if !reflect.DeepEqual(input2, expected2) {
		t.Errorf("Sort(%v, less) = %v; want %v", input2, input2, expected2)
	}
}

func TestGroup(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := map[uint64][]int{
		0: {2, 4},
		1: {1, 3, 5},
	}
	output := Group(input, func(v int) uint64 { return uint64(v % 2) })
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Group(%v, fn) = %v; want %v", input, output, expected)
	}
}
func TestIndex(t *testing.T) {
	{
		// 构造一个元素类型为 string 的切片。
		slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
		// 定义一个从元素中提取键的函数。
		fn := func(v string) rune {
			return rune(v[0])
		}
		// 调用 Index 函数将切片转换为映射。
		index := Index(slice, fn)
		// 检查映射是否包含预期的键值对。
		assert.Equal(t, "apple", index['a'])
		assert.Equal(t, "banana", index['b'])
		assert.Equal(t, "cherry", index['c'])
		assert.Equal(t, "date", index['d'])
		assert.Equal(t, "elderberry", index['e'])
	}
	{
		type Data struct {
			ID   uint64
			Name string
		}
		list := []Data{
			{1, "nimo"},
			{2, "nico"},
			{3, "tim"},
		}
		index := Index(list, func(v Data) uint64 {
			return v.ID
		})
		assert.Equal(t, "nimo", index[1].Name)
		assert.Equal(t, "nico", index[2].Name)
		assert.Equal(t, "tim", index[3].Name)
		assert.Equal(t, 3, len(index))
	}
	{
		type Data struct {
			ID   uint64
			Name string
		}
		list := []Data{
			{1, "nimo"},
			{2, "nico"},
			{3, "tim"},
			{2, "nico-2"},
		}
		index := Index(list, func(v Data) uint64 {
			return v.ID
		})
		assert.Equal(t, "nimo", index[1].Name)
		assert.Equal(t, "nico-2", index[2].Name)
		assert.Equal(t, "tim", index[3].Name)
		assert.Equal(t, 3, len(index))
	}
}
func TestRandElem(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < 1000; i++ {
		elem, has := RandElem(slice)
		if !has {
			t.Error("RandElem returned false for a non-empty slice")
		}

		if elem < 1 || elem > 5 {
			t.Errorf("RandElem returned an out-of-range value: %d", elem)
		}
	}

	emptySlice := []int{}
	_, emptyHas := RandElem(emptySlice)

	if emptyHas {
		t.Error("RandElem returned true for an empty slice")
	}
}

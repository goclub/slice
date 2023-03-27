package xslice_test

import (
	xslice "github.com/goclub/slice"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func TestIndexOfFromIndex(t *testing.T) {
	source := []string{"a", "b", "c"}
	{
		index, found := xslice.IndexOf(source, "a")
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(0))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "a", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(0))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "a", 1)
		assert.Equal(t, found, false)
		assert.Equal(t, index, int(0))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "a", 2)
		assert.Equal(t, found, false)
		assert.Equal(t, index, int(0))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "b", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(1))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "b", 1)
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(1))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "b", 2)
		assert.Equal(t, found, false)
		assert.Equal(t, index, int(0))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "c", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(2))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "c", 1)
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(2))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "c", 2)
		assert.Equal(t, found, true)
		assert.Equal(t, index, int(2))
	}
	{
		index, found := xslice.IndexOfFromIndex(source, "c", 3)
		assert.Equal(t, found, false)
		assert.Equal(t, index, int(0))
	}
}
func TestIntersection(t *testing.T) {
	{
		result := xslice.Intersection(
			[]int{1},
		)
		assert.Equal(t, result, []int{1})
	}
	{
		result := xslice.Intersection(
			[]int{1, 111, 222, 333, 2, 3},
			[]int{4, 111, 222, 333, 5, 6},
		)
		assert.Equal(t, len(result), 3)
		xslice.Contains(result, 111)
		xslice.Contains(result, 222)
		xslice.Contains(result, 333)
	}
	{
		result := xslice.Intersection(
			[]int{1, 111, 222, 333, 2, 3},
			[]int{4, 111, 222, 333, 5, 6},
			[]int{7, 111, 222, 333, 8, 9},
		)
		assert.Equal(t, len(result), 3)
		xslice.Contains(result, 111)
		xslice.Contains(result, 222)
		xslice.Contains(result, 333)
	}
}

func TestUniqueStrings(t *testing.T) {
	// Test case 1: slice with duplicate strings
	input1 := []string{"apple", "banana", "orange", "banana", "peach", "orange"}
	expected1 := []string{"apple", "banana", "orange", "peach"}
	output1 := xslice.Unique(input1)
	if !reflect.DeepEqual(expected1, output1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: empty slice of strings
	input2 := []string{}
	expected2 := []string{}
	output2 := xslice.Unique(input2)
	if !reflect.DeepEqual(expected2, output2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: slice with all xslice.Unique strings
	input3 := []string{"hello", "world", "this", "is", "a", "test"}
	expected3 := []string{"hello", "world", "this", "is", "a", "test"}
	output3 := xslice.Unique(input3)
	if !reflect.DeepEqual(expected3, output3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}
}

func TestUniqueInts(t *testing.T) {
	// Test case 1: slice with duplicate ints
	input1 := []int{1, 2, 3, 2, 4, 3, 5}
	expected1 := []int{1, 2, 3, 4, 5}
	output1 := xslice.Unique(input1)
	if !reflect.DeepEqual(expected1, output1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: empty slice of ints
	input2 := []int{}
	expected2 := []int{}
	output2 := xslice.Unique(input2)
	if !reflect.DeepEqual(expected2, output2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: slice with all unique ints
	input3 := []int{1, 2, 3, 4, 5, 0}
	expected3 := []int{1, 2, 3, 4, 5, 0}
	output3 := xslice.Unique(input3)
	if !reflect.DeepEqual(expected3, output3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: slice with negative and positive ints
	input4 := []int{-1, 2, 3, 2, -4, 3, 5}
	expected4 := []int{-1, 2, 3, -4, 5}
	output4 := xslice.Unique(input4)
	if !reflect.DeepEqual(expected4, output4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}

	// Test case 5: slice with all negative ints
	input5 := []int{-5, -4, -3, -2, -1}
	expected5 := []int{-5, -4, -3, -2, -1}
	output5 := xslice.Unique(input5)
	if !reflect.DeepEqual(expected5, output5) {
		t.Errorf("Test case 5 failed: expected %v, but got %v", expected5, output5)
	}
}
func TestUniqueFloats(t *testing.T) {
	// Test case 1: slice with duplicate floats
	input1 := []float64{1.1, 2.2, 3.3, 2.2, 4.4, 3.3, 5.5}
	expected1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	output1 := xslice.Unique(input1)
	if !reflect.DeepEqual(expected1, output1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: empty slice of floats
	input2 := []float64{}
	expected2 := []float64{}
	output2 := xslice.Unique(input2)
	if !reflect.DeepEqual(expected2, output2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: slice with all unique floats
	input3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expected3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	output3 := xslice.Unique(input3)
	if !reflect.DeepEqual(expected3, output3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: slice with negative and positive floats
	input4 := []float64{-1.1, 2.2, 3.3, 2.2, -4.4, 3.3, 5.5}
	expected4 := []float64{-1.1, 2.2, 3.3, -4.4, 5.5}
	output4 := xslice.Unique(input4)
	if !reflect.DeepEqual(expected4, output4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}

	// Test case 5: slice with all negative floats
	input5 := []float64{-5.5, -4.4, -3.3, -2.2, -1.1}
	expected5 := []float64{-5.5, -4.4, -3.3, -2.2, -1.1}
	output5 := xslice.Unique(input5)
	if !reflect.DeepEqual(expected5, output5) {
		t.Errorf("Test case 5 failed: expected %v, but got %v", expected5, output5)
	}
}

func TestUniqueAlias(t *testing.T) {
	type ID int
	ids := []ID{1, 2, 2, 3}
	assert.Equal(t, []ID{1, 2, 3}, xslice.Unique(ids))
}

func TestFirstStringSlice(t *testing.T) {
	// Test case 1: slice with one element
	input1 := []string{"a"}
	expected1 := "a"
	output1, has1 := xslice.First(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []string{"a", "b", "c"}
	expected2 := "a"
	output2, has2 := xslice.First(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []string{}
	expected3 := ""
	output3, has3 := xslice.First(input3)
	if has3 {
		t.Errorf("Test case 3 failed: expected has=false, but got has=true")
	}
	if output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}
	// Test case 4: empty string slice
	var s []string
	expected4 := ""
	output4, has4 := xslice.First(s)
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
	output1, has1 := xslice.First(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []int{1, 2, 3}
	expected2 := 1
	output2, has2 := xslice.First(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []int{}
	expected3 := 0
	output3, has3 := xslice.First(input3)
	if has3 {
		t.Errorf("Test case 3 failed: expected has=false, but got has=true")
	}
	if output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	var s []int
	expected4 := 0
	output4, has4 := xslice.First(s)
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
	output1, has1 := xslice.Last(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []string{"a", "b", "c"}
	expected2 := "c"
	output2, has2 := xslice.Last(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []string{}
	expected3 := ""
	output3, has3 := xslice.Last(input3)
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
	output1, has1 := xslice.Last(input1)
	if !has1 || output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: slice with multiple elements
	input2 := []int{1, 2, 3}
	expected2 := 3
	output2, has2 := xslice.Last(input2)
	if !has2 || output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: empty slice
	input3 := []int{}
	expected3 := 0
	output3, has3 := xslice.Last(input3)
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
	output4, has4 := xslice.Last(s)
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
	output5, has5 := xslice.Last(s)
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
	output1 := xslice.Union(input1...)
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
	output2 := xslice.Union(input2...)
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
	output3 := xslice.Union(input3...)
	sort.Strings(output3)
	if !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	var s []string
	expected4 := []string{}
	output4 := xslice.Union(s)
	sort.Strings(output4)
	if !reflect.DeepEqual(output4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestUnionFloat64Slice(t *testing.T) {
	// Test case 1: multiple non-empty slices
	input1 := [][]float64{
		{1.1, 2.2, 3.3},
		{2.2, 3.3, 4.4},
		{3.3, 4.4, 5.5},
	}
	expected1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	output1 := xslice.Union(input1...)
	sort.Float64s(output1)
	if !reflect.DeepEqual(output1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: multiple empty slices
	input2 := [][]float64{
		{},
		{},
		{},
	}
	expected2 := []float64{}
	output2 := xslice.Union(input2...)
	sort.Float64s(output2)
	if !reflect.DeepEqual(output2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: one non-empty slice and two empty slices
	input3 := [][]float64{
		{1.1, 2.2, 3.3},
		{},
		{},
	}
	expected3 := []float64{1.1, 2.2, 3.3}
	output3 := xslice.Union(input3...)
	sort.Float64s(output3)
	if !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	var s []float64
	expected4 := []float64{}
	output4 := xslice.Union(s)
	sort.Float64s(output4)
	if !reflect.DeepEqual(output4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestDifferenceIntSlice(t *testing.T) {
	// Test case 1: multiple non-empty slices
	input1 := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	expected1 := []int{1}
	output1 := xslice.Difference(input1...)
	sort.Ints(output1)
	if !reflect.DeepEqual(output1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, output1)
	}

	// Test case 2: multiple empty slices
	input2 := [][]int{
		{},
		{},
		{},
	}
	expected2 := []int{}
	output2 := xslice.Difference(input2...)
	sort.Ints(output2)
	if !reflect.DeepEqual(output2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, output2)
	}

	// Test case 3: one non-empty slice and two empty slices
	input3 := [][]int{
		{1, 2, 3},
		{},
		{},
	}
	expected3 := []int{1, 2, 3}
	output3 := xslice.Difference(input3...)
	sort.Ints(output3)
	if !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, output3)
	}

	// Test case 4: variable definition with no elements
	var s []int
	expected4 := []int{}
	output4 := xslice.Difference(s)
	sort.Ints(output4)
	if !reflect.DeepEqual(output4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, output4)
	}
}

func TestFilter(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	evenFilter := func(a int) bool {
		return a%2 == 0
	}
	evenSlice := xslice.Filter(intSlice, evenFilter)
	expectedEvenSlice := []int{2, 4}

	if !reflect.DeepEqual(evenSlice, expectedEvenSlice) {
		t.Errorf("Filter failed, expected %v but got %v", expectedEvenSlice, evenSlice)
	}

	stringSlice := []string{"Hello", "World", "Go"}
	lengthFilter := func(a string) bool {
		return len(a) > 3
	}
	longStringSlice := xslice.Filter(stringSlice, lengthFilter)
	expectedLongStringSlice := []string{"Hello", "World"}

	if !reflect.DeepEqual(longStringSlice, expectedLongStringSlice) {
		t.Errorf("Filter failed, expected %v but got %v", expectedLongStringSlice, longStringSlice)
	}
}

func TestFind(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	// Test finding an even number
	isEven := func(a int) bool { return a%2 == 0 }
	v, has := xslice.Find(intSlice, isEven)
	if !has || v != 2 {
		t.Errorf("Find failed, expected 2 but got %d", v)
	}

	// Test finding an odd number
	isOdd := func(a int) bool { return a%2 == 1 }
	v, has = xslice.Find(intSlice, isOdd)
	if !has || v != 1 {
		t.Errorf("Find failed, expected 1 but got %d", v)
	}

	// Test finding a number not in the slice
	isNegative := func(a int) bool { return a < 0 }
	v, has = xslice.Find(intSlice, isNegative)
	if has {
		t.Errorf("Find failed, expected false but got true")
	}
}

func TestEvery(t *testing.T) {
	// Test all elements are even
	isEven := func(a int) bool { return a%2 == 0 }
	allEven := xslice.Every([]int{2, 4, 6, 8, 10}, isEven)
	if !allEven {
		t.Errorf("Every failed, expected true but got false")
	}

	// Test some elements are odd
	allEven = xslice.Every([]int{2, 4, 6, 7, 10}, isEven)
	if allEven {
		t.Errorf("Every failed, expected false but got true")
	}

	// Test an empty slice
	allEven = xslice.Every([]int{}, isEven)
	if !allEven {
		t.Errorf("Every failed, expected true but got false")
	}
}

func TestShuffle(t *testing.T) {
	// 创建一个包含 10 个元素的整数切片。
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 100; i++ {
		// 记录初始顺序。
		original := xslice.Copy(slice)

		// 打乱切片顺序。
		xslice.Shuffle(slice)

		// 检查切片是否被打乱。
		if reflect.DeepEqual(slice, original) {
			t.Errorf("Slice not shuffled: %v", slice)
		}

		// 检查切片中的元素是否保持不变。
		for _, elem := range original {
			if !xslice.Contains(slice, elem) {
				t.Errorf("Element missing from shuffled slice: %d", elem)
			}
		}
	}
}

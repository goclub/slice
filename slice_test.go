package xslice_test

import (
	xslice "github.com/goclub/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexOfFromIndex(t *testing.T) {
	source := []string{"a", "b", "c"}
	{
		found, index := xslice.IndexOf(source, "a")
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(0))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "a", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(0))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "a", 1)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "a", 2)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "b", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(1))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "b", 1)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(1))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "b", 2)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "c", 0)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(2))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "c", 1)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(2))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "c", 2)
		assert.Equal(t, found, true)
		assert.Equal(t, index, uint64(2))
	}
	{
		found, index := xslice.IndexOfFormIndex(source, "c", 3)
		assert.Equal(t, found, false)
		assert.Equal(t, index, uint64(0))
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

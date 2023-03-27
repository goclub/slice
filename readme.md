# goclub/slice

[![Go Reference](https://pkg.go.dev/badge/github.com/goclub/slice.svg)](https://pkg.go.dev/github.com/goclub/slice)

## 安装

```bash
go get github.com/goclub/slice
```

xslice是一个Go语言中的切片工具包，提供了一些常用的切片操作函数，如查找、去重、合并等。

xslice中包含以下函数：

```go
// IndexOfFormIndex 返回从给定索引（包括该索引）开始在切片中搜索给定值的第一个匹配项的索引和是否找到
func IndexOfFormIndex(slice []T, value T, fromIndex uint64) (index uint64, found bool)

// IndexOf 在整个切片中搜索给定值的第一个匹配项的索引和是否找到
func IndexOf(slice []T, value T) (index uint64, found bool)

// ContainsFormIndex 返回从给定索引（包括该索引）开始在切片中搜索给定值是否存在
func ContainsFormIndex(slice []T, value T, fromIndex uint64) (found bool)

// Contains 返回整个切片中是否存在给定值
func Contains(slice []T, value T) (found bool)

// Union 返回给定切片的并集，即所有不同值的无序组合
func Union(slices ...[]T) (unorderedResult []T)

// Intersection 返回给定切片的交集，即所有出现在所有切片中的不同值的无序集合
func Intersection(slices ...[]T) (unorderedResult []T)

// Difference 返回第一个切片中不在任何其他切片中出现的所有不同值的有序集合
func Difference(slices ...[]T) (result []T)

// Unique 返回给定切片中所有不同的值的有序集合
func Unique(slice []T) (unique []T)

// First 返回切片中的第一个元素和是否存在
func First(slice []T) (v T, has bool)

// Last 返回切片中的最后一个元素和是否存在
func Last(slice []T) (v T, has bool)

// Filter 过滤切片中的元素，返回满足条件的元素组成的新切片
func Filter[T any](slice []T, fn func(a T) (save bool)) (newSlice []T)
```

这些函数都支持泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。如果 T 类型不满足 comparable 接口，则编译时会产生错误。



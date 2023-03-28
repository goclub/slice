# goclub/slice

[![Go Reference](https://pkg.go.dev/badge/github.com/goclub/slice.svg)](https://pkg.go.dev/github.com/goclub/slice)

## 安装

```bash
go get github.com/goclub/slice
```
goclub/slice 提供了一些常用的切片操作函数，如查找、去重、合并等。

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

// Find 返回满足条件的第一个元素和是否存在
func Find[T any](slice []T, fn func(a T) (found bool)) (v T, has bool)

// Every 函数接受一个类型为 T 的切片和一个接受类型为 T 的元素并返回布尔值的函数 fn
func Every[T any](slice []T, fn func(a T) (ok bool)) bool

// Some 用于检查给定的切片中是否有至少一个元素满足指定条件。
func Some[T any](slice []T, fn func(a T) (ok bool)) bool

// Shuffle 使用 Fisher-Yates shuffle 算法打乱输入的切片 slice 中元素的顺序。
// 注意，本函数会修改输入的切片 slice 中元素的顺序，而不会返回一个新的切片。
func Shuffle[T any](slice []T)

// Pluck 是一个通用函数，用于从切片中提取特定属性并生成一个新的切片。例如，从一个结构体切片中提取 ID 属性生成一个新的 ID 切片。
func Pluck[T any, Attr any](slice []T, fn func(v T) Attr) []Attr 

// Copy 函数接收一个泛型类型的切片 slice，并返回一个新的切片 newSlice，
func Copy[T any](slice []T) []T

// Pluck 是一个通用函数，用于从切片中提取特定属性并生成一个新的切片。
func Pluck[T any, Attr any](slice []T, fn func(v T) Attr) []Attr 
```

这些函数都支持泛型类型 T，并对其进行了类型约束，要求 T 实现了 comparable 接口，因为在 Go 中，泛型类型的比较需要使用比较运算符，因此需要确保 T 类型实现了该接口。如果 T 类型不满足 comparable 接口，则编译时会产生错误。



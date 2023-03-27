# Sort

官方对 Sort 包的定义如下：

> Package sort provides primitives for sorting slices and user-defined collections.
>

我们可以通过 `sort` 包实现以下功能：

- 对内置类型进行排序
- 自定义比较器/对自定义数据结构进行排序
- 检查是否排序
- 查找数据
- ...

## API

可以看出，`sort` 包并不仅限于对某种特定类型进行排序，而是为开发人员提供了更多的功能。因此，我们最好先了解一下 `sort` 包提供了哪些 API。这对于我们使用 `sort` 包和深入学习 `sort` 包源码都很有帮助。

### 1. 最常用的 API - 对内置类型进行排序

调用 sort 包，一般都是为了排序切片，因此它的排序功能是最常被使用的。包提供了 3 种内置类型的排序方法，分别是 [func Ints(x []int)](https://pkg.go.dev/sort#Ints)，[func Float64s(x []float64)](https://pkg.go.dev/sort#Float64s) 和 [func Strings(x []string)](https://pkg.go.dev/sort#Strings)

我们先来一个 demo 看看这 3 种 sort 方法都做了什么事情

```go
func sortByTypeDemo() {
	s1 := []int{1, 5, 9, 8, 2, 0, 1}
	sort.Ints(s1)
	fmt.Printf("sort int slice: %v\n", s1)
	
	s2 := []float64{3.14, 2.99, 9.87654321, 0.45, 9.99}
	sort.Float64s(s2)
    fmt.Printf("sort float slice: %v\n", s2)

    s3 := []string{"nut", "cat", "hello", "world", "cute"}
    sort.Strings(s3)
    fmt.Printf("sort string slice: %v\n", s3)
}
```

利用单测看看结果

```
=== RUN   Test_sortByTypeDemo
sort int slice: [0 1 1 2 5 8 9]
sort float slice: [0.45 2.99 3.14 9.87654321 9.99]
sort string slice: [cat cute hello nut world]
--- PASS: Test_sortByTypeDemo (0.00s)
PASS
ok      go_all_clear/chapter07_standard_library/code
```

- 总结：
    - 3 种 type 排序都是按照升序排序的
    - string 的排序按首字母排序

`sort.Float64s`、`sort.Ints` 和 `sort.Strings` 都是 Go 语言 `sort` 包中提供的函数，用于**对相应类型的切片进行升序排序**。

> 简述 sort 的排序
>

这些函数都可以对切片进行**原地排序**，也就是说，排序后的结果将覆盖原来的切片。排序函数并不会返回任何结果，所以你使用 `new_slice := sort.Ints(int_slice)`，那么会报错 `sort.Ints(s2) (no value) used as value`。如果需要保留原来的切片，可以先复制一份再进行排序。

> 算法复杂度
>

通常，我们认为 sort 包排序的时间复杂度为 `O(nlog(n))`，n 是 slice 的长度。因为 `sort.Float64s`、`sort.Ints` 和 `sort.Strings` 内部都使用了一种经典的快速排序算法。所以更具体的说，sort 包排序的平均时间复杂度为 `O(nlog(n))`，最坏时间复杂度为 `O(n^2)`

实际上，sort 包内部提供了多种排序算法，包括冒泡排序、插入排序、快速排序、堆排序和归并排序等。不同的排序算法具有不同的时间复杂度，因此在选择排序算法时需要综合考虑数据规模和数据特征。

### 2. 判断是否排序

sort 包提供了以下几个函数用于判断切片是否已递增排序，这些函数都接受一个切片作为参数，并返回一个布尔值：

- `func IntsAreSorted(a []int) bool`
- `func Float64sAreSorted(a []float64) bool`
- `func StringsAreSorted(a []string) bool`

通过以下代码演示这些函数的使用：

```go
func isSortedDemo() {
    s1 := []int{1, 2, 3, 4, 5}
    fmt.Printf("Is s1 sorted? %v\\n", sort.IntsAreSorted(s1))

    s2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
    fmt.Printf("Is s2 sorted? %v\\n", sort.Float64sAreSorted(s2))

    s3 := []string{"apple", "banana", "cherry", "date", "eggplant"}
    fmt.Printf("Is s3 sorted? %v\\n", sort.StringsAreSorted(s3))
}

```

输出结果：

```
Is s1 sorted? true
Is s2 sorted? true
Is s3 sorted? false

```

- 总结：
    - `IntsAreSorted` 和 `Float64sAreSorted` 只能判断是否升序排列，不能判断是否降序排列。
    - `StringsAreSorted` 只能判断按字典序（首字母顺序）排列的字符串切片。
    - 如果切片是已排序的，则返回 true，否则返回 false。

### 3. 查找元素

sort 包中还提供了一些函数用于在已排序的切片中查找元素。这些函数都接受一个已排序的切片和要查找的元素作为参数，并返回一个整数值：

- `func Search(n int, f func(int) bool) int`
- `func SearchInts(a []int, x int) int`
- `func SearchFloat64s(a []float64, x float64) int`
- `func SearchStrings(a []string, x string) int`

其中，`Search` 函数接受一个函数作为参数，该函数用于比较元素。这个函数接受一个整数作为参数，表示要比较的元素的索引，返回一个布尔值，表示该元素是否等于要查找的元素。**如果在切片中找到了要查找的元素，则返回该元素的索引；否则，返回应该插入它的位置，以保证切片仍然是有序的。**

以下是一个使用 `Search` 函数查找元素的示例代码：

```go
func searchDemo(s []int, x int) string {
    i := sort.Search(len(s), func(i int) bool { return s[i] >= x })
    if i < len(s) && s[i] == x {
            return fmt.Sprintf("%d is found at index %d", x, i)
    } else {
            return fmt.Sprintf("%d is not found, it should be inserted at index %d", x, i)
    }
}
```

单元测试

```go
func Test_searchDemo(t *testing.T) {
    sliceCase := []int{1, 2, 3, 4, 6, 7}
    res1 := searchDemo(sliceCase, 6)
    t.Log(res1)
    res2 := searchDemo(sliceCase, 5)
    t.Log(res2)
}
```

输出结果：

```
=== RUN   Test_searchDemo
    code_test.go:16: 6 is found at index 4
    code_test.go:18: 5 is not found, it should be inserted at index 4
--- PASS: Test_searchDemo (0.00s)
PASS

```

如果要查找的元素不在切片中，则可以通过返回值计算出它应该插入的位置。例如，上面的示例中，要查找的元素是 6，它的索引是 4。如果要插入一个不存在与切片中的元素 5 到切片中，则应该将它插入到索引 4 的位置。

另外，sort 包还提供了 `SearchInts`、`SearchFloat64s` 和 `SearchStrings` 函数，这些函数是 `Search` 函数的特化版本，用于在整型、浮点型和字符串切片中查找元素。这些函数的用法与 `Search` 函数类似。

### 4. 其他排序方式

我们上面遇到的排序方式都是比较常用的排序方式，而且，都是采用了递增排序，现在，我们来看看这两个 API

> 逆向排序
>

`sort.Reverse` 函数是 Go `sort` 包中提供的一个函数，用于**逆向排序**，即将升序排序转化为降序排序。

`sort.Reverse` 函数接受一个实现了 `sort.Interface` 接口的对象，并返回一个新的对象，该对象是原对象的逆序排列。例如，如果我们有一个 `[]int` 类型的切片，想要将其从大到小排序，可以使用以下代码：

```
s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
sort.Sort(sort.Reverse(sort.IntSlice(s)))
fmt.Println(s) // Output: [9 6 5 5 5 4 3 3 2 1 1]

```

这里的 `sort.IntSlice(s)` 是一个 `[]int` 类型的切片，它实现了 `sort.Interface` 接口。`sort.Reverse` 函数将其转换为一个新的对象，该对象是原对象的逆序排列。`sort.Sort` 函数按照逆序排列对其进行排序，并将结果保存回原对象。

`sort.Reverse` 函数还可以用于自定义类型的排序。例如，如果我们有一个自定义类型 `Person`，其中包含一个 `Age` 属性，我们可以按照 `Age` 逆序排列 `Person` 类型的切片：

```
type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
    people := []Person{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 20},
        {"Dave", 35},
    }

    sort.Sort(sort.Reverse(ByAge(people)))

    fmt.Println(people) // Output: [{Dave 35} {Bob 30} {Alice 25} {Charlie 20}]
}

```

这里的 `ByAge` 类型实现了 `sort.Interface` 接口，用于按照 `Age` 属性排序。`sort.Reverse` 函数将其转换为一个新的对象，该对象是原对象的逆序排列。`sort.Sort` 函数按照逆序排列对其进行排序，并将结果保存回原对象。
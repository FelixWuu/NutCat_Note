# Sort

官方对 Sort 包的定义如下：

> Package sort provides primitives for sorting slices and user-defined collections.

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

```go
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

这些函数都可以对切片进行**原地排序**，也就是说，排序后的结果将覆盖原来的切片。排序函数并不会返回任何结果，所以你使用 `new_slice := sort.Ints(int_slice)`，那么会报错 `sort.Ints(s2) (no value) used as value`。如果需要保留原来的切片，可以先复制一份再进行排序。

> 算法复杂度

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

### 4. 自定义排序

直接上例子

```go
func sliceDemo() string {
	people := []Person{
		{"NutCat", 18, 100},
		{"Wuu", 27, 100},
		{"Cat", 10, 100},
		{"BAO", 27, 90},
		{"Felix", 26, 99},
		{"AAA", 26, 83},
		{"Felix", 29, 87},
	}

	// sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})

	s1 := fmt.Sprintf("sort by name: %v\n", people)

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	s2 := fmt.Sprintf("sort by age: %v\n", people)

	// sort by score
	sort.Slice(people, func(i, j int) bool {
		return people[i].Score > people[j].Score
	})

	s3 := fmt.Sprintf("sort by score: %v\n", people)

	return strings.Join([]string{s1, s2, s3}, "\n")
}
```

单测结果;

```
=== RUN   Test_sliceDemo
    code_test.go:29: sort by name: [{Wuu 27 100} {NutCat 18 100} {Felix 26 99} {Felix 29 87} {Cat 10 100} {BAO 27 90} {AAA 26 83}]
        
        sort by age: [{Cat 10 100} {NutCat 18 100} {AAA 26 83} {Felix 26 99} {BAO 27 90} {Wuu 27 100} {Felix 29 87}]
        
        sort by score: [{Cat 10 100} {NutCat 18 100} {Wuu 27 100} {Felix 26 99} {BAO 27 90} {Felix 29 87} {AAA 26 83}]
        
--- PASS: Test_sliceDemo (0.00s)
PASS
```

### 5. 其他排序方式

我们上面遇到的排序方式都是比较常用的排序方式，而且，都是采用了递增排序，现在，我们来看看这两个 API

> 逆向排序

`sort.Reverse` 函数是 Go `sort` 包中提供的一个函数，用于**逆向排序**，即将升序排序转化为降序排序。

`sort.Reverse` 函数接受一个实现了 `sort.Interface` 接口的对象，并返回一个新的对象，该对象是原对象的逆序排列。例如，如果我们有一个 `[]int` 类型的切片，想要将其从大到小排序，可以使用以下代码：

```go
func reverseIntDemo(s []int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	return fmt.Sprintf("%#v", s)
}
```

单元测试：

```go
func Test_reverseIntDemo(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	res := reverseIntDemo(s)
	t.Log(res)
}
```

输出结果：

```go
=== RUN   Test_reverseDemo
    code_test.go:24: []int{9, 6, 5, 5, 5, 4, 3, 3, 2, 1, 1}
--- PASS: Test_reverseDemo (0.00s)
PASS
```

> 稳定排序

`sort` 包还提供了一个名为 `stable` 的函数，可以对切片进行**稳定排序**。稳定排序指的是，如果两个元素的值相等，它们在排序后的顺序不会改变。

我们可以使用**SliceStable()**函数对切片进行稳定排序。我们还可以给定指定的 less 函数，对切片进行自定义的排序，同时保持相等元素的原始顺序。

语法：`func SliceStable(a_slice interface{}, less func(p, q int) bool)`

例如，一个包含多个 Person 结构体的切片，每个人的年龄和姓名都是一个字段，我们想要按照年龄排序：

```go
type Person struct {
	Name string
	Age  int
}

func stableDemo() string {
	people := []Person{
		{"NutCat", 18, 100},
		{"Wuu", 27, 100},
		{"Cat", 10, 100},
		{"BAO", 27, 90},
		{"Felix", 26, 99},
		{"AAA", 26, 83},
		{"Felix", 29, 87},
	}

	// sort by name
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})

	s1 := fmt.Sprintf("sort by name: %v\n", people)

	// sort by age
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	s2 := fmt.Sprintf("sort by age: %v\n", people)

	// sort by score
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Score > people[j].Score
	})

	s3 := fmt.Sprintf("sort by score: %v\n", people)

	return strings.Join([]string{s1, s2, s3}, "\n")
}

```

单元测试：

```go
func Test_stableDemo(t *testing.T) {
	res := stableDemo()
	t.Log(res)
}

```

输出结果：

```
=== RUN   Test_stableDemo
    code_test.go:29: sort by name: [{Wuu 27 100} {NutCat 18 100} {Felix 26 99} {Felix 29 87} {Cat 10 100} {BAO 27 90} {AAA 26 83}]
        
        sort by age: [{Cat 10 100} {NutCat 18 100} {Felix 26 99} {AAA 26 83} {Wuu 27 100} {BAO 27 90} {Felix 29 87}]
        
        sort by score: [{Cat 10 100} {NutCat 18 100} {Wuu 27 100} {Felix 26 99} {BAO 27 90} {Felix 29 87} {AAA 26 83}]
        
--- PASS: Test_stableDemo (0.00s)
PASS

```

- 名字：从大到小排序，名字相同的人排序的先后顺序与原来一样
- 年龄：从小到大排序，年龄相同的人排序的先后顺序与原来一样, 与第4节相比，这里的 Wuu 和 BAO 的顺序没有变化，仍然是 Wuu 在前，BAO 在后。
- 分数：从大到小排序，分数相同的人排序的先后顺序与原来一样

### 6. 各种类型的排序

官方文档中，有一些参数类型是 data Interface，如：

- [func IsSorted(data Interface) bool](https://pkg.go.dev/sort#IsSorted)
  - IsSorted函数用于检查一个序列是否已经按照升序排列
- [func Sort(data Interface)](https://pkg.go.dev/sort#Sort)
  - Sort函数用于对一个序列进行升序排序
- [func Stable(data Interface)](https://pkg.go.dev/sort#Stable)
  - Stable函数用于对一个序列进行稳定排序

我们直接看看例子

```go
func dataInterfaceSortDemo() string {
	ints := []int{2, 3, 5, 7, 11, 13}
	strs := []string{"nut", "cat", "hello", "felix", "wuu"}

	// is sort
	res1 := fmt.Sprintf(
		"is ints sorted: %v, is strs sorted: %v\n",
		sort.IsSorted(sort.IntSlice(ints)),
		sort.IsSorted(sort.StringSlice(strs)),
	)

	// sort
	sort.Sort(sort.StringSlice(strs))
	res2 := fmt.Sprintf("Sorted string slice :%v\n", strs)

	// stable
	float64s := []float64{3.14, 2.99, 8.76, 0.98, 2.99, 1.00, 5.55}
	sort.Stable(sort.Float64Slice(float64s))
	res3 := fmt.Sprintf("Sorted float64 slice :%v\n", float64s)

	return strings.Join([]string{res1, res2, res3}, "\n")
}
```

它将输出：

```go
=== RUN   Test_dataInterfaceSortDemo
    code_test.go:39: is ints sorted: true, is strs sorted: false
        
        Sorted string slice :[cat felix hello nut wuu]
        
        Sorted float64 slice :[0.98 1 2.99 2.99 3.14 5.55 8.76]
        
--- PASS: Test_dataInterfaceSortDemo (0.00s)
PASS
```

我们看到参数类型 —— data Interface，似乎不是一个 slice，那么这些类型是什么呢？又有着什么特性，这就是我们第7节的内容

### 7. 聊聊 sort 包中的各种 type

我们前文看到的 data Interface 参数实际上是待排序的数据集合。sort 包内部已支持 **IntSlice 类型、Float64Slice 类型和 StringSlice 类型。**

我们在官方文档可以看到这三种 type，并且它们实现了 `sort.Interface` 的接口。所以，我们先来看看最基本的 `sort.Interface` 接口做了什么。

`sort.Interface` 是 Go 语言中 sort 包的核心接口，该接口定义了排序所需的全部方法。如果要使用 sort 包中的任何排序算法，都需要实现 `sort.Interface` 接口。

sort.Interface 接口有三个方法：

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

sort.Interface 接口的方法解释如下：

- `Len()` 方法返回待排序的元素数量。
- `Less(i, j int) bool` 方法用于比较索引为 i 和 j 的元素。如果元素 i 应该排在元素 j 之前，则返回 true；否则，返回 false。
- `Swap(i, j int)` 方法用于交换索引为 i 和 j 的元素的位置。

sort.Interface 接口的实现可以是任何类型的序列，包括数组、切片、映射甚至是通道。在实现 sort.Interface 接口时，需要根据实际情况更新三个方法，以便它们适用于你的数据类型。

例如，我们来对一个结构体切片进行排序，它应该包含 `Len()`、`Less(i, j int) bool` 和 `Swap(i, j int)` 方法。实现该接口后，我们还能够使用 sort 包中的 `Sort()`、`Stable()` 或 `Search()` 等函数对其进行排序。

下面是一个实现 sort.Interface 接口的例子：

```go
type Author struct {
	Person
	Books []string
}

type AuthorSlice []Author

func (a AuthorSlice) Len() int {
	return len(a)
}

func (a AuthorSlice) Less(i, j int) bool {
	// return a[i].Age < a[j].Age  // sort by age
	return len(a[i].Books) > len(a[j].Books)
}

func (a AuthorSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AuthorSlice) SortAuthors() {
	sort.Sort(AuthorSlice(a))
}
```

测试：

```go
func Test_Authors(t *testing.T) {
	authors := []Author{
		{
			Person: Person{Name: "NutCat", Age: 26, Score: 99},
			Books:  []string{"母猪的产后护理", "公鸡下蛋教程", "配角如何抢走主角的戏", "宫廷玉液酒"},
		},
		{
			Person: Person{Name: "鲁迅", Age: 10000, Score: 100},
			Books:  []string{"孔乙己", "骆驼祥子"},
		},
	}

	a := AuthorSlice(authors)
	a.SortAuthors()
	t.Logf("the author slice is sorted?: %v", sort.IsSorted(a))
	t.Logf("the sorted authors slice: %v", authors)
}
```

结果：

```
=== RUN   Test_Authors
    code_test.go:59: the author slice is sorted?: true
    code_test.go:60: the sorted authors slice: [{{NutCat 26 99} [母猪的产后护理 公鸡下蛋教程 配角如何抢走主角的戏 宫廷玉液酒]} {{鲁迅 10000 100} [孔乙己 骆驼祥子]}]
--- PASS: Test_Authors (0.00s)
PASS
```

我们在 `sort.Sort()` 函数中传递了 `AuthorSlice` 类型的切片，这是因为 `AuthorSlice` 类型实现了 `sort.Interface` 接口。由于实现了该接口，因此 `AuthorSlice` 类型可以传递到任何 `sort` 包中的函数中进行排序。

### 8. 总结

- 最常用的 API 就是第一节的 API，它基本符合我们平时的排序需求
- 当我们要查找元素时，使用 Search（第二节内容）
  - 它内部是使用二分查找的，算法复杂度为 **O(logN)**
  - 当我们需要使用二分查找算法来寻找元素或寻找元素插入位置时，直接使用 search api 即可（需要先排好序）
- 我们也可以自定义排序，包括自定义排序顺序，或者对不同的数据类型排序
  - 特殊情况下，我们需要稳定的排序


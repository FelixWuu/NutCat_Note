package code_7

import (
	"fmt"
	"sort"
	"strings"
)

func sortByTypeDemo() {
	s1 := []int{1, 5, 9, 8, 2, 0, 1}
	sort.Ints(s1)
	fmt.Printf("sort int slice: %v\n", s1)

	s2 := []float64{3.14, 2.99, 9.87654321, 0.45, 9.99}
	fmt.Printf("sort float slice: %v\n", s2)

	s3 := []string{"nut", "cat", "hello", "world", "cute"}
	sort.Strings(s3)
	fmt.Printf("sort string slice: %v\n", s3)
}

func isSortedDemo() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Is s1 sorted? %v\\n", sort.IntsAreSorted(s1))

	s2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Printf("Is s2 sorted? %v\\n", sort.Float64sAreSorted(s2))

	s3 := []string{"apple", "banana", "cherry", "date", "eggplant"}
	fmt.Printf("Is s3 sorted? %v\\n", sort.StringsAreSorted(s3))
}

func searchDemo(s []int, x int) string {
	i := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if i < len(s) && s[i] == x {
		return fmt.Sprintf("%d is found at index %d", x, i)
	} else {
		return fmt.Sprintf("%d is not found, it should be inserted at index %d", x, i)
	}
}

func reverseIntDemo(s []int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	return fmt.Sprintf("%#v", s)
}

type Person struct {
	Name  string
	Age   int
	Score int
}

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

package code_7

import (
	"fmt"
	"sort"
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

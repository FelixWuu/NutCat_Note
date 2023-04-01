package code_7

import (
	"sort"
	"testing"
)

func Test_sortByTypeDemo(t *testing.T) {
	sortByTypeDemo()
}

func Test_isSortedDemo(t *testing.T) {
	isSortedDemo()
}

func Test_searchDemo(t *testing.T) {
	sliceCase := []int{1, 2, 3, 4, 6, 7}
	res1 := searchDemo(sliceCase, 6)
	t.Log(res1)
	res2 := searchDemo(sliceCase, 5)
	t.Log(res2)
}

func Test_reverseIntDemo(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	res := reverseIntDemo(s)
	t.Log(res)
}

func Test_sliceDemo(t *testing.T) {
	res := sliceDemo()
	t.Log(res)
}

func Test_stableDemo(t *testing.T) {
	res := stableDemo()
	t.Log(res)
}

func Test_dataInterfaceSortDemo(t *testing.T) {
	res := dataInterfaceSortDemo()
	t.Log(res)
}

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

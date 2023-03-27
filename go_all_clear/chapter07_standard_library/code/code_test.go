package code_7

import "testing"

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

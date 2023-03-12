package main

import "fmt"

func main() {
	//fmt.Println("testing NilChannelsDemo")
	//NilChannelDemo()
	//fmt.Println("\nfinished")

	//fmt.Println("testing  GetStringComparisonSpeed")
	//GetStringComparisonSpeed()
	//fmt.Println("\nfinished")

	fmt.Println("testing  Struct Compare")
	c1 := &Company{
		Name: "搜索 NutCat",
		Employees: []*Person{
			{Name: "NutCat", Age: 18},
			{Name: "猫坚果", Age: 19},
		},
	}
	c2 := &Company{
		Name: "搜索 NutCat",
		Employees: []*Person{
			{Name: "NutCat", Age: 18},
			{Name: "猫坚果", Age: 19},
		},
	}

	if c1.Equals(c2) {
		fmt.Println("c1 and c2 are equal")
	} else {
		fmt.Println("c1 and c2 are not equal")
	}

	fmt.Println("\nfinished")
}

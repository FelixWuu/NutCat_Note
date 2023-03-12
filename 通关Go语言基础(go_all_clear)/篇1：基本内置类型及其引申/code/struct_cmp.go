package main

type Person struct {
	Name string
	Age  int
}

type Company struct {
	Name      string
	Employees []*Person // 包含不可比较类型的字段
}

func (c *Company) Equals(other *Company) bool {
	if c.Name != other.Name {
		return false
	}
	if len(c.Employees) != len(other.Employees) {
		return false
	}
	for i := range c.Employees {
		if c.Employees[i].Name != other.Employees[i].Name ||
			c.Employees[i].Age != other.Employees[i].Age {
			return false
		}
	}
	return true
}

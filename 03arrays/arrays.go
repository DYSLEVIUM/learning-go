package main

import "fmt"

func main() {
	grades := [3]int{10, 120, 51}
	grades2 := [...]int{10, 120, 51}
	grades3 := []int{10, 120, 51}

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)
	fmt.Printf("Grades: %v\n", grades3)

	var students [3]string
	fmt.Printf("Students: %v\n", students)

	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"

	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[0])
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	fmt.Printf("Identity Matix: %v\n", identityMatrix)

	// makes a copy
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// using pointer
	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

	// slice; references as default for []
	s := []int{1, 2, 3}
	s2 := s
	s2[0] = 0

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Printf("Length %v\n", len(s))
	fmt.Printf("Capacity %v\n", cap(s))

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := x[:]
	g := x[3:]
	h := x[:6]
	i := x[1:2]
	x[5] = 42

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	// len and capacity similar to vector in c++

	m := make([]int, 1, 2) // similar to reserve in c++ std::vector
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	m = append(m, 1)
	m = append(m, 1)
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	m = append(m, []int{15, 151, 51, 57, 12, 1}...) // javascript spread
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	middle := append(m[:2], m[3:]...)
	fmt.Println(middle)
}

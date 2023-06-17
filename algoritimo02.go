package main

import "fmt"

func unirMapas(m1 map[string]int, m2 map[string]int) map[string]int {

	m3 := make(map[string]int)

	for range1, range2 := range m1 {

		m3[range1] = range2

	}

	for range3, range4 := range m2 {

		m3[range3] = range4

	}

	return m3

}

func main() {

	m1 := map[string]int{

		"pa":       1,
		"f": 4,
		"putz":       2,
	}

	m2 := map[string]int{

		"gh": 2,
		"pa":              999,
		"3":                 2,
	}

	fmt.Println(unirMapas(m1, m2))

}

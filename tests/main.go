package main

import "fmt"

func main() {

	s1 := make([]int, 1, 10)
	fmt.Printf("%p", &s1)
	change(s1)
	fmt.Println(s1)

}

func change(s2 []int) {

	fmt.Println()

	s2 = append(s2, 1)
	fmt.Printf("change 中的 %p \n", &s2)

	s2[0] = 2
	fmt.Println(s2)
	fmt.Println("change结束")
}

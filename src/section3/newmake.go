package main

import "fmt"

func generater() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// makeは空を返す
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	// new はポインタを返す
	var p *int = new(int)
	fmt.Printf("%T\n", p)

	// ints := generater()

	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())

	// ints2 := generater()

	// fmt.Println(ints2())
	// fmt.Println(ints2())
	// fmt.Println(ints2())
	// fmt.Println(ints2())
	// fmt.Println(ints2())
	// fmt.Println(ints2())

	if b := 100; b == 100 {
		fmt.Println("if")
	}

	var a int = 0
	if a := 100; a == 100 {
		fmt.Println(a)
	}
	fmt.Println(a, "aaa")

}

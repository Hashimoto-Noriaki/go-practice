package main

import "fmt"

//算術演算子

func main() {
	fmt.Println(1 + 2)
	fmt.Println("Water" + "Melon")

	fmt.Println(10 - 3)

	fmt.Println(5 * 8)

	fmt.Println(50 / 5)

	fmt.Println(13 % 3)

	n:= 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n --
	fmt.Println(n)

	s := "Water"
	s += "Melon"
	fmt.Println(s) //s= "Water" + "Melon"
}
/*
3
WaterMelon
7
40
10
1
4
5
4
WaterMelon
*/

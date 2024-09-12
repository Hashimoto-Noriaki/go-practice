package main

import "fmt"

func Double(i int){
	i = i * 2
}

func Doublev2(i *int){
	*i = *i*2
}

func Doublev3(s []int){
	for i,v := range s{
		s[i] = v * 2
	}
}

//ポインタ
func main(){
	var n int = 1000
	fmt.Println(n)

	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	//ポインタ宣言
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	/*
	*p = 3000
	fmt.Println(n)
	n = 2000
	fmt.Println(*p)
	*/
	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	var sl []int = []int{1,2,3}
	Doublev3(sl)
	fmt.Println(sl)
}

/*
1000
0xc000090020
1000
0xc000090020
1000
2000
4000
[2 4 6]
*/

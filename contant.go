package main
import "fmt"
//定数(グローバルに書くことが多い)
const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)
//暗号の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

const Big int = 34567898 + 1

func main(){
	fmt.Println(Pi)
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL,SiteName)

	fmt.Println(A,B,C,D,E,F)

	fmt.Println(Big - 1)

	fmt.Println(c0,c1,c2)
}

/*
3.14
http://xxx.co.jp test
1 1 1 A A A
34567898
0 1 2
*/

package main
import "fmt"
//変数

//i5 := 500
var i5 int = 500

func main(){
	//明示的な定義
	// var 変数名　型　= 値
	var i int = 1000
	fmt.Println(i)
	// 1000

	var s string = "Goの練習だ"
	fmt.Println(s)
	//Goの練習だ

	var t,f bool = true,false
	fmt.Println(f,t)
	//true false

	var(
		i2 int = 2000
		s2 string = "Goをやるぞ"
	)
	fmt.Println(i2,s2)
	//2000 Goをやるぞ

	var i3 int 
	var s3 string 
	fmt.Println(i3,s3)
	//0

	i3 = 300
	s3 = "Go"
	fmt.Println(i3,s3)
	//300 Go

	i = 1500
	fmt.Println(i)
	/*
	1500
	値の更新をした
	*/

	/*
	暗黙的な定義
	変数名　:= 値 
	*/
	i4 := 4000
	fmt.Println(i4)
	//4000

	i4 = 4500
	fmt.Println(i4)
	//4500

	// i4 := 5000
	// fmt.Println(i4)
	// i4 = "いえい"
	/* fmt.Println(i4)
	# command-line-arguments
	./valuable.go:59:7: cannot use "いえい" (untyped string constant) as int value in assignment */

	fmt.Println(i5)
	//500  上で定義されている
}


package main
//変数
import (
	"fmt"
)
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
}


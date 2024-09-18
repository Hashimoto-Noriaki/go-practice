package main
//関数
//無名関数
import "fmt"

func main(){
	f := func(x,y int) int {
		return x + y
	}
	i := f(10,20)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(50,30)
	fmt.Println(i2)
}
//30
//80

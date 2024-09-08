package main //Goはpackageの宣言から
//fmtはformatパッケージ
import (
	"fmt"
	"time"
)

//簡単な出力　　main関数がエントリーポイントでプログラムの本体
func main(){
	fmt.Println("Goの練習")
	fmt.Println(time.Now())
}
/*
Goの練習と出力
今の日付が出力
*/

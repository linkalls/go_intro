package main //* かならず何らかのパッケージを指定する
//* 宣言できるのは一つに限られる
// import "fmt"
// import "time"
import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())
	var count int64 = 100
	var number int
	var text string //* string型の初期値は空文字列
	fmt.Println(count)
	fmt.Println(text, number)
	some := 100 //* 関数外部では使用不可
	fmt.Println(some)
	fmt.Printf("%T\n", some) //* %Tっていうのはintを表す　これで型がわかる

	fmt.Println(int32(count)) //* これで文字列に変換できる
	//* 型変換は基本的に 型名(変数名) で行う

	fl := 2.5
	fmt.Printf("%T", fl)            //* これで型がわかる
	fmt.Println()                   //* 改行
	fmt.Println(fl + float64(some)) //* これで計算できるようにintをfloat64に変換

	fmt.Println(reflect.TypeOf(fl)) //* これも型がわかる

	var byteText byte = 72 //* byte型は0~255までの値しか取れない
	fmt.Println(text, string(byteText))

	//* こっから配列とスライス
	var array [3]int //* 配列は要素数を指定して宣言する
	array[0] = 1
	array[1] = 10
	array[2] = 100
	fmt.Println(array)

	var slice []string             //* スライスは要素数を指定しない
	slice = append(slice, "apple") //* appendで要素を追加 jsのスプレッド構文みたいな感じ

	nums := make([]int, 2, 5) // 長さ2, 容量5のスライス
	fmt.Println("長さ:", len(nums), "容量:", cap(nums))
	nums = append(nums, 10, 20, 30)
	fmt.Println("長さ:", len(nums), "容量:", cap(nums))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

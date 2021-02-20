package main

import "fmt"

func main(){
	array1 := [5]float32{}
	array2 := [6]int{1,2,3,4,5,6}
	array3 := [...]string{"今日","明日","明後日"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	array3[0]="月"
	array3[2]="火"
	fmt.Println(array3)
}

/*
配列の定義の方法3種類
1. 変数名:=[要素数]int{}
2. 変数名:=[要素数]int{1,2,3,4,...}
3. 変数名:=[...]int{1,2,3,4,...}
 */
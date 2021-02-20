package main

import "fmt"

func f1(){
	defer fmt.Println("deferで呼び出されます")
	fmt.Println("パニックを発生させます")
	panic("パニック発生しました")
	fmt.Println("3")
}


func main(){
	f1()
}

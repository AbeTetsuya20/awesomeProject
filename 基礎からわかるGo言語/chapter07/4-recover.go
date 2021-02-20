package main

import "fmt"

func main(){
	f1(false)
	fmt.Println("falseを代入しました。")
	f1(true)
	fmt.Println("trueを代入しました。")
}

func f1(p bool){
	defer func(){
		fmt.Println("遅延関数開始")

		if err := recover(); err != nil{
			fmt.Println("パニックを中断します:", err)
		}
		fmt.Println("遅延関数終了")
	}()

	if p{
		panic("パニック発生")
	}
}


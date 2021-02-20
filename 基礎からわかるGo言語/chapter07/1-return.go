package main

import(
	"fmt"
	"os"
)

func main(){
	file,err := os.Open("test.txt")

	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("ファイルが開けませんでした。")
		fmt.Println(err)
		os.Exit(1)
	}

	file.Close()
	fmt.Println("OK")
}
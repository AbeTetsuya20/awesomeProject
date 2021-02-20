package main

import "fmt"

func main(){
	var array1 [1]byte
	var array2 [2]*int
	var array3 [8][3]int64
	var array4 [2]struct{x,y int}

	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(array3),len(array3[0]))
	fmt.Println(len(array4))
}

//二次元配列→array[8][3]
//array→{{1,2,3},{4,5,6},{　…　}}
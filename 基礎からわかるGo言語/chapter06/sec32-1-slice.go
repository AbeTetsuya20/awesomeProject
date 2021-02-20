package main

import "fmt"

func main(){
	/*
	x:=[5]string{"a","b","c","d","e"}
	var s1 []string

	s1=x[:]
	fmt.Println(s1)

	s2:=x[1:4]
	fmt.Println(s2)

	s3:=x[3:]
	fmt.Println(s3)

	s4:=x[:4]
	fmt.Println(s4)
	 */

	values:=[...]int{0,1,2,3,4}

	fmt.Println(values[:])
}

func double(values []int){
	for i:=0;i<len(values);i++{
		values[i]*=2
	}
}

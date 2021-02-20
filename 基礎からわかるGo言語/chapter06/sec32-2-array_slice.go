package main

import "fmt"

func double(values []int){
	for i:=0;i<len(values);i++{
		values[i]*=2
	}
}

func double2(values []int){
	for value1,value2:=range values{
		values[value1]=value2*2
	}
}

func main(){
	values:=[...]int{0,1,2,3,4}
	double(values[:])
	fmt.Println(values)

	double(values[:])
	fmt.Println(values)
}
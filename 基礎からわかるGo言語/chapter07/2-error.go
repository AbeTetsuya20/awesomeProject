package main

import "fmt"

type MyError struct{
	message string
}

func (err MyError) Error() string{
	return err.message
}

func hex2int(hex string)(val int,err error){
	for _,r:=range hex{
		val *= 16
		switch{
		case '0' <= r && r <= '9':
			val += int(r-'0')
		default:
			return 0,MyError{"不正な文字です。:"+string(r)}
		}
	}
	return
}

func main(){
	val,err:=hex2int("1")
	fmt.Println(val,err)

	val,err=hex2int("z")
	fmt.Println(val,err)
}
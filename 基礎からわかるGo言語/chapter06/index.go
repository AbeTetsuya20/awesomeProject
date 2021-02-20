package main

import "fmt"

func main(){
	var date [7]string
	date[0]="日曜日"
	date[1]="月曜日"
	date[2]="火曜日"
	date[3]="水曜日"
	date[4]="木曜日"
	date[5]="金曜日"
	date[6]="土曜日"

	fmt.Print(date)

	for i:=0;i<7;i++{
		fmt.Print(i," ",date[i]," ")
	}

	fmt.Println()

	for value1,value2 := range date{
		fmt.Print(value1,value2," ")
	}
	/*
	配列の取り出し方は3種類
	1.fmt.Print(配列名)
	2,3. for文で取り出す
	for i:=0;i<要素数;i++{
	fmt.Print(配列名[i])
	}
	for 変数1,変数2:= range 変数名{
	fmt.Print(変数2)
	}
	ちなみに変数1はインデックス番号、変数2は要素
	 */


}
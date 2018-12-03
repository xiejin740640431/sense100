package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	var tHour float64 = 60 * 60 * 1000
	var tMin = 60 * 1000
	var num float64
	num = 1.5
	var timeCount = tHour*num

	fmt.Println(tMin,timeCount)
	//var hour = int(timeCount / tHour)
	//var min = timeCount % tHour / tMin
	//var sec = timeCount % tHour % tMin
	//
	//fmt.Println(hour, min, sec)
}

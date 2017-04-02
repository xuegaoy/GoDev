package main

import (
"fmt"
"time"
)

func main() {
	fmt.Print(time.Now())
	fmt.Println("222222222")
	//时间加一天显示
	fmt.Print(time.Now().AddDate(0,0,1))
	fmt.Println("222222222")
	//时间加2个小时显示
	fmt.Print(time.Now().Add(time.Hour).Add(time.Hour))
}

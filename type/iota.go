package main

import (
	"fmt"
	"strings"
)
type CouponStatus byte

const (
	NotAssigned1 CouponStatus = iota
	NotAssigned2 CouponStatus = iota
	NotAssigned3 CouponStatus = iota
	aa1          string       = "1"
)
type CoupleDetail struct {
	abc CouponStatus
}
func main() {
	var aaa CoupleDetail
	type Coup2 string
	a := "1"
	if (strings.EqualFold(a,"1")){
		//拿到的常量
		aaa.abc = NotAssigned3
	}
	fmt.Println(aaa.abc)
	fmt.Println(NotAssigned1)
	fmt.Println(NotAssigned2)


}

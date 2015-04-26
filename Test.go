/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.26]
*/

package main

import (
	"fmt"
	"github.com/tjCFeng/GoRK3288/RK3288"
)

func main() {
	P8A1 := RK3288.CreateGPIO(RK3288.P8, RK3288.A1)
	P8A1.Flip()
	fmt.Println(P8A1.GetLevel())
	RK3288.FreeGPIO(P8A1)
}

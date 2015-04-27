/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.26]
*/

package main

import (
	"fmt"
	"time"
	"github.com/tjCFeng/GoRK3288/RK3288"
)

func main() {
	P8A1 := RK3288.CreateGPIO(RK3288.P8, RK3288.A1)
	P8A1.Flip()
	fmt.Println(P8A1.GetLevel())
	RK3288.FreeGPIO(P8A1)
	
	P7A1 := RK3288.CreateGPIO(RK3288.P7, RK3288.A1)
	P7A1.SetDir(RK3288.Output)
	P7A1.SetPP(RK3288.PP_PU)
	P7A1.SetSR(RK3288.SR_Fast)
	P7A1.SetE(RK3288.E_12mA)
	P7A1.Flip()
	fmt.Println(P7A1.GetLevel())
	time.Sleep(time.Second)
	P7A1.Flip()
	fmt.Println(P7A1.GetLevel())
	RK3288.FreeGPIO(P7A1)
	
	PWM1 := RK3288.CreatePWM(RK3288.PWM_1)
	PWM1.SetInactivePolarity(RK3288.Positive)
	PWM1.SetPERIOD(500000)
	PWM1.SetDUTY(250000)
	PWM1.SetCNT(0)
	PWM1.Start()
	time.Sleep(time.Second * 5)
	fmt.Println(PWM1.GetCNT())
	PWM1.Stop()
	RK3288.FreePWM(PWM1)

	defer RK3288.FreeRK3288()
}

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
	defer RK3288.FreeRK3288()
	
	//所有的Create具有是否创建成功返回值
	P8A1, ok := RK3288.CreateGPIO(RK3288.P8, RK3288.A1)
	if ok {
		P8A1.Flip()
		fmt.Println(P8A1.GetData())
		RK3288.FreeGPIO(P8A1)
	}
	
	//可以忽略是否创建成功的返回值
	P7A1, _ := RK3288.CreateGPIO(RK3288.P7, RK3288.A1)
	P7A1.SetDir(RK3288.Output)
	P7A1.SetPP(RK3288.PP_PU)
	P7A1.SetSR(RK3288.SR_Fast)
	P7A1.SetE(RK3288.E_12mA)
	P7A1.Flip()
	fmt.Println(P7A1.GetData())
	time.Sleep(time.Second)
	P7A1.Flip()
	fmt.Println(P7A1.GetData())
	RK3288.FreeGPIO(P7A1)
	
	PWM1, _ := RK3288.CreatePWM(RK3288.PWM_1)
	PWM1.SetInactivePolarity(RK3288.Positive)
	PWM1.SetPERIOD(500000)
	PWM1.SetDUTY(250000)
	PWM1.SetCNT(0)
	PWM1.Start()
	time.Sleep(time.Second * 5)
	fmt.Println(PWM1.GetCNT())
	PWM1.Stop()
	RK3288.FreePWM(PWM1)

	//所有的单例没有创建是否成功的返回值，直接使用
	Data :=RK3288.ITSADC().GetData(RK3288.TSADC_1)
	fmt.Println(Data)
	fmt.Println(RK3288.ITSADC().GetTemperature(Data))
	RK3288.FreeTSADC()
	
	//判断单例是否成功的方法
	WD := RK3288.IWDT()
	if WD != nil {
		WD.FeedSecond = 10
		WD.Start()
		time.Sleep(time.Second * 60)
		WD.Stop() //一旦启动不能被关闭，所以此句无效
		RK3288.FreeWDT()
	}
}

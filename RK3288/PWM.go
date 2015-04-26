/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.26]
*/

package RK3288

import (
	
)

const (PWM_0, PWM_1, PWM_2, PWM_3, PWM_4 = 0, 1, 2, 3, 4)
const (OneShot, Continuous, Capture = 0, 1, 2)
const (Left, Center = 0, 1)

const BasePWM = 0xFF680000

type PWM struct {
	pwmX 		uint8
	hMem		[]uint8
	
	PWM_CNT	*uint32
	PWM_PERIOD	*uint32
	PWM_DUTY	*uint32
	PWM_CTRL	*uint32
}

func CreatePWM(PWMx uint8) *PWM {
	pwm := &PWM{pwmX: PWMx}
	
	pwm.hMem, _ = IRK3288().GetMMap(BasePWM + int64(pwm.pwmX) * 0x10)
	pwm.PWM_CNT, _ = IRK3288().Register(pwm.hMem, 0x0000)
	pwm.PWM_PERIOD, _ = IRK3288().Register(pwm.hMem, 0x0004)
	pwm.PWM_DUTY, _ = IRK3288().Register(pwm.hMem, 0x0008)
	pwm.PWM_CTRL, _ = IRK3288().Register(pwm.hMem, 0x000C)
	
	switch pwm.pwmX {
		case PWM_0:	IGRF().IOMUX_PWM(PWM0_IOMUX)
		case PWM_1:	IGRF().IOMUX_PWM(PWM1_IOMUX)
		case PWM_2:	IGRF().IOMUX_PWM(PWM2_IOMUX)
		case PWM_3:	IGRF().IOMUX_PWM(PWM3_IOMUX)
		//case PWM_4:	IGRF().IOMUX_PWM(PWM4_IOMUX)
	}
	
	return pwm
}

func FreePWM(pwm *PWM) {
	IRK3288().FreeMMap(pwm.hMem)
}

func (this *PWM) GetState() bool {
	return (*this.PWM_CTRL & 0x1) == 0x1
}

func (this *PWM) GetCNT() uint32 {
	return *this.PWM_CNT
}

func (this *PWM) GetPERIOD() uint32 {
	return *this. PWM_PERIOD
}

func (this *PWM) GetDUTY() uint32 {
	return *this.PWM_DUTY
}

func (this *PWM) SetCNT(cnt uint32) {
	*this.PWM_CNT = cnt
}

func (this *PWM) SetPERIOD(period uint32) {
	*this.PWM_PERIOD = period
}

func (this *PWM) SetDUTY(duty uint32) {
	*this.PWM_DUTY = duty
}

func (this *PWM) Start() {
	*this.PWM_CTRL |= 0x1
}

func (this *PWM) Stop() {
	*this.PWM_CTRL &^= 0x1
}

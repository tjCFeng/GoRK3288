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
const (AlignedLeft, AlignedCenter = 0, 1)
const (Negative, Positive = 0, 1)

const BasePWM = 0xFF680000

type PWM struct {
	pwmX 		uint8
	hMem		[]uint8
	
	PWM_CNT	*uint32
	PWM_PERIOD	*uint32
	PWM_DUTY	*uint32
	PWM_CTRL	*uint32
}

func CreatePWM(PWMx uint8) (*PWM, bool) {
	var Result bool = false
	var offset uint32

	if PWMx > PWM_4 { return nil, Result}
	pwm := &PWM{pwmX: PWMx}
	offset = uint32(PWMx * 0x10)

	pwm.hMem, Result = IRK3288().GetMMap(BasePWM)
	if !Result { return nil, Result}
	
	pwm.PWM_CNT, _ = IRK3288().Register(pwm.hMem, 0x0000 + offset)
	pwm.PWM_PERIOD, _ = IRK3288().Register(pwm.hMem, 0x0004 + offset)
	pwm.PWM_DUTY, _ = IRK3288().Register(pwm.hMem, 0x0008 + offset)
	pwm.PWM_CTRL, _ = IRK3288().Register(pwm.hMem, 0x000C + offset)
	
	switch pwm.pwmX {
		case PWM_0:	IGRF().IOMUX_PWM(PWM0_IOMUX)
		case PWM_1:	IGRF().IOMUX_PWM(PWM1_IOMUX)
		case PWM_2:	IGRF().IOMUX_PWM(PWM2_IOMUX)
		case PWM_3:	IGRF().IOMUX_PWM(PWM3_IOMUX)
		//case PWM_4:	IGRF().IOMUX_PWM(PWM4_IOMUX)
	}
	
	*pwm.PWM_CTRL = (0x7 << 12) + (0x1 << 1)
	
	return pwm, Result
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

func (this *PWM) SetOutputMode(mode uint8) {
	switch mode {
		case 0: *this.PWM_CTRL &^= (0x1 << 5)
		default: *this.PWM_CTRL |= (0x1 << 5)
	}
}

func (this *PWM) SetInactivePolarity(polarity uint8) {
	switch polarity {
		case 0: *this.PWM_CTRL &^= (0x1 << 4)
		default: *this.PWM_CTRL |= (0x1 << 4)
	}
}

func (this *PWM) SetDutyPolarity(polarity uint8) {
	switch polarity {
		case 0: *this.PWM_CTRL &^= (0x1 << 3)
		default: *this.PWM_CTRL |= (0x1 << 3)
	}
}

func (this *PWM) Start() {
	*this.PWM_CTRL |= 0x1
}

func (this *PWM) Stop() {
	*this.PWM_CTRL &^= 0x1
}


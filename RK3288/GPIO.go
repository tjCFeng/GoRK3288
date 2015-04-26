/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.26]
*/

package RK3288

import (
	
)

const (P1, P2, P3, P4, P5, P6, P7, P8 = 0, 1, 2, 3, 4, 5, 6, 7)

const (
	A0, A1, A2, A3, A4, A5, A6, A7 = 0, 1, 2, 3, 4, 5, 6, 7
	B0, B1, B2, B3, B4, B5, B6, B7 = 8, 9, 10, 11, 12, 13, 14, 15
	C0, C1, C2, C3, C4, C5, C6, C7 = 16, 17, 18, 19, 20, 21, 22, 23
	D0, D1, D2, D3, D4, D5, D6, D7 = 24, 25, 26, 27, 28, 29, 30, 31
)

const (
	Input = false
	Output = true
)

const BaseGPIO = 0xFF780000

/*****************************************************************************/
type GPIOGROUP struct {
	port				uint8
	hMem			[]uint8
	
	SWPORT_DR		*uint32
	SWPORT_DDR		*uint32
	INTEN			*uint32
	INTMASK		*uint32
	INTTYPE_LEVEL	*uint32
	INT_POLARITY	*uint32
	INT_STATUS		*uint32
	INT_RAWSTATUS	*uint32
	DEBOUNCE		*uint32
	PORT_EOI		*uint32
	EXT_PORT		*uint32
	FLS_SYNC		*uint32
}

func CreateGPIOGROUP(Port uint8) *GPIOGROUP {
	gpiogroup := &GPIOGROUP{port: Port}
	
	gpiogroup.hMem, _ = IRK3288().GetMMap(BaseGPIO + int64(gpiogroup.port) * 0x10000)
	gpiogroup.SWPORT_DR, _ = IRK3288().Register(gpiogroup.hMem, 0x0000)
	gpiogroup.SWPORT_DDR, _ = IRK3288().Register(gpiogroup.hMem, 0x0004)
	gpiogroup.INTEN, _ = IRK3288().Register(gpiogroup.hMem, 0x0030)
	gpiogroup.INTMASK, _ = IRK3288().Register(gpiogroup.hMem, 0x0034)
	gpiogroup.INTTYPE_LEVEL, _ = IRK3288().Register(gpiogroup.hMem, 0x0038)
	gpiogroup.INT_POLARITY, _ = IRK3288().Register(gpiogroup.hMem, 0x003C)
	gpiogroup.INT_STATUS, _ = IRK3288().Register(gpiogroup.hMem, 0x0040)
	gpiogroup.INT_RAWSTATUS, _ = IRK3288().Register(gpiogroup.hMem, 0x0044)
	gpiogroup.DEBOUNCE, _ = IRK3288().Register(gpiogroup.hMem, 0x0048)
	gpiogroup.PORT_EOI, _ = IRK3288().Register(gpiogroup.hMem, 0x004C)
	gpiogroup.EXT_PORT, _ = IRK3288().Register(gpiogroup.hMem, 0x0050)
	gpiogroup.FLS_SYNC, _ = IRK3288().Register(gpiogroup.hMem, 0x0060)
	
	return gpiogroup
}

func FreeGPIOGROUP(gpiogroup *GPIOGROUP) {
	IRK3288().FreeMMap(gpiogroup.hMem)
	FreeRK3288()
}

/*****************************************************************************/
type GPIO struct {
	port				*GPIOGROUP
	pin				uint8
	bit				uint32
}

func CreateGPIO(Port uint8, Pin uint8)  *GPIO {
	gpio := &GPIO{}
	gpio.port = CreateGPIOGROUP(Port)
	gpio.pin = Pin
	gpio.bit = (0x1 << gpio.pin)

	return gpio
}

func FreeGPIO(gpio *GPIO) {
	FreeGPIOGROUP(gpio.port)
}

func (this *GPIO) SetLevel(level bool) {
	switch level {
		case true: *this.port.SWPORT_DR |= (0x1 << this.pin)
		case false: *this.port.SWPORT_DR &^= (0x1 << this.pin)
	}
}

func (this *GPIO) GetLevel() bool {
	return (*this.port.SWPORT_DR & this.bit) == this.bit
}

func (this *GPIO) Flip() {
	*this.port.SWPORT_DR ^= (0x1 << this.pin)
}

func (this *GPIO) SetDir(dir bool) {
	switch dir {
		case true: *this.port.SWPORT_DDR |= (0x1 << this.pin)
		case false: *this.port.SWPORT_DDR &^= (0x1 << this.pin)
	}
}

/*func (this *GPIO) SetDebounce(deb bool) {
	switch deb {
		case true: *this.port.DEBOUNCE |= (0x1 << this.pin)
		case false: *this.port.DEBOUNCE &^= (0x1 << this.pin)
	}
}*/

func (this *GPIO) GetInputValue() bool {
	return (*this.port.EXT_PORT & this.bit) == this.bit
}

/*func main() {
	P8A1 := CreateGPIO(P8, A1)
	P8A1.Flip()
	FreeGPIO(P8A1)
}*/

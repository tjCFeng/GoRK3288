/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.26]
*/

package RK3288

import (

)

const BaseGRF = 0xFF770000

var iGRF *GRF = nil

func IGRF() (*GRF) {
	if (iGRF == nil) {
		iGRF = &GRF{hMem: nil}
		iGRF.hMem, _ = IRK3288().GetMMap(BaseGRF)
		
		iGRF.GPIO1D_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x000C)
		iGRF.GPIO2A_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0010)
		iGRF.GPIO2B_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0014)
		iGRF.GPIO2C_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0018)
		iGRF.GPIO3A_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0020)
		iGRF.GPIO3B_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0024)
		iGRF.GPIO3C_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0028)
		iGRF.GPIO3DL_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x002C)
		iGRF.GPIO3DH_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0030)
		iGRF.GPIO4AL_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0034)
		iGRF.GPIO4AH_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0038)
		iGRF.GPIO4BL_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x003C)
		iGRF.GPIO4C_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0044)
		iGRF.GPIO4D_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0048)
		iGRF.GPIO5B_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0050)
		iGRF.GPIO5C_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0054)
		iGRF.GPIO6A_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x005C)
		iGRF.GPIO6B_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0060)
		iGRF.GPIO6C_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0064)
		iGRF.GPIO7A_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x006C)
		iGRF.GPIO7B_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0070)
		iGRF.GPIO7CL_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0074)
		iGRF.GPIO7CH_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0078)
		iGRF.GPIO8A_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0080)
		iGRF.GPIO8B_IOMUX, _ = IRK3288().Register(iGRF.hMem, 0x0084)
	}

	return iGRF
}

func FreeGRF(grf *GRF) {
	IRK3288().FreeMMap(grf.hMem)
}

/*****************************************************************************/

const  (
	PWM0_IOMUX = iota
	PWM1_IOMUX
	PWM2_IOMUX
	PWM3_IOMUX
	PWM4_IOMUX
	I2C0_IOMUX
	I2C1_IOMUX
	I2C2_IOMUX
	I2C3_IOMUX
	I2C4_IOMUX
	I2C5_IOMU
	SPI0_CS0_IOMUX
	SPI0_CS1_IOMUX
	SPI1_CS0_IOMUX
	SPI1_CS1_IOMUX
	SPI2_CS0_IOMUX
	SPI2_CS1_IOMU
	UART0_IOMUX
	UART1_IOMUX
	UART2_IOMUX
	UART3_IOMUX
	UART4_IOMU
	LCDC0_IOMUX
	LCDC1_IOMU
	EMMC_IOMU
	SDCARD_IOMU
	HDMI_IOMUX
)

type GRF struct {
	hMem			[]uint8
	
	//Function Select : GPIO | ...
	GPIO1D_IOMUX	*uint32
	GPIO2A_IOMUX	*uint32
	GPIO2B_IOMUX	*uint32
	GPIO2C_IOMUX	*uint32
	GPIO3A_IOMUX	*uint32
	GPIO3B_IOMUX	*uint32
	GPIO3C_IOMUX	*uint32
	GPIO3DL_IOMUX	*uint32
	GPIO3DH_IOMUX	*uint32
	GPIO4AL_IOMUX	*uint32
	GPIO4AH_IOMUX	*uint32
	GPIO4BL_IOMUX	*uint32
	GPIO4C_IOMUX	*uint32
	GPIO4D_IOMUX	*uint32
	GPIO5B_IOMUX	*uint32
	GPIO5C_IOMUX	*uint32
	GPIO6A_IOMUX	*uint32
	GPIO6B_IOMUX	*uint32
	GPIO6C_IOMUX	*uint32
	GPIO7A_IOMUX	*uint32
	GPIO7B_IOMUX	*uint32
	GPIO7CL_IOMUX	*uint32
	GPIO7CH_IOMUX	*uint32
	GPIO8A_IOMUX	*uint32
	GPIO8B_IOMUX	*uint32
	
	//GPIO Rate Select : Slow | Fast
	GPIO1H_SR		*uint32
	GPIO2L_SR		*uint32
	GPIO2H_SR		*uint32
	GPIO3L_SR		*uint32
	GPIO3H_SR		*uint32
	GPIO4L_SR		*uint32
	GPIO4H_SR		*uint32
	GPIO5L_SR		*uint32
	GPIO5H_SR		*uint32
	GPIO6L_SR		*uint32
	GPIO6H_SR		*uint32
	GPIO7L_SR		*uint32
	GPIO7H_SR		*uint32
	GPIO8L_SR		*uint32
	
	//GPIO PU/PD Select : None | PU | PD | Repeater
	GPIO1D_P		*uint32
	GPIO2A_P		*uint32
	GPIO2B_P		*uint32
	GPIO2C_P		*uint32
	GPIO3A_P		*uint32
	GPIO3B_P		*uint32
	GPIO3C_P		*uint32
	GPIO3D_P		*uint32
	GPIO4A_P		*uint32
	GPIO4B_P		*uint32
	GPIO4C_P		*uint32
	GPIO4D_P		*uint32
	GPIO5B_P		*uint32
	GPIO5C_P		*uint32
	GPIO6A_P		*uint32
	GPIO6B_P		*uint32
	GPIO6C_P		*uint32
	GPIO7A_P		*uint32
	GPIO7B_P		*uint32
	GPIO7C_P		*uint32
	GPIO8A_P		*uint32
	GPIO8B_P		*uint32
	
	//GPIO Drive Strength Select : 2 | 4 | 8 | 12 mA
	GPIO1D_E		*uint32
	GPIO2A_E		*uint32
	GPIO2B_E		*uint32
	GPIO2C_E		*uint32
	GPIO3A_E		*uint32
	GPIO3B_E		*uint32
	GPIO3C_E		*uint32
	GPIO3D_E		*uint32
	GPIO4A_E		*uint32
	GPIO4B_E		*uint32
	GPIO4C_E		*uint32
	GPIO4D_E		*uint32
	GPIO5B_E		*uint32
	GPIO5C_E		*uint32
	GPIO6A_E		*uint32
	GPIO6B_E		*uint32
	GPIO6C_E		*uint32
	GPIO7A_E		*uint32
	GPIO7B_E		*uint32
	GPIO7C_E		*uint32
	GPIO8A_E		*uint32
	GPIO8B_E		*uint32
}

func (this *GRF) IOMUX_GPIO(Port uint8, Pin uint8)  {
	switch Port {
		case P7:
			switch Pin {
				case A0:
					*this.GPIO7A_IOMUX |= (0x3 << 16)
					*this.GPIO7A_IOMUX &^= (0x3 << 0)
					//*this.GPIO7A_IOMUX &^= (0x3 << 16)
				case A1:
					*this.GPIO7A_IOMUX |= (0x1 << 18)
					*this.GPIO7A_IOMUX &^= (0x1 << 2)
					//*this.GPIO7A_IOMUX &^= (0x3 << 16)
			}
		default:
			return
	}
}

func (this *GRF) IOMUX_PWM(PWMx uint8)  {
	switch PWMx {
		case PWM0_IOMUX:
			*this.GPIO7A_IOMUX |= (0x3 << 16)
			*this.GPIO7A_IOMUX &^= (0x3 << 0)
			*this.GPIO7A_IOMUX |= (0x1 << 0)
			//*this.GPIO7A_IOMUX &^= (0x3 << 16)
		case PWM1_IOMUX:
			*this.GPIO7A_IOMUX |= (0x1 << 18)
			*this.GPIO7A_IOMUX |= (0x1 << 2)
			//*this.GPIO7A_IOMUX &^= (0x1 << 18)
		default:
			return
	}
}

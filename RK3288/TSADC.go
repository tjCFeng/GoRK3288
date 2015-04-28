/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.28]
*/

package RK3288

import (
	
)

const (TSADC_0, TSADC_1, TSADC_2, TSADC_3 = 7, 6, 5, 4)

const BaseTSADC = 0xFF280000

var iTSADC *TSADC = nil

func ITSADC() (*TSADC) {
	if (iTSADC == nil) {
		iTSADC = &TSADC{hMem: nil}
		iTSADC.bit = (0x1 << 12)
		iTSADC.hMem, _ = IRK3288().GetMMap(BaseTSADC)
		
		iTSADC.USER_CON, _ = IRK3288().Register(iTSADC.hMem, 0x0000)
		iTSADC.AUTO_CON, _ = IRK3288().Register(iTSADC.hMem, 0x0004)
		iTSADC.INT_EN, _ = IRK3288().Register(iTSADC.hMem, 0x0008)
		iTSADC.INT_PD, _ = IRK3288().Register(iTSADC.hMem, 0x000C)
		iTSADC.DATA0, _ = IRK3288().Register(iTSADC.hMem, 0x0020)
		iTSADC.DATA1, _ = IRK3288().Register(iTSADC.hMem, 0x0024)
		iTSADC.DATA2, _ = IRK3288().Register(iTSADC.hMem, 0x0028)
		iTSADC.DATA3, _ = IRK3288().Register(iTSADC.hMem, 0x002C)
		iTSADC.COMP0_INT, _ = IRK3288().Register(iTSADC.hMem, 0x0030)
		iTSADC.COMP1_INT, _ = IRK3288().Register(iTSADC.hMem, 0x0034)
		iTSADC.COMP2_INT, _ = IRK3288().Register(iTSADC.hMem, 0x0038)
		iTSADC.COMP3_INT, _ = IRK3288().Register(iTSADC.hMem, 0x003C)
		iTSADC.COMP0_SHU, _ = IRK3288().Register(iTSADC.hMem, 0x0040)
		iTSADC.COMP1_SHU, _ = IRK3288().Register(iTSADC.hMem, 0x0044)
		iTSADC.COMP2_SHU, _ = IRK3288().Register(iTSADC.hMem, 0x0048)
		iTSADC.COMP3_SHU, _ = IRK3288().Register(iTSADC.hMem, 0x004C)
		iTSADC.HIGHT_INT_DEBOUNCE, _ = IRK3288().Register(iTSADC.hMem, 0x0060)
		iTSADC.HIGHT_TSHUT_DEBOUNCE, _ = IRK3288().Register(iTSADC.hMem, 0x0064)
		iTSADC.AUTO_PERIOD, _ = IRK3288().Register(iTSADC.hMem, 0x0068)
		iTSADC.AUTO_PERIOD_HT, _ = IRK3288().Register(iTSADC.hMem, 0x006C)
		
		*iTSADC.USER_CON = (0x3F << 6) + (0x1 << 3)
	}

	return iTSADC
}

func FreeTSADC() {
	if iTSADC != nil {
		IRK3288().FreeMMap(iTSADC.hMem)
	}
}

type TSADC struct {
	hMem					[]uint8
	bit						uint32
	
	USER_CON				*uint32
	AUTO_CON				*uint32
	INT_EN					*uint32
	INT_PD					*uint32
	DATA0					*uint32
	DATA1					*uint32
	DATA2					*uint32
	DATA3					*uint32
	COMP0_INT				*uint32
	COMP1_INT				*uint32
	COMP2_INT				*uint32
	COMP3_INT				*uint32
	COMP0_SHU				*uint32
	COMP1_SHU				*uint32
	COMP2_SHU				*uint32
	COMP3_SHU				*uint32
	HIGHT_INT_DEBOUNCE		*uint32
	HIGHT_TSHUT_DEBOUNCE	*uint32
	AUTO_PERIOD			*uint32
	AUTO_PERIOD_HT			*uint32
}

func (this *TSADC) GetData(channel uint8) uint32 {
	Converted := make(chan uint8)
	
	*iTSADC.USER_CON &^= (0x7 << 0)
	*iTSADC.USER_CON |= uint32(channel << 0)
	*iTSADC.USER_CON |= (0x1 << 5)

	go func() {
		for {
			if (*this.USER_CON & this.bit == 0) {break}	
		}
		Converted <- 1
	}()
	
	<-Converted

	switch channel {
		case TSADC_0: return *this.DATA0
		case TSADC_1: return *this.DATA1
		case TSADC_2: return *this.DATA2
		case TSADC_3: return *this.DATA3
		default: return 0
	}
}

var DataTable = []uint32{
	3800, 3792, 3783, 3774, 3765, 3756, 3747, 3737, 
	3728, 3718, 3708, 3698, 3688, 3678, 3667, 3656,
	3645, 3634, 3623, 3611, 3600, 3588, 3575, 3563,
	3550, 3537, 3524, 3510, 3496, 3482, 3467, 3452,
	3437, 3421,
}

func (this *TSADC) GetTemperature(data uint32) float32 {
	if data >3800 {return -255}
	if data < 3421 {return 255}

	var I = 0
	for I = 0; I < len(DataTable); I++ {
		if data >= DataTable[I] {break}
	}

	Temp := 5 / float32(DataTable[I - 1] - DataTable[I])
	Temp *= float32(data - DataTable[I])
	Temp += float32(-40 + I * 5)
	
	return Temp
}

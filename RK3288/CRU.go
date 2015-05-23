/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.05.22]
*/

package RK3288

import (
	
)

const BaseCRU = 0xFF760000

var iCRU *CRU = nil

func ICRU() *CRU {
	if (iCRU == nil) {
		iCRU = &CRU{}
		iCRU.hMem, _ = IRK3288().GetMMap(BaseCRU)
		
		for I:= 0; I < 3; I++ { iCRU.APLL[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0000)) }
		for I:= 0; I < 3; I++ { iCRU.DPLL[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0010)) }
		for I:= 0; I < 3; I++ { iCRU.CPLL[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0020)) }
		for I:= 0; I < 3; I++ { iCRU.GPLL[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0030)) }
		for I:= 0; I < 3; I++ { iCRU.NPLL[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0040)) }
		iCRU.MODE, _ = IRK3288().Register(iCRU.hMem, 0x0050)
		for I:= 0; I < 43; I++ { iCRU.CLKSEL[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0060)) }
		for I:= 0; I < 19; I++ { iCRU.CLKGATE[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0160)) }
		iCRU.GLB_SRST_FST_VALUE, _ = IRK3288().Register(iCRU.hMem, 0x01B0)
		iCRU.GLB_SRST_SND_VALUE, _ = IRK3288().Register(iCRU.hMem, 0x01B4)
		for I:= 0; I < 12; I++ { iCRU.SOFTRST[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x01B8)) }
		iCRU.MISC, _ = IRK3288().Register(iCRU.hMem, 0x01E8)
		iCRU.GLB_CNT_TH, _ = IRK3288().Register(iCRU.hMem, 0x01EC)
		iCRU.GLB_RST, _ = IRK3288().Register(iCRU.hMem, 0x01F0)
		iCRU.GLB_RST_ST, _ = IRK3288().Register(iCRU.hMem, 0x01F8)
		for I:= 0; I < 2; I++ { iCRU.SDMMC[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0200)) }
		for I:= 0; I < 2; I++ { iCRU.SDIO0[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0208)) }
		for I:= 0; I < 2; I++ { iCRU.SDIO1[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0210)) }
		for I:= 0; I < 2; I++ { iCRU.EMMC[I], _ = IRK3288().Register(iCRU.hMem, uint32(0x4 * I + 0x0218)) }
	}
	
	return iCRU
}

func FreeCRU() {
	if iCRU != nil {
		IRK3288().FreeMMap(iCRU.hMem)
	}
}

type CRU struct {
	hMem					[]uint8
	
	APLL						[4]*uint32
	DPLL						[4]*uint32
	CPLL						[4]*uint32
	GPLL						[4]*uint32
	NPLL						[4]*uint32
	MODE					*uint32
	CLKSEL					[43]*uint32
	CLKGATE				[19]*uint32
	GLB_SRST_FST_VALUE		*uint32
	GLB_SRST_SND_VALUE	*uint32
	SOFTRST				[12]*uint32
	MISC						*uint32
	GLB_CNT_TH		*uint32
	GLB_RST				*uint32
	GLB_RST_ST		*uint32
	SDMMC				[2]*uint32
	SDIO0					[2]*uint32
	SDIO1					[2]*uint32
	EMMC					[2]*uint32
}

/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.05.02]
*/

package RK3288

import (
	"time"
)

const BaseWDT = 0xFF800000

var iWDT *WDT = nil

func IWDT() (*WDT) {
	if (iWDT == nil) {
		iWDT = &WDT{hMem: nil, running:false, FeedSecond:3}
		iWDT.hMem, _ = IRK3288().GetMMap(BaseWDT)
		
		iWDT.CR, _ = IRK3288().Register(iWDT.hMem, 0x0000)
		iWDT.TORR, _ = IRK3288().Register(iWDT.hMem, 0x0004)
		iWDT.CCVR, _ = IRK3288().Register(iWDT.hMem, 0x0008)
		iWDT.CRR, _ = IRK3288().Register(iWDT.hMem, 0x000C)
		iWDT.STAT, _ = IRK3288().Register(iWDT.hMem, 0x0010)
		iWDT.EOI, _ = IRK3288().Register(iWDT.hMem, 0x0014)
		
		*iWDT.CR = (0x3 << 2) + (0x1 << 1) //0x0000001C
		*iWDT.TORR = 0x0000000F
	}
	
	return iWDT
}

func FreeWDT() {
	if iWDT != nil {
		IRK3288().FreeMMap(iWDT.hMem)
	}
}

type WDT struct {
	hMem				[]uint8
	running			bool
	FeedSecond	uint8
	
	CR						*uint32
	TORR					*uint32
	CCVR					*uint32
	CRR					*uint32
	STAT					*uint32
	EOI					*uint32
}

func (this *WDT) Start() {
	this.running = true
	*this.CR |= 0x1
	this.FeedWD()
}

func (this *WDT) Stop() { //一旦启动不能被关闭
	*this.CR &^= 0x1
	this.running = false
}

func (this *WDT) FeedWD() {
	go func() {
		for {
			*this.CRR = 0x76
			if this.running == false { return }
			time.Sleep(time.Second * time.Duration(this.FeedSecond))
	}
	}()
}

func (this *WDT) Reset() {
	*this.CR = 0x1
}

/*
	Author: tjCFeng(LiuYang)
	EMail: tjCFeng@163.com
	[2015.04.26]
*/

package RK3288

import (
	"os"
	"unsafe"
	"syscall"
)

var iRK3288 *RK3288 = nil

func IRK3288() (*RK3288) {
	if (iRK3288 == nil) {
		iRK3288 = &RK3288{nil}
		iRK3288.hFile, _ = os.OpenFile("/dev/mem", os.O_RDWR, 0)
	}

	return iRK3288
}

func FreeRK3288() {
	FreeGRF(IGRF())
	iRK3288.hFile.Close()
}

/*****************************************************************************/

type RK3288 struct {
	hFile	*os.File
}

func (this *RK3288) GetMMap(PhysicalAddr int64) ([]uint8, bool) {
	if this.hFile == nil {
		return nil, false
	}

	mem, err := syscall.Mmap(int(this.hFile.Fd()), PhysicalAddr, os.Getpagesize(), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	return mem, (err != nil)
}

func (this *RK3288) FreeMMap(hMem []uint8) {
	syscall.Munmap(hMem)
}

func (this *RK3288) Register(hMem []uint8, offset uint32) (*uint32, bool) {
	if hMem == nil {
		return nil, false
	}
	return  (*uint32)(unsafe.Pointer(&hMem[offset])), true
}

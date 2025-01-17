//go:build linux || android || openwrt
// +build linux android openwrt

package rkvendorstorage

import (
	"encoding/binary"
	"os"
	"syscall"
	"unsafe"
)

// Write writes a value into the vendor storage with the given ID as the key.
//
// If the size of the input exceeds RequestMaxSize, the request will be denied and no data will be written.
func Write(id uint16, data []byte) (err error) {
	if len(data) > RequestMaxSize {
		return ErrorDataTooLong
	}

	f, err := os.OpenFile(DevicePath, syscall.O_RDWR, 0)
	if err != nil {
		return
	}
	defer f.Close()

	buf := make([]byte, RequestBufferAllocationSize)
	binary.LittleEndian.PutUint32(buf[0:4], RequestTag)
	binary.LittleEndian.PutUint16(buf[4:6], id)
	binary.LittleEndian.PutUint16(buf[6:8], uint16(len(data)))
	copy(buf[8:], data)

	err = IOCtl(f.Fd(), IORequestWrite, uintptr(unsafe.Pointer(&buf[0])))
	return
}

//go:build !(linux || android || openwrt)
// +build !linux,!android,!openwrt

package rkvendorstorage

import "errors"

// Read reads the value corresponding to the given key from the vendor storage.
//
// If the on-disk data size exceeds RequestMaxSize, only the first [0:RequestMaxSize] bytes will be returned, and err
// will be set to ErrorDataTooLong. There is no way to read the trailing parts of the data.
func Read(id uint16) (size uint16, data []byte, err error) {
	err = errors.New("not support")
	return
}

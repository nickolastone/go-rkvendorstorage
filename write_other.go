//go:build !(linux || android || openwrt)
// +build !linux,!android,!openwrt

package rkvendorstorage

import (
	"errors"
)

// Write writes a value into the vendor storage with the given ID as the key.
//
// If the size of the input exceeds RequestMaxSize, the request will be denied and no data will be written.
func Write(id uint16, data []byte) (err error) {
	err = errors.New("not support")
	return
}

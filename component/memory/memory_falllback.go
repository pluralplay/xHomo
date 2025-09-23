//go:build !darwin && !linux && !freebsd && !openbsd && !windows

package memory

import "errors"

var ErrNotImplementedError = errors.New("not implemented yet")

func GetMemoryInfo(pid int32) (*MemoryInfoStat, error) {
	return nil, ErrNotImplementedError
}

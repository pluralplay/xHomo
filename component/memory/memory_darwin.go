package memory

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

const (
	PROC_PIDTASKINFO = 4
)

type ProcTaskInfo struct {
	Virtual_size      uint64
	Resident_size     uint64
	Total_user        uint64
	Total_system      uint64
	Threads_user      uint64
	Threads_system    uint64
	Policy            int32
	Faults            int32
	Pageins           int32
	Cow_faults        int32
	Messages_sent     int32
	Messages_received int32
	Syscalls_mach     int32
	Syscalls_unix     int32
	Csw               int32
	Threadnum         int32
	Numrunning        int32
	Priority          int32
}

// Library represents a dynamic library loaded by purego.
type Library struct {
	addr  uintptr
	path  string
	close func()
}

func NewLibrary(path string) (*Library, error) {
	lib, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return nil, err
	}

	closeFunc := func() {
		purego.Dlclose(lib)
	}

	return &Library{
		addr:  lib,
		path:  path,
		close: closeFunc,
	}, nil
}

func (lib *Library) Dlsym(symbol string) (uintptr, error) {
	return purego.Dlsym(lib.addr, symbol)
}

func GetFunc[T any](lib *Library, symbol string) T {
	var fptr T
	purego.RegisterLibFunc(&fptr, lib.addr, symbol)
	return fptr
}

func (lib *Library) Close() {
	lib.close()
}

// library paths
const (
	System = "/usr/lib/libSystem.B.dylib"
)

// System functions and symbols.
type (
	ProcPidInfoFunc func(pid, flavor int32, arg uint64, buffer uintptr, bufferSize int32) int32
)

const (
	ProcPidInfoSym = "proc_pidinfo"
)

type dlFuncs struct {
	lib *Library

	procPidInfo ProcPidInfoFunc
}

func loadProcFuncs() (*dlFuncs, error) {
	lib, err := NewLibrary(System)
	if err != nil {
		return nil, err
	}

	return &dlFuncs{
		lib:         lib,
		procPidInfo: GetFunc[ProcPidInfoFunc](lib, ProcPidInfoSym),
	}, nil
}

func (f *dlFuncs) Close() {
	f.lib.Close()
}

func GetMemoryInfo(pid int32) (*MemoryInfoStat, error) {
	funcs, err := loadProcFuncs()
	if err != nil {
		return nil, err
	}
	defer funcs.Close()

	var ti ProcTaskInfo
	funcs.procPidInfo(pid, PROC_PIDTASKINFO, 0, uintptr(unsafe.Pointer(&ti)), int32(unsafe.Sizeof(ti)))

	ret := &MemoryInfoStat{
		RSS: uint64(ti.Resident_size),
		VMS: uint64(ti.Virtual_size),
	}
	return ret, nil
}

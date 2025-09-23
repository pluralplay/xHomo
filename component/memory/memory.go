// Package memory return MemoryInfoStat
// modify from https://github.com/shirou/gopsutil/tree/v4.25.8/process
package memory

type MemoryInfoStat struct {
	RSS uint64 `json:"rss"` // bytes
	VMS uint64 `json:"vms"` // bytes
}

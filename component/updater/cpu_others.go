//go:build !amd64

package updater

// getGOAMD64level is always return 0 when not in amd64 platfrom.
func getGOAMD64level() int32 {
	return 0
}

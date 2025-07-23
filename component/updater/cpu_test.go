package updater

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGOAMD64level(t *testing.T) {
	level := getGOAMD64level()
	fmt.Printf("GOAMD64=%d\n", level)
	if runtime.GOARCH == "amd64" {
		assert.True(t, level > 0)
		assert.True(t, level <= 4)
	} else {
		assert.Equal(t, level, int32(0))
	}
}

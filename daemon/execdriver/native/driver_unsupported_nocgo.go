// +build linux,!cgo

package native

import (
	"fmt"

	"github.com/dotcloud/docker/daemon/execdriver"
)

func NewDriver(root, initPath string) (execdriver.Driver, error) {
	return nil, fmt.Errorf("native driver not supported on non-linux")
}

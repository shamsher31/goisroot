// Package root Check if the process is running as root user, eg. started with sudo
package root // import "github.com/shamsher31/goisroot"

import (
	"os"
)

func Is() bool {
	if os.Getuid() == 0 {
		return true
	}
	return false
}

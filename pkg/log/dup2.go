//go:build !windows

package log

import (
	"os"
	"syscall"
)

var (
	stdout = syscall.Stdout
	stderr = syscall.Stderr
)

func dup2(f *os.File, fd int) error {
	return syscall.Dup3(int(f.Fd()), fd, 0)
}

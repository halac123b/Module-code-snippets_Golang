package main

import "syscall"

const (
	// syscall.SIGTERM, tín hiệu để kết thúc chương trình
	sigTerm = syscall.Signal(15)
)

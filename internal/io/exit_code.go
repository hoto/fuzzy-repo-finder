package io

type ExitCode int

const (
	CONTINUE ExitCode = iota + 1
	NORMAL_EXIT
	ABNORMAL_EXIT
)

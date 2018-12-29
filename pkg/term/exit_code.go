package term

type ExitCode int

const (
	ContinueRunning ExitCode = iota + 1
	NormalExit
	AbnormalExit
)

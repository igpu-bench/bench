package job

type JobState int

const (
	Empty JobState = iota
	Pending
	Running
	Complete
	Failed
	Suspended
)

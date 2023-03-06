package job

import (
	"github.com/google/uuid"
	"github.com/igpu-bench/ibench/common/result"
)

type Job struct {
	id     uuid.UUID
	result result.Result
}

func New() *Job {
	var j *Job

	uu, _ := uuid.NewUUID()
	j.id = uu

	return j
}

func (j *Job) Start() error {
	// TODO actual processing of the job
	return nil
}

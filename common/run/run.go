package run

import (
	"github.com/google/uuid"
	"github.com/igpu-bench/ibench/common/job"
)

type Run struct {
	id   uuid.UUID
	jobs []*job.Job
}

func New() *Run {
	var r *Run

	uu, _ := uuid.NewUUID()
	r.id = uu

	return r
}

func (run *Run) StartAll() error {
	// TODO
	return nil
}

package scheduler

import "GoAuth/src/scheduler/job"

func Init() error {
	err := job.CleanUpToken()
	if err != nil {
		return err
	}

	return nil
}

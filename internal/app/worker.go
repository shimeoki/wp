package app

// An action that needs to be done in a transaction with the passed provider.
// If action is successful, nil should be returned.
type Job[P Provider] func(P) error

// A factory for the jobs.
type Worker[P Provider] interface {
	// Creates a new job with the specified context. This context should be
	// reused in the job as well. If an error was encountered while creating
	// the job, error should be returned. If an error encountered in the job,
	// the error should be returned as well.
	Work(Ctx, Job[P]) error
}

// An adapter to satisfy the Worker interface.
type WorkerFunc[P Provider] func(Ctx, Job[P]) error

func (f WorkerFunc[P]) Work(ctx Ctx, j Job[P]) error {
	return f(ctx, j)
}

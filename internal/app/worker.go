package app

type Job[P Provider] func(P) error

type Worker[P Provider] interface {
	Work(Ctx, Job[P]) error
}

type WorkerFunc[P Provider] func(Ctx, Job[P]) error

func (f WorkerFunc[P]) Work(ctx Ctx, j Job[P]) error {
	return f(ctx, j)
}

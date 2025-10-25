package app

type Worker[P Provider] interface {
	Work(Ctx, func(P) error) error
}

type WorkerFunc[P Provider] func(Ctx, func(P) error) error

func (f WorkerFunc[P]) Work(ctx Ctx, fn func(P) error) error {
	return f(ctx, fn)
}

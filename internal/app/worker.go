package app

type Worker[P Providers] interface {
	Do(Ctx, func(P) error) error
}

type WorkerFunc[P Providers] func(Ctx, func(P) error) error

func (f WorkerFunc[P]) Do(ctx Ctx, fn func(P) error) error {
	return f(ctx, fn)
}

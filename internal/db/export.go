package db

import (
	"context"
	"io"
)

type Exporter interface {
	Export(ctx context.Context, out io.Writer) error
}

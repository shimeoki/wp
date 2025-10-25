package main

import (
	"context"

	"github.com/shimeoki/wp/internal/cli"
	"github.com/shimeoki/wp/internal/config"
)

func main() {
	app := config.NewApp(config.New())
	cmd := cli.New(app)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd.Execute(ctx)
}

package main

import (
	"context"

	"github.com/shimeoki/wp/internal/cli"
	"github.com/shimeoki/wp/internal/config"
)

func main() {
	cli.New(config.NewApp(config.New())).Execute(context.Background())
}

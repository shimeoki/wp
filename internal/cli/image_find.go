package cli

import (
	"fmt"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) imageFindCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "find [hash]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.FindWallpaper

			res, err := h.Handle(cmd.Context(), &app.FindWallpaperQuery{
				Hash: args[0],
			})

			if err != nil {
				fatal(err)
			}

			fmt.Println(res.Format)
		},
	}
}

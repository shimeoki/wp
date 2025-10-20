package cli

import (
	"io"
	"os"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) imageShowCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "show [hash]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.ShowWallpaper

			res, err := h.Handle(cmd.Context(), &app.ShowWallpaperQuery{
				Hash: args[0],
			})

			if err != nil {
				fatal(err)
			}

			defer res.Image.Close()
			if _, err := io.Copy(os.Stdout, res.Image); err != nil {
				fatal(err)
			}
		},
	}
}

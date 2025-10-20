package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) imageCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "create [path]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.CreateWallpaper

			img, err := os.Open(args[0])
			if err != nil {
				fatal(err)
			}

			defer img.Close()
			res, err := h.Handle(cmd.Context(), &app.CreateWallpaperCommand{
				Image:  img,
				Format: filepath.Ext(args[0])[1:],
			})

			if err != nil {
				fatal(err)
			}

			fmt.Println(res.Hash)
		},
	}
}

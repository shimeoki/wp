package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) imageDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "delete [hash]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.DeleteWallpaper

			if _, err := h.Handle(cmd.Context(), &app.DeleteWallpaperCommand{
				Hash: args[0],
			}); err != nil {
				fatal(err)
			}
		},
	}
}

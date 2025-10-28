package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) sourceRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "remove [hash] [id]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.RemoveSource

			if _, err := h.Handle(cmd.Context(), &app.RemoveSourceCommand{
				WallpaperHash: args[0],
				SourceID:      args[1],
			}); err != nil {
				fatal(err)
			}
		},
	}
}

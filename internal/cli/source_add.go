package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) sourceAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "add [hash] [id]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.AddSource

			if _, err := h.Handle(cmd.Context(), &app.AddSourceCommand{
				WallpaperHash: args[0],
				SourceID:      args[1],
			}); err != nil {
				cli.fatal(err)
			}
		},
	}
}

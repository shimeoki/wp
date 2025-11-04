package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) tagRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "remove [hash] [name]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.RemoveTag

			if _, err := h.Handle(cmd.Context(), &app.RemoveTagCommand{
				WallpaperHash: args[0],
				TagName:       args[1],
			}); err != nil {
				cli.fatal(err)
			}
		},
	}
}

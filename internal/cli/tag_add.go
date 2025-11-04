package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) tagAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "add [hash] [name]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.AddTag

			if _, err := h.Handle(cmd.Context(), &app.AddTagCommand{
				WallpaperHash: args[0],
				TagName:       args[1],
			}); err != nil {
				cli.fatal(err)
			}
		},
	}
}

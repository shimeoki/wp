package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) aliasRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "remove [hash] [name]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.RemoveAlias

			if _, err := h.Handle(cmd.Context(), &app.RemoveAliasCommand{
				WallpaperHash: args[0],
				AliasName:     args[1],
			}); err != nil {
				fatal(err)
			}
		},
	}
}

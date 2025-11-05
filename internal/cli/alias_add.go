package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) aliasAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "add [hash] [name]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.AddAlias

			if _, err := h.Handle(cmd.Context(), &app.AddAliasCommand{
				WallpaperHash: args[0],
				AliasName:     args[1],
			}); err != nil {
				cli.fatal(err)
			}
		},
	}
}

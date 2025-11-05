package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) sourceDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "delete [id]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.DeleteSource

			_, err := h.Handle(cmd.Context(), &app.DeleteSourceCommand{
				ID: args[0],
			})

			if err != nil {
				cli.fatal(err)
			}
		},
	}
}

package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) tagCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "create [name]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.CreateTag

			_, err := h.Handle(cmd.Context(), &app.CreateTagCommand{
				Name: args[0],
			})

			if err != nil {
				fatal(err)
			}
		},
	}
}

package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) tagDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "delete [name]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.DeleteTag

			_, err := h.Handle(cmd.Context(), &app.DeleteTagCommand{
				Name: args[0],
			})

			if err != nil {
				fatal(err)
			}
		},
	}
}

package cli

import (
	"fmt"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) sourceCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "create [name] [link?]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.CreateSource

			var link *string
			if len(args) > 1 {
				link = &args[1]
			}

			r, err := h.Handle(cmd.Context(), &app.CreateSourceCommand{
				Name: args[0],
				Link: link,
			})

			if err != nil {
				cli.fatal(err)
			}

			fmt.Println(r.ID)
		},
	}
}

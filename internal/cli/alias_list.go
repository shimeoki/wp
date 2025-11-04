package cli

import (
	"fmt"
	"strings"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) aliasListCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.ListAliases

			res, err := h.Handle(cmd.Context(), &app.ListAliasesQuery{})
			if err != nil {
				fatal(err)
			}

			fmt.Println(strings.Join(res.Aliases, "\n"))
		},
	}
}

package cli

import (
	"fmt"
	"strings"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) tagListCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.ListTags

			res, err := h.Handle(cmd.Context(), &app.ListTagsQuery{})
			if err != nil {
				fatal(err)
			}

			fmt.Println(strings.Join(res.Names, "\n"))
		},
	}
}

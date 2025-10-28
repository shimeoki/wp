package cli

import (
	"fmt"
	"strings"

	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) sourceListCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.ListSources

			res, err := h.Handle(cmd.Context(), &app.ListSourcesQuery{})
			if err != nil {
				fatal(err)
			}

			var lines []string
			for _, s := range res.List {
				line := fmt.Sprintf("id: %s, name: %s", s.ID, s.Name)
				if s.Link != nil {
					line += fmt.Sprintf(", link: %s", *s.Link)
				}

				lines = append(lines, line)
			}

			fmt.Println(strings.Join(lines, "\n"))
		},
	}
}

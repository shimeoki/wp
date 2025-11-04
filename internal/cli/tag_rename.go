package cli

import (
	"github.com/shimeoki/wp/internal/app"
	"github.com/spf13/cobra"
)

func (cli *CLI) tagRenameCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "rename [before] [after]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			h := cli.handlers.RenameTag

			if _, err := h.Handle(cmd.Context(), &app.RenameTagCommand{
				Before: args[0],
				After:  args[1],
			}); err != nil {
				cli.fatal(err)
			}
		},
	}
}

package cli

import "github.com/spf13/cobra"

func (cli *CLI) tagCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tag",
	}

	cmd.AddCommand(
		cli.tagCreateCommand(),
		cli.tagDeleteCommand(),
		cli.tagListCommand(),
		cli.tagAddCommand(),
		cli.tagRemoveCommand(),
		cli.tagRenameCommand(),
	)

	return cmd
}

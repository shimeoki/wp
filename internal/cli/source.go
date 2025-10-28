package cli

import "github.com/spf13/cobra"

func (cli *CLI) sourceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "source",
	}

	cmd.AddCommand(
		cli.sourceCreateCommand(),
		cli.sourceDeleteCommand(),
		cli.sourceListCommand(),
		cli.sourceAddCommand(),
		cli.sourceRemoveCommand(),
	)

	return cmd
}

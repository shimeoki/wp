package cli

import "github.com/spf13/cobra"

func (cli *CLI) aliasCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "alias",
	}

	cmd.AddCommand(
		cli.aliasListCommand(),
	)

	return cmd
}

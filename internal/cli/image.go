package cli

import "github.com/spf13/cobra"

func (cli *CLI) imageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "image",
	}

	cmd.AddCommand(
		cli.imageCreateCommand(),
		cli.imageDeleteCommand(),
		cli.imageFindCommand(),
		cli.imageShowCommand(),
	)

	return cmd
}

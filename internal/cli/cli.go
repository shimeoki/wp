package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/shimeoki/wp/internal/config"
	"github.com/spf13/cobra"
)

type CLI struct {
	cmd      *cobra.Command
	cfg      *config.Config
	app      *config.App
	handlers *config.Handlers
}

func New(cfg *config.Config) *CLI {
	cli := &CLI{cfg: cfg}
	cli.cmd = cli.command()
	return cli
}

func (cli *CLI) Execute(ctx context.Context) {
	if err := cli.cmd.ExecuteContext(ctx); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func (cli *CLI) command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "wp",

		// runs before every children command - initialize an app instance
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cli.app != nil {
				return
			}

			a, err := config.NewApp(cmd.Context(), cli.cfg)
			if err != nil {
				fatal(err)
			}

			cli.app = a
			cli.handlers = a.Handlers()
		},

		// cleanup after the commands
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			cli.app.Close()
		},
	}

	flags := cmd.PersistentFlags()

	flags.StringVar(
		&cli.cfg.DB.DataSourceName,
		"db-dsn",
		cli.cfg.DB.DataSourceName,
		"database data source name",
	)

	flags.StringVar(
		&cli.cfg.Store.Path,
		"store-path",
		cli.cfg.Store.Path,
		"store location",
	)

	cmd.AddCommand(
		cli.imageCommand(),
		cli.tagCommand(),
	)

	return cmd
}

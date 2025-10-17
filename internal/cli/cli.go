package cli

import (
	"fmt"
	"os"

	"github.com/shimeoki/wp/internal/config"
	"github.com/spf13/cobra"
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use: "wp",
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func init() {
	initTag()
	initImage()

	rootCmd.PersistentFlags().StringVar(
		&cfg.DB.DataSourceName, "db-dsn", "",
		"database data source name")

	rootCmd.PersistentFlags().StringVar(
		&cfg.Store.Path, "store-path", "",
		"store location")

	rootCmd.AddCommand(tagCmd)
	rootCmd.AddCommand(imageCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fatal(err)
	}
}

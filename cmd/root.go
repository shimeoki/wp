package cmd

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

func init() {
	rootCmd.PersistentFlags().StringVar(
		&cfg.DB.DataSourceName, "db-dsn",
		"", "database data source name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

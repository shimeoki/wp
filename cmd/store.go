package cmd

import (
	"context"
	"io"
	"os"

	"github.com/shimeoki/wp/internal/api"
	"github.com/shimeoki/wp/internal/db"
	"github.com/spf13/cobra"
)

var storeFile string

func storeReader() io.ReadCloser {
	if storeFile == "" {
		return os.Stdin
	}

	file, err := os.Open(storeFile)
	if err != nil {
		fatal(err)
	}

	return file
}

func storeWriter() io.WriteCloser {
	if storeFile == "" {
		return os.Stdout
	}

	file, err := os.Create(storeFile)
	if err != nil {
		fatal(err)
	}

	return file
}

func jsoner() *api.JSONer {
	repo, err := db.NewSQLiteRepo(&cfg.DB)
	if err != nil {
		fatal(err)
	}

	return api.NewJSONer(repo)
}

func storeExportRun(cmd *cobra.Command, args []string) {
	w := storeWriter()
	defer w.Close()

	if err := jsoner().Export(context.Background(), w); err != nil {
		fatal(err)
	}
}

func storeImportRun(cmd *cobra.Command, args []string) {
	r := storeReader()
	defer r.Close()

	if err := jsoner().Import(context.Background(), r); err != nil {
		fatal(err)
	}
}

var storeCmd = &cobra.Command{
	Use: "store",
}

var storeExportCmd = &cobra.Command{
	Use:  "export",
	Args: cobra.NoArgs,
	Run:  storeExportRun,
}

var storeImportCmd = &cobra.Command{
	Use:  "import",
	Args: cobra.NoArgs,
	Run:  storeImportRun,
}

func initStore() {
	storeCmd.PersistentFlags().StringVarP(
		&storeFile, "file", "f", "",
		"path of the file to use (stdin or stdout if blank)")

	storeCmd.AddCommand(storeExportCmd)
	storeCmd.AddCommand(storeImportCmd)
}

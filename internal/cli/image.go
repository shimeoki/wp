package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shimeoki/wp/internal/app"
	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use: "image",
}

var imageCreateCmd = &cobra.Command{
	Use:  "create [path]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		repo := sqlite.NewWallpaperRepo(openDB(ctx))

		store := openStore()
		defer store.Close()

		h := app.NewCreateWallpaperHandler(store, repo)

		img, err := os.Open(args[0])
		if err != nil {
			fatal(err)
		}

		defer img.Close()
		res, err := h.Handle(ctx, &app.CreateWallpaperCommand{
			Image:  img,
			Format: filepath.Ext(args[0])[1:],
		})

		if err != nil {
			fatal(err)
		}

		fmt.Println(res.Hash)
	},
}

var imageDeleteCmd = &cobra.Command{
	Use:  "delete [hash]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		repo := sqlite.NewWallpaperRepo(openDB(ctx))

		store := openStore()
		defer store.Close()

		h := app.NewDeleteWallpaperHandler(store, repo)

		_, err := h.Handle(ctx, &app.DeleteWallpaperCommand{Hash: args[0]})
		if err != nil {
			fatal(err)
		}
	},
}

var imageFindCmd = &cobra.Command{
	Use:  "find [hash]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		repo := sqlite.NewWallpaperRepo(openDB(ctx))

		h := app.NewFindWallpaperHandler(repo)

		res, err := h.Handle(ctx, &app.FindWallpaperQuery{Hash: args[0]})
		if err != nil {
			fatal(err)
		}

		fmt.Println(res.Format)
	},
}

var imageShowCmd = &cobra.Command{
	Use:  "show [hash]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		repo := sqlite.NewWallpaperRepo(openDB(ctx))

		store := openStore()
		defer store.Close()

		h := app.NewShowWallpaperHandler(store, repo)

		res, err := h.Handle(ctx, &app.ShowWallpaperQuery{Hash: args[0]})
		if err != nil {
			fatal(err)
		}

		defer res.Image.Close()
		if _, err := io.Copy(os.Stdout, res.Image); err != nil {
			fatal(err)
		}
	},
}

func initImage() {
	imageCmd.AddCommand(imageCreateCmd)
	imageCmd.AddCommand(imageDeleteCmd)
	imageCmd.AddCommand(imageFindCmd)
	imageCmd.AddCommand(imageShowCmd)
}

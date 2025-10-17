package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/shimeoki/wp/internal/app"
	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use: "tag",
}

var tagCreateCmd = &cobra.Command{
	Use:  "create [name]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := sqlite.Open(ctx, &cfg.DB)
		if err != nil {
			fatal(err)
		}

		h := app.NewCreateTagHandler(sqlite.NewTagRepo(conn))

		_, err = h.Handle(ctx, &app.CreateTagCommand{Name: args[0]})
		if err != nil {
			fatal(err)
		}
	},
}

var tagDeleteCmd = &cobra.Command{
	Use:  "delete [name]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := sqlite.Open(ctx, &cfg.DB)
		if err != nil {
			fatal(err)
		}

		h := app.NewDeleteTagHandler(sqlite.NewTagRepo(conn))

		_, err = h.Handle(ctx, &app.DeleteTagCommand{Name: args[0]})
		if err != nil {
			fatal(err)
		}
	},
}

var tagListCmd = &cobra.Command{
	Use:  "list",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := sqlite.Open(ctx, &cfg.DB)
		if err != nil {
			fatal(err)
		}

		h := app.NewListTagsHandler(sqlite.NewTagRepo(conn))

		res, err := h.Handle(ctx, &app.ListTagsQuery{})
		if err != nil {
			fatal(err)
		}

		fmt.Println(strings.Join(res.Names, "\n"))
	},
}

func initTag() {
	tagCmd.AddCommand(tagCreateCmd)
	tagCmd.AddCommand(tagDeleteCmd)
	tagCmd.AddCommand(tagListCmd)
}

package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"libragen/util"
	"log"
	"os/exec"
)

// techMojoCmd represents the taskrm command
var libragenCmd = &cobra.Command{
	Use:   "lg",
	Short: "libragen",
	Long:  `只能在我家里面的电脑使用`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jekyllDir := `D:\code\libragen.cn`
		err := util.ParseUrlPageLibraGen(args[0], "div.article__content", jekyllDir)
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		commd := exec.CommandContext(ctx, "bundle", "exec", "jekyll", "serve")
		commd.Dir = jekyllDir
		err = commd.Start()
		if err != nil {
			cancel()
		}
	},
}

func init() {
	rootCmd.AddCommand(libragenCmd)
}
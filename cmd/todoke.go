package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unoharu/hikyaku/internal/edo"
	"github.com/unoharu/hikyaku/internal/fileops"
)

var todokeCmd = &cobra.Command{
	Use:   "todoke [src] [dst]",
	Short: "届け（移動）: ファイルを移動し、元場所を空き地にする",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		dst := args[1]
		fmt.Printf("届けるぜ！[%s] から [%s] へ、跡形もなく運んでやる！\n", src, dst)

		if err := fileops.Move(src, dst); err != nil {
			fmt.Println(edo.ErrorMessage(err))
			return err
		}

		fmt.Println("ガッテンだ！無事に荷を届けたぜ。受け取りの判をもらってきな！")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todokeCmd)
}

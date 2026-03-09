package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var todokeCmd = &cobra.Command{
	Use:   "todoke [src] [dst]",
	Short: "届け（移動）: ファイルを移動し、元場所を空き地にする",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		dst := args[1]
		fmt.Printf("届けるぜ！[%s] から [%s] へ、跡形もなく運んでやる！\n", src, dst)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todokeCmd)
}

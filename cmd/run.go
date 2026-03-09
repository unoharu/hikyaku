package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unoharu/hikyaku/internal/edo"
	"github.com/unoharu/hikyaku/internal/fileops"
)

var runCmd = &cobra.Command{
	Use:   "run [src] [dst]",
	Short: "荷運び(コピー): 飛脚が走りながらファイルをコピーする",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		dst := args[1]
		fmt.Printf("走るぜ！[%s] から [%s] へ届けてみせる！\n", src, dst)
		
		if err := fileops.Copy(src, dst); err != nil {
			fmt.Println(edo.ErrorMessage(err))
			return err			
		}
	
		fmt.Println("ガッテンだ！無事に荷を届けたぜ。受け取りの判をもらってきな！")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

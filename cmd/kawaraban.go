package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unoharu/hikyaku/internal/edo"
)

var kawarabanCmd = &cobra.Command{
	Use:   "kawaraban [path]",
	Short: "瓦版（一覧）: ディレクトリを木札風レイアウトで表示する",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		fmt.Printf("≡ 瓦版 ≡ [%s]\n", path)
		fmt.Println("────────────────────")

		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("📁 %s/\n", entry.Name())
			} else {
				info, err := entry.Info()
				if err != nil {
					return err
				}
				fmt.Printf("📄 %-20s %s\n", entry.Name(), edo.FormatSize(info.Size()))
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(kawarabanCmd)
}

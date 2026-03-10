package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unoharu/hikyaku/internal/edo"
	"github.com/unoharu/hikyaku/internal/store"
)

var meiboCmd = &cobra.Command{
	Use:   "meibo",
	Short: "台帳（履歴）: 過去の転送記録を表示する",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := store.Load()
		if err != nil {
			return err
		}

		if len(entries) == 0 {
			fmt.Println("台帳が空っぽだぜ。まずは荷を運んでみな！")
			return nil
		}

		fmt.Println("≡ 飛脚台帳 ≡")
		fmt.Println("────────────────────────────────────────────")

		var totalBytes int64
		for _, e := range entries {
			date := edo.FormatDate(e.Time)
			size := edo.FormatSize(e.Bytes)
			fortune := e.Fortune
			if fortune == "" {
				fortune = "−"
			}
			fmt.Printf("%s  %s → %s  %s  %s\n", date, e.Src, e.Dst, size, fortune)
			totalBytes += e.Bytes
		}

		fmt.Println("────────────────────────────────────────────")
		fmt.Printf("累計：%s\n", edo.FormatSize(totalBytes))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(meiboCmd)
}

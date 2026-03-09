package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/unoharu/hikyaku/internal/edo"
	"github.com/unoharu/hikyaku/internal/fileops"
	"github.com/unoharu/hikyaku/internal/ui"
)

var runCmd = &cobra.Command{
	Use:   "run [src] [dst]",
	Short: "荷運び(コピー): 飛脚が走りながらファイルをコピーする",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		dst := args[1]
		fmt.Printf("走るぜ！[%s] から [%s] へ届けてみせる！\n", src, dst)

		info, err := os.Stat(src)
		if err != nil {
			return err
		}
		fmt.Println(edo.FormatSize(info.Size()))
		fmt.Println(edo.WeightComment(info.Size()))

		w := edo.RandomWeather()
		fmt.Printf("%s 「%s」\n", w.Label, w.Line)

		if err := fileops.Copy(src, dst); err != nil {
			fmt.Println(edo.ErrorMessage(err))
			return err
		}

		// プログレスバーを表示（疑似進捗）
		p := tea.NewProgram(ui.NewModel())
		if _, err := p.Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

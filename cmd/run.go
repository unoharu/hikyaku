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

var runKakugo bool

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

		if _, err := os.Stat(dst); err == nil {
			if runKakugo {
				fmt.Println("上書きしたぞ。後悔すんなよ。")
			} else {
				fmt.Print("おっと、そこには先客がいるようだ。蹴散らして（上書き）も構わねぇかい？ [y/n]: ")
				var answer string
				fmt.Scanln(&answer)
				if answer != "y" && answer != "Y" {
					fmt.Println("そうかい、引き返すなら今のうちだぜ。")
					return nil
				}
			}
		}

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
	runCmd.Flags().BoolVarP(&runKakugo, "kakugo", "k", false, "上書き確認をスキップする（覚悟の上で）")
}

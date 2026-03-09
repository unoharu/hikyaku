package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/unoharu/hikyaku/internal/edo"
	"github.com/unoharu/hikyaku/internal/fileops"
	"github.com/unoharu/hikyaku/internal/store"
)

var todokeKakugo bool
var todokeYonige bool

var todokeCmd = &cobra.Command{
	Use:   "todoke [src] [dst]",
	Short: "届け（移動）: ファイルを移動し、元場所を空き地にする",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		dst := args[1]

		if !todokeYonige {
			fmt.Printf("届けるぜ！[%s] から [%s] へ、跡形もなく運んでやる！\n", src, dst)
		}

		info, err := os.Stat(src)
		if err != nil {
			return err
		}

		if info.Size() == 0 {
			fmt.Println(edo.EmptyFileMessage)
			return nil
		}

		if !todokeYonige {
			fmt.Println(edo.FormatSize(info.Size()))
			fmt.Println(edo.WeightComment(info.Size()))
			w := edo.RandomWeather()
			fmt.Printf("%s 「%s」\n", w.Label, w.Line)
		}

		if _, err := os.Stat(dst); err == nil {
			if todokeKakugo {
				if !todokeYonige {
					fmt.Println("上書きしたぞ。後悔すんなよ。")
				}
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

		if err := fileops.Move(src, dst); err != nil {
			fmt.Println(edo.ErrorMessage(err))
			return err
		}

		fortune, err := edo.FortuneFromFile(dst)
		if err != nil {
			fortune = ""
		}

		if err := store.Append(store.Entry{
			Time:    time.Now(),
			Src:     src,
			Dst:     dst,
			Bytes:   info.Size(),
			Fortune: fortune,
		}); err != nil {
			fmt.Printf("台帳への記録に失敗したぜ: %v\n", err)
		}

		if !todokeYonige {
			fmt.Println("ガッテンだ！無事に荷を届けたぜ。受け取りの判をもらってきな！")
			if fortune != "" {
				fmt.Printf("おみくじ：【%s】%s\n", fortune, edo.FortuneMessage(fortune))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todokeCmd)
	todokeCmd.Flags().BoolVarP(&todokeKakugo, "kakugo", "k", false, "上書き確認をスキップする（覚悟の上で）")
	todokeCmd.Flags().BoolVarP(&todokeYonige, "yonige", "y", false, "静音モード：メッセージを非表示")
}

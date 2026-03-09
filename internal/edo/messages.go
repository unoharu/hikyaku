package edo

import (
	"errors"
	"os"
)

func ErrorMessage(err error) string {
	if errors.Is(err, os.ErrNotExist) {
		return "てやんでぇ！存在しねぇ荷物を運べってのか？化かされてんじゃねぇよ！"
	}
	if errors.Is(err, os.ErrPermission) {
		return "おとといきやがれ！ここは一般人の立ち入りは禁止されてんだ！"
	}
	return "どうも具合が悪いぜ：" + err.Error()
}

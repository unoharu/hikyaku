package edo

import (
	"fmt"
	"time"
)

var kanjiMonth = [13]string{
	"", "一月", "二月", "三月", "四月", "五月", "六月",
	"七月", "八月", "九月", "十月", "十一月", "十二月",
}

var kanjiDay = [32]string{
	"", "一日", "二日", "三日", "四日", "五日", "六日", "七日",
	"八日", "九日", "十日", "十一日", "十二日", "十三日", "十四日",
	"十五日", "十六日", "十七日", "十八日", "十九日", "二十日",
	"二十一日", "二十二日", "二十三日", "二十四日", "二十五日",
	"二十六日", "二十七日", "二十八日", "二十九日", "三十日", "三十一日",
}

// FormatDate は time.Time を「天保元年三月一日」形式の文字列に変換する
// 天保元年 = 1830年 を起点とした簡易換算
func FormatDate(t time.Time) string {
	year := t.Year() - 1830 + 1
	month := int(t.Month())
	day := t.Day()
	return fmt.Sprintf("天保%d年%s%s", year, kanjiMonth[month], kanjiDay[day])
}

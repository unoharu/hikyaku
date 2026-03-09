package edo

import "fmt"

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

func FormatSize(bytes int64) string {
	switch {
	case bytes < KB:
		return fmt.Sprintf("%d 厘", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.1f 分", float64(bytes)/float64(KB))
	case bytes < GB:
		return fmt.Sprintf("%.1f 寸", float64(bytes)/float64(MB))
	case bytes < TB:
		return fmt.Sprintf("%.1f 尺", float64(bytes)/float64(GB))
	default:
		return fmt.Sprintf("%.1f 丈", float64(bytes)/float64(TB))
	}
}

func WeightComment(bytes int64) string {
	switch {
	case bytes < MB: // 〜1MB
		return "これくらい羽より軽いぜ"
	case bytes < 100*MB: // 1MB〜100MB
		return "ちょうどいい荷だ"
	case bytes < GB: // 100MB〜1GB
		return "うっ…重い。これは米俵か？"
	default:
		return "こいつは…無理言うなよ旦那…"
	}
}

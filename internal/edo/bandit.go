package edo

import (
	"fmt"
	"math/rand"
	"time"
)

// MaybeBandit は1%の確率で盗賊イベントを発生させる
// yonige（静音モード）が true のときは何もしない
func MaybeBandit(yonige bool) {
	if yonige {
		return
	}
	if rand.Intn(100) != 0 {
		return
	}
	fmt.Println("泥棒だー！荷物を狙いやがった！")
	time.Sleep(5 * time.Second)
	fmt.Println("取り返したぜ！行くぞ！")
}

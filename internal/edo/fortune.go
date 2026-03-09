package edo

import (
	"crypto/sha256"
	"io"
	"os"
)

var fortunes = []string{
	"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶",
}

var fortuneMessages = map[string]string{
	"大吉": "このコードはバグが出ねぇ予感だぜ！",
	"吉":  "今日は順風満帆だ。どんどん運べ！",
	"中吉": "まずまずの荷だな。悪くはねぇぜ。",
	"小吉": "ちょっと気を引き締めて行くとしよう。",
	"末吉": "焦らず慎重にな。急ぎすぎは禁物だぜ。",
	"凶":  "気をつけろ！今日は厄日かもしれねぇ。",
	"大凶": "旦那、今日は大人しくしておいた方がいいぜ…",
}

// FortuneFromFile はファイルの SHA-256 ハッシュの末尾バイトで運勢を決定する
func FortuneFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	sum := h.Sum(nil)
	fortune := fortunes[sum[len(sum)-1]%byte(len(fortunes))]
	return fortune, nil
}

// FortuneMessage は運勢名に対応するメッセージを返す
func FortuneMessage(fortune string) string {
	return fortuneMessages[fortune]
}

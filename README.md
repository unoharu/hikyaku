# 飛脚（Hikyaku）

江戸時代の飛脚が荷物を運ぶ体験として昇華した、Go 製のファイル操作 CLI ツール。
`cp` / `mv` の機能を、粋な江戸っ子口調と東海道53次の旅で味わえる。

## インストール

### バイナリ直接ダウンロード

[Releases](https://github.com/unoharu/hikyaku/releases) から OS に合ったアーカイブをダウンロードして展開する。

### ソースからビルド

```bash
git clone https://github.com/unoharu/hikyaku.git
cd hikyaku
go build -o hikyaku .
```

## 使い方

### `run` — 荷運び（コピー）

```bash
hikyaku run ./src.txt ./dst.txt
```

- 東海道53次の宿場町を巡りながらファイルをコピーする
- 完了後におみくじで運勢を表示
- 1%の確率で盗賊が出現し、5秒停止する

### `todoke` — 届け（移動）

```bash
hikyaku todoke ./src.txt ./dst.txt
```

- ファイルを移動し、元の場所を空き地にする

### `kawaraban` — 瓦版（一覧）

```bash
hikyaku kawaraban ./some-dir
```

- ディレクトリの中身を木札風に表示する
- ファイルサイズは江戸単位（厘・分・寸・尺・丈）で表示

### `meibo` — 台帳（履歴）

```bash
hikyaku meibo
```

- 過去の転送記録と運勢を台帳形式で表示する
- 履歴は `~/.hikyaku/log.json` に保存される

## フラグ

| フラグ | 短縮 | 説明 |
| :--- | :--- | :--- |
| `--kakugo` | `-k` | 上書き確認をスキップする |
| `--yonige` | `-y` | 静音モード（メッセージ・プログレスバーを非表示） |

```bash
# 上書き確認なしでコピー
hikyaku run ./src.txt ./dst.txt --kakugo

# 静かに移動（バックグラウンド実行向け）
hikyaku todoke ./src.txt ./dst.txt --yonige
```

## 江戸単位

| 現代単位 | 江戸単位 |
| :--- | :--- |
| 1 Byte | 1 厘 |
| 1 KB | 1 分 |
| 1 MB | 1 寸 |
| 1 GB | 1 尺 |
| 1 TB | 1 丈 |

## ライセンス

MIT

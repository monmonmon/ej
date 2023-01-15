# ej - 英和辞書コマンド

[英和辞書データ "ejdict-hand"](https://github.com/kujirahand/EJDict) (SQLite 形式版) のためのコマンドラインインターフェース

## インストール

1. 上記 URL から SQLite 形式の辞書データベースファイルをダウンロード、解凍し、 `ejdict.sqlite3` を `~/lib` 以下に設置
1. 本リポジトリを取得、 `go build` して出来たバイナリ `ej` をパスの通ったディレクトリへ移動

## 使い方

```bash
$ ej hello
hello
    (電話の応答で)『もしもし』;(あいさつ・呼びかけ・驚きの声などに用いて)『こんにちは』,やあ,おい,ちょっと,おや,まあ
    こんにちは(やあ,おいなど)という呼びかけ(あいさつ)

$ ej japan
japan
         黒い漆(うるし);漆器(しっき)
         …‘に'黒漆を塗る
Japan
         『日本』
Japan Current
         =Black Stream
Japanese
         『日本の』;『日本人』(『語』)『の』
         〈C〉『日本人』;《集合的に》日本人
         〈U〉《冠詞をつけないで》『日本語』
Japanese beetle
         マメコガネ(植物の害虫)
Japanese lantern
         =Chinese lantern
Japanese quince
         ボケ
Japanize
         …‘を'日本化する,日本ふうにする
Japanise
         …‘を'日本化する,日本ふうにする
```

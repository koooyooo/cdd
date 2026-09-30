# シェル補完

## 概要

`cdd init` が出力するシェル関数に、登録済みエイリアス名の補完を載せる。Cobra 標準の `cdd completion` はサブコマンドとフラグ用のままとし、ジャンプ先の名前は出さない。

## 要件

- `eval "$(cdd init bash)"` または `eval "$(cdd init zsh)"` のあと、`cdd <Tab>` でエイリアス名とサブコマンドが出る
- エイリアス名は `list` と同じ順。番号は候補にしない
- 次のサブコマンド（と短縮名）の直後の引数はエイリアス名: `remove` `rm` `delete` `del` `move-up` `up` `move-down` `down` `print` `p`
- `init` の次の引数は `bash` と `zsh`
- `add` `a` のパス引数（3番目の語）はファイル名補完。新しいエイリアス名の位置では補完しない
- 第1引数以外で、上記に当てはまらない位置ではファイル名を出さない
- zsh は `compdef` があるときだけ補完を登録する。`compinit` 前でもジャンプ用の関数は動く
- bash は 3.2（macOS の `/bin/bash`）で動く
- 補完は `command cdd` を呼び、シェル関数の再帰を避ける
- `completion` と補完用の隠しコマンドは、エイリアスへのジャンプとして扱わない

## 設計

- 隠しコマンド `complete-aliases` がエイリアス名を1行1件で stdout に出す。`list` の表はパースしない
- 名前の取得は `repo.List()` の順。空の名前は捨てる
- `init` は bash と zsh で同じ `cdd` 関数を出し、そのあとにシェル別の補完定義を出す
- bash は `complete -F`。zsh は `compadd` と `compdef`
- フラグを挟む補完（例: `cdd print -o <Tab>`）はしない

## 検証

- `go test ./...`
- 生成スクリプトの `bash -n` / `zsh -n`
- bash では偽の `cdd` を PATH に置き、候補の中身を確認する

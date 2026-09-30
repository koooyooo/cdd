![cdd](./logo.jpg)

# cdd

[English](./README.md) | **日本語**

[![test](https://github.com/koooyooo/cdd/actions/workflows/test.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/test.yaml)
[![lint](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml)

事前登録したディレクトリへ、短いエイリアスでジャンプする `cd` です。

深い階層への移動を、名前か番号だけで済ませられます。パス解決は `cdd` バイナリが行い、実際の移動はシェルの `builtin cd` に任せます（シェル連携が必要です）。

```bash
# Before
$ cd ../../Documents/projects/cdd/chart

# After
$ cdd chart
```

## インストール

```bash
go install github.com/koooyooo/cdd@latest
```

シェル連携を有効にします（zsh の例）:

```bash
# 対話: 関数定義を表示し、必要なら eval 行をクリップボードへコピー
cdd init zsh

# または rc に直接追記
echo 'eval "$(cdd init zsh)"' >> ~/.zshrc
```

bash の場合は `cdd init bash` と `~/.bashrc` を使います。

zsh でエイリアス名を補完するには、`compinit` のあとに `eval` します。

```bash
autoload -Uz compinit && compinit
eval "$(cdd init zsh)"
```

bash は追加の設定なしで `cdd <Tab>` がエイリアス名とサブコマンドを出します。

## 使い方

`$ cdd {command}` で各種コマンドを実行します。引数なしの `cdd` はヘルプを表示します。

### ジャンプ（シェル連携時）

エイリアス名、または `list` の番号を渡すと、対象ディレクトリへ `cd` します。

```bash
$ cdd docs
$ pwd
/Users/me/Documents

$ cdd 1
$ pwd
/Users/me/Documents
```

### コマンド一覧

| コマンド | 短縮 | 説明 |
| --- | --- | --- |
| `print` | `p` | エイリアスを絶対パスに解決して stdout へ出力 |
| `init` | — | bash / zsh 向けシェル連携コードを出力 |
| `list` | — | 登録済みエイリアスを一覧表示 |
| `add` | — | エイリアスを追加 |
| `remove` | `rm` | エイリアスを削除 |
| `up` / `down` | — | 一覧上の並び順を変更 |
| `edit` | — | 設定ファイルをエディタで開く |

#### `print` / `p`

エイリアス名・番号を絶対パスに解決します。シェル連携やスクリプトから利用します。

```bash
$ cdd print docs
/Users/me/Documents

$ cd "$(cdd print docs)"
```

#### `init`

bash / zsh 向けのシェル連携コードを stdout に出力します。ジャンプ用の関数に加え、エイリアス名とサブコマンドの補完も登録します。`remove` / `up` / `down` / `print` の次の引数はエイリアス名、`init` の次は `bash` / `zsh`、`add` のパス引数はファイル名です。stdout が TTY のときは、rc に貼る用の `eval "$(cdd init …)"` 一行をクリップボードへコピーするか Y/n で確認します。

```bash
$ eval "$(cdd init zsh)"
$ cdd init zsh
```

#### `list`

登録されたエイリアスを一覧表示します。初回は `home` と `docs` が登録されています。

```bash
$ cdd list
    0 | home | ${HOME}
    1 | docs | ${HOME}/Documents
```

#### `add`

`$ cdd add <name> <path>` でエイリアスを登録します。`<path>` には絶対パス・相対パスのどちらも使えます。

```bash
$ cdd add dls /Users/me/Downloads
$ cdd add dls .                          # カレントディレクトリ
$ cdd add docs '${HOME}/Documents'       # ${HOME} はクォートしてシェル展開を防ぐ
```

#### `remove` / `rm`

`$ cdd remove <name>` で削除します。`list` の番号でも指定できます。

```bash
$ cdd remove dls
```

#### `up` / `down`

`list` 上の並び順を変更します。省略時は 1 行、第 2 引数で移動量を指定できます。

```bash
$ cdd up dls
$ cdd up dls 2
$ cdd down dls 2
```

#### `edit`

設定ファイルを既定のエディタで開きます（現状 Windows 未対応）。

```bash
$ cdd edit
```

## 設定ファイル

エイリアスは `${HOME}/.cdd.yaml` に保存されます。直接編集することもできます。

```yaml
- name: home
  dir: ${HOME}
- name: docs
  dir: ${HOME}/Documents
```

`dir` には絶対パスを指定します。例外として `${HOME}` からの相対パスも使えます。その他の環境変数は展開されません。

![cdd](./logo.jpg)

# cdd

[![test](https://github.com/koooyooo/cdd/actions/workflows/test.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/test.yaml)
[![lint](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml)

- `cdd` は事前登録したディレクトリにジャンプ可能な `cd`です
- 作業ディレクトリから遠く離れたディレクトリへの遷移が可能です
- ただし遷移対象のディレクトリは事前の登録が必要です

#### Before
```bash
$ cd ../../Documents/projects/cdd/chart
```

#### After
```bash
$ cdd chart
```

## Install
```bash
$ go install github.com/koooyooo/cdd@latest
```

## Usage
`$ cdd {command}` の形式で各種コマンドを実行します。

### Sub Commands
#### `(alias-name)`
`cdd` コマンドに Aliasの名前を渡すことにより、対象のディレクトリにジャンプできます。
```bash
$ cdd docs

$ pwd
/Users/me/Documents
```

`list` コマンドでリストアップされた際の番号を入力しても、同等の挙動になります。
```bash
$ cdd 1

$ pwd
/Users/me/Documents
```
#### `list`
登録された Aliasをリストアップします。
- デフォルトで 2つのAliasが登録されています。
```bash
$ cdd list
    0 | home | ${HOME}
    1 | docs | ${HOME}/Documents
```

#### `add`
新規に Aliasを登録します。
フォーマットは `$ cdd add ${name} ${absolute-path}` 形式です。
```bash
$ cdd add dls "/Users/me/Downloads"
```

> Note: 
> - `${absolute-path}` 部分はスペースを含まなければ " " で囲む必要はありません
> ```bash
> $ cdd add dls /Users/me/Downloads
> ```
> - `${absolute-path}` 部分に${HOME}を指定する場合は ' 'で囲みシェル展開を防ぎます
> ```bash
> $ cdd add docs '${HOME}/Documents'
> ```

#### `remove` `rm`
既存の Aliasを削除します。
フォーマットは `$ cdd remove ${name}` 形式です。
```bash
$ cdd remove dls
```
> Note: `{name}` 部分は `list` コマンドで表示される番号でも指定可能)

#### `up`
`list` 表示における指定 Aliasの順序を引き上げます。
フォーマットは `$ cdd up ${name}` 形式です。
```bash
$ cdd up dls

# 2行分 up
$ cdd up dls 2
```

#### `down`
`list` 表示における指定 Aliasの順序を引き下げます。
フォーマットは `$ cdd down ${name}` 形式です。
```bash
$ cdd down dls

# 2行分 down
$ cdd down dls 2
```

#### `edit`
設定ファイルを既定のエディタで開きます。
後述の `config` ファイルを編集する際のショートカットとして利用できます。
```bash
$ cdd edit
```

## config
`add`コマンド等で設定した Aliasは `${HOME}/.cdd.yaml` ファイルに保存されます。これを直接変更して Aliasの一覧を編集することも可能です。
- `edit` コマンドで 編集することも可能です。(現状 win環境未対応)
```yaml
- name: home
  dir: ${HOME}
- name: docs
  dir: ${HOME}/Documents
```

> Note: `dir` 部分には 絶対パスを指定しますが、例外的に `${HOME}`からの相対パスを指定することも可能です。その他の環境変数は読み込みません。

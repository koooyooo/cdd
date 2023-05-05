![cdd](./logo.jpg)

# cdd

[![test](https://github.com/koooyooo/cdd/actions/workflows/test.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/test.yaml)
[![lint](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml)

- `cdd` は事前登録したディレクトリにジャンプ可能な `cd`です。
- ディレクトリに直接 `cd`することが可能です。

## Install
```bash
$ go install github.com/koooyooo/cdd@latest
```

## Usage
`$ cdd {command}` の形式で各種コマンドを実行します


### `list`
登録された Aliasをリストアップします。
- デフォルトで 2つのAliasが登録されています。
```bash
$ cdd list
    0 | home | ${HOME}
    1 | docs | ${HOME}/Documents
```

### `add`
新規に Aliasを登録します。
フォーマットは `$ cdd add ${name} ${absolute-path}` 形式です。
```bash
$ cdd add dls "/Users/me/Downloads"
```

### `(alias-name)`
`cdd` コマンドに Aliasの名前を渡すことにより、対象のディレクトリにジャンプできます。
```bash
$ cdd docs

[me@mac]% pwd
/Users/me/Documents
```

`list` コマンドでリストアップされた際の番号を入力しても、同等の挙動になります。
```bash
$ cdd 0

[me@mac]% pwd
/Users/me/Documents
```

## config
`add`コマンド等で設定した Aliasは `${HOME}/.cdd.yaml` ファイルに保存されます。これを直接変更して Aliasの一覧を編集することも可能です。
```yaml
- name: home
  dir: ${HOME}
- name: docs
  dir: ${HOME}/Documents
```

> Note: `dir` 部分には 絶対パスを指定しますが、例外的に `${HOME}`からの相対パスを指定することも可能です。その他の環境変数は読み込みません。

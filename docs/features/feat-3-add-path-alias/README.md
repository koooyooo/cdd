# feat-3: `cdd add` の path のみ登録と `--alias`

## 概要

`cdd add <path>` だけで、パス末尾のディレクトリ名を alias として登録できるようにする。明示指定用に `--alias` / `-a` を追加する。既存の `cdd add <name> <path>` は維持する。

## 要件

- `cdd add <path>` → 解決後パスの `filepath.Base` を alias にする
- `cdd add <name> <path>` → 現行どおり
- `cdd add <path> --alias <name>` / `-a <name>` → path のみ＋明示 alias
- 位置引数の `<name>` と `--alias` を同時指定したらエラー
- `.` / `..` / 相対パスは絶対パス化してから basename を取る。basename が `.` / `..` / 空ならエラー
- `${HOME}` / `~` 付きパスは `Replace4Get` 相当で展開してから basename を取る
- シェル補完: `add` / `a` の直後（第2語）と、その次（第3語）をパス補完する
- README（EN/JA）の add 説明を更新する

## 設計

- `cmd/add.go` に `--alias` / `-a` フラグを追加
- 引数解決はテスト可能な `resolveAddArgs(args, aliasFlag)` に切り出す
- パス由来 alias は `aliasFromPath(path)`（展開 → `filepath.Abs` → `Base`）
- パスの保存は従来どおり `common.Replace4Store`
- 補完テンプレートは `pathArgWords` の対象を COMP_CWORD 2 と 3（zsh は CURRENT 3 と 4）にする

# feat-2: `up` / `down` を `move-up` / `move-down` に改名

## 概要

一覧上の並び替えコマンド名を、ディレクトリ移動と誤解しにくい `move-up` / `move-down` に変更する。

## 要件

- 正式名は `move-up` / `move-down`
- 既存の `up` / `down` は互換エイリアスとして残す
- シェル補完（第1引数の候補、エイリアス引数の対象）に新名・旧名の両方を含める
- README / AGENTS の記載を更新する

## 設計

- Cobra の `Use` を `move-up` / `move-down` にし、`Aliases: []string{"up"}` / `{"down"}` を付ける
- `cmd/shell_integration.go` の `completionCommandWords` と `aliasArgWords` に両系統を載せる

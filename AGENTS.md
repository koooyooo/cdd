# AGENTS.md

AI コーディングエージェント向けのリポジトリガイドです。

## What this is

`cdd` は、事前登録したディレクトリへジャンプする Go CLI です。エイリアスは `${HOME}/.cdd.yaml` に YAML で保存されます。パス解決はバイナリの `print`、実際の移動はシェル関数経由の `builtin cd` です。CLI は [Cobra](https://github.com/spf13/cobra) ベースです。

## Project structure

- `main.go` — エントリポイント。`cmd.Execute()` を呼ぶだけ
- `cmd/` — Cobra サブコマンド（`add` / `list` / `remove` / `move-up` / `move-down` / `edit` / `print` / `init`）。ルートはヘルプのみ。`up` / `down` は互換エイリアス
- `model/` — `Alias`（`name` / `dir`）とパス展開
- `repo/` — `.cdd.yaml` の読み書きとインメモリキャッシュ（シングルトン）
- `common/` — パス解決・`${HOME}` 置換・ファイル存在確認
- `common/constant/` — 設定ファイル名（`.cdd.yaml`）

## Build, test, lint

```bash
go test ./...                 # ユニットテスト（CI と同じ）
go build -o cdd main.go       # ローカルビルド
make install                  # ビルドして ${GOPATH}/bin へ配置
make install-gh               # go install github.com/koooyooo/cdd@latest
golangci-lint run             # .golangci.yml 準拠（gocyclo / staticcheck / govet / revive など）
```

Go バージョンは `go 1.20` 以上（CI も `>=1.20.0`）。

## Architecture notes

- **設定の永続化**: `repo` が `${HOME}/.cdd.yaml` を読み書きする。初回未存在時は `home` / `docs` を初期登録する
- **パス置換**: 保存時は `common.Replace4Store`、取得時は `common.Replace4Get`。`${HOME}` のみ特別扱い
- **ジャンプ**: `cdd print {name|num}` が絶対パスを stdout に出す。シェル連携（`eval "$(cdd init zsh)"` 等）がそれを受けて `builtin cd` する。入れ子シェルは使わない
- **コマンド追加**: `cmd/*.go` に Cobra コマンドを定義し、`init()` で `rootCmd.AddCommand(...)` する既存パターンに合わせる

## Coding style

- `gofmt` 済みの idiomatic Go。パッケージは小文字、公開 API は CamelCase
- エラーは呼び出し側で扱い、CLI 境界では `log.Fatal` / メッセージ表示が既存パターン
- 依存は必要最小限（cobra / yaml.v3 / tablewriter / testify）
- コメントやユーザー向けメッセージは既存どおり日本語・英語混在でよい。新規は簡潔に

## Testing

- テストファイルは `*_test.go`。`model/` / `common/` / `cmd/`（`findPathWithRepo`）など
- ビジネスロジック（パス置換、エイリアス解決）を触る変更ではテストを追加・更新する
- 実行は `go test ./...`

## Commits & PRs

- 短い命令形。`fix:` / `feat:` / `chore:` などの scoped prefix 可。日本語も可
- PR は目的・変更点・実行したテスト（`go test ./...`、必要なら lint）を書く
- 秘密情報やローカルの `.cdd.yaml` 実体はコミットしない

## Do not

- 設定ファイルパスをハードコードで散在させない（`common` / `constant` 経由）
- `repo` のキャッシュと YAML の不整合を残す変更をしない（書き込み後はキャッシュも更新）
- 親シェルの cwd を変えるために子シェルを起動しない（`print` + シェル関数を使う）
- Windows 非対応機能（例: `edit`）を壊す・偽って対応済みにしない

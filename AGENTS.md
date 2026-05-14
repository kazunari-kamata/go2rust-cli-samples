# AGENTS.md

このリポジトリは `go2rust-cli` の変換確認に使う Go サンプル集です。

## 方針

- 変換器本体は `go2rust-cli` リポジトリで管理する。
- このリポジトリには、変換入力として使いやすい小さな Go ファイルを置く。
- サンプルは完全なアプリケーションよりも、変換対象構文が分かる最小コードを優先する。
- 未対応構文を確認するためのサンプルも残す。
- ドキュメントやコメントにローカル環境固有のパス、ユーザー名、作業ディレクトリを含めない。

## バージョン管理

- ローカル作業は `jj` を優先する。
- `jj status`、`jj log`、`jj diff`、`jj bookmark`、`jj git push` を使う。
- author はこのリポジトリ配下の設定で `Kazunari Kamata <14287197+kazunari-kamata@users.noreply.github.com>` にする。

## 検証

Go toolchain がある環境では次を実行する。

```sh
go test ./...
```

`go2rust-cli` が隣接ディレクトリにある環境では、次のように変換結果を確認できる。

```sh
cargo run --manifest-path ../go2rust-cli/Cargo.toml -- convert -i samples/basic/hello/main.go
```

# go2rust-cli-samples

`go2rust-cli` の変換入力として使う Go サンプル集です。

`go2rust-cli` は Go ソースコードを Rust ソースコードの雛形へ変換する CLI です。このリポジトリでは、対応済み構文と未対応構文の出力を確認しやすい小さな Go ファイルを管理します。

## 使い方

隣接ディレクトリに `go2rust-cli` がある場合は、次のように変換できます。

```sh
cargo run --manifest-path ../go2rust-cli/Cargo.toml -- convert -i samples/basic/hello/main.go
```

出力ファイルを指定する例です。

```sh
cargo run --manifest-path ../go2rust-cli/Cargo.toml -- convert \
  -i samples/basic/variables/main.go \
  -o output/variables.rs
```

変換できるかだけを確認する例です。

```sh
cargo run --manifest-path ../go2rust-cli/Cargo.toml -- convert \
  -i samples/basic/hello/main.go \
  --check
```

## サンプル

- `samples/basic/hello/main.go`: `package main`、`import "fmt"`、`func main()`、`fmt.Println`
- `samples/basic/variables/main.go`: `string` と `int` の変数宣言
- `samples/basic/return_value/main.go`: `return x` と、未対応の通常関数宣言
- `samples/unsupported/control_flow/main.go`: 未対応の `if`、`for` を TODO 出力として確認するサンプル
- `samples/unsupported/functions/main.go`: 引数や戻り値を持つ関数が未対応として残ることを確認するサンプル

## 検証

Go toolchain がある環境では、サンプルが Go コードとして壊れていないことを確認できます。

```sh
go test ./...
```

## 関連リポジトリ

- https://github.com/kazunari-kamata/go2rust-cli

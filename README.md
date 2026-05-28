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
- `samples/basic/variables/main.go`: 初期値付き変数宣言、`bool`、`float64`、代入
- `samples/basic/return_value/main.go`: 戻り値を持つ通常関数と `return x`
- `samples/basic/functions/main.go`: 引数、戻り値、短変数宣言を持つ通常関数
- `samples/basic/control_flow/main.go`: 単純な `if` / `else if` / `else` ブロックと `fmt.Print`
- `samples/basic/early_return/main.go`: 値なし `return` を含む早期リターン
- `samples/basic/loops/main.go`: 条件付き `for` と無限 `for` の変換サンプル
- `samples/unsupported/control_flow/main.go`: 未対応の `range` を TODO 出力として確認するサンプル
- `samples/unsupported/functions/main.go`: 構造体とメソッドが未対応として残ることを確認するサンプル

## 検証

Go toolchain がある環境では、サンプルが Go コードとして壊れていないことを確認できます。

```sh
go test ./...
```

## 関連リポジトリ

- https://github.com/kazunari-kamata/go2rust-cli

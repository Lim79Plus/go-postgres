### Go(Echo) and Postgres Sample App with Clean Architecture

## reference

- [Dokcer ComposeでGoとPostgreSQLの開発環境構築](https://qiita.com/muroya2355/items/d48c384a4a82c7ed34ae)

- コンテナ内で「go test」が実行できなかった際に参照
    - [GO1.11.9のコンテナでgo testしようとするとexec: "gcc": executable file not found in $PATH と怒られる](https://qiita.com/trewanek/items/579e0065fd203e22f7cd)

- Goで環境変数のテスト実施したが、うまく行かなかった。.bash_profileに記述して再起動してOKに。

## Go applicaation

- [Go envconfig Github](https://github.com/kelseyhightower/envconfig)

## Go JSON

- [Go言語でJSONを扱う](https://qiita.com/nayuneko/items/2ec20ba69804e8bf7ca3)

## Go MOCK

- [Go Mockでインタフェースのモックを作ってテストする #golang](https://qiita.com/tenntenn/items/24fc34ec0c31f6474e6d)

```:
mockgen -source interface/controllers/context.go -destination mock/mock_context.go
```

## Test

- [Go言語のHTTPサーバのテスト事始め](https://qiita.com/theoden9014/items/ac8763381758148e8ce5)
    - apiのテスト
- [package assert GoDoc](https://godoc.org/github.com/stretchr/testify/assert)
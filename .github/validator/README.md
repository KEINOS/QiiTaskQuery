# JSON Schema Validator

このディレクトリは、JSON データのバリデーション用ツールのソースコードです。（Go 言語）

このツールは JSON データが JSON スキーマに準拠しているかを確認します。

- 第 1 引数: JSON スキーマの定義ファイルのパス
- 第 2 引数: 検証したい JSON ファイルのパス

## ツールのビルド

このツールは Go 言語で書かれていますが、`docker`/`docker-compose` がローカルにインストールされている場合は Go は不要です。

その場合、リポジトリのルートで以下を実行すると、JSON スキーマに準拠しているか検証します。

```bash
docker-compose run validator <JSONスキーマ定義ファイルのパス> <検証したいJSONファイル>
```

以下は `query.schema.json` で `query.json` を検証する例です。

```shellsession
$ docker-compose run validator ./query.schema.json ./query.json



### Go でビルド&インストール

Go がインストールされている場合、このディレクトリに移動し以下を実行すると `validator` コマンドが使えるようになります。

```bash
go install .
```

## ツールのデバッグ on VS Code

VS Code 上で、このツールをデバッグする場合は、リポジトリのルートに `go.mod` がないため `gopls` による恩恵が受けられません。

その場合、このディレクトリにある `workspace.code-workspace` をクリックして、右下に表示される「ワークスペースで開く」ボタンを押してください。このディレクトリをルートにした別ワークスペースが開きます。

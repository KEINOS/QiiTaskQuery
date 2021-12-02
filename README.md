# QiiTask Query

QiiTask コマンドに同梱される質問一覧専用のリポジトリです。追加したい質問があれば遠慮なく PR してください。

- 質問集ファイル: [query.json](query.json)
    - HTTPS ダウンロード: [https://git.io/JMPFq](https://git.io/JMPFq)
    - IPFS ダウンロード: [リリースページの CID 参照](https://github.com/KEINOS/QiiTaskQuery/releases/latest)
- `query.json` のスキーマ定義: [query.schema.json](query.schama.json)
    - ダウンロード: [https://git.io/JMPF7](https://git.io/JMPF7)

## コントリビュート

- PR 先: `main` ブランチ
- イタズラの類いでなければ、[CI](./github/workflows/) さえ通ればマージ致します。
- `docker-compose` がインストールされていればローカルでテストできます。
    ```bash
    docker-compose run validator ./query.schema.json ./query.json
    ```

## ライセンス

- MIT
- Copyright (c) [The QiiTask Query Contributors](https://github.com/KEINOS/QiiTaskQuery/graphs/contributors).

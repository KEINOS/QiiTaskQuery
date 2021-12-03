<!-- Code generated using ./.github/gen_readme.sh; DO NOT EDIT. -->
[![IsValidJSON](https://github.com/KEINOS/QiiTaskQuery/actions/workflows/validate_json.yaml/badge.svg)](https://github.com/KEINOS/QiiTaskQuery/actions/workflows/validate_json.yaml)
[![](https://img.shields.io/badge/IPFS-QmWBf2hYCzFWtQp5w5NLD3oVSdV6nmdvFAdG4KGgsoBMNh-blue?logo=ipfs)](https://ipfs.io/ipfs/QmWBf2hYCzFWtQp5w5NLD3oVSdV6nmdvFAdG4KGgsoBMNh "IPFS Gateway")

# QiiTask Query

QiiTask コマンドに同梱される質問一覧専用のリポジトリです。追加したい質問があれば遠慮なく PR してください。

- 質問一覧ファイル
    - ソース: [query.json](query.json)（ダウンロード: [https://git.io/JMPFq](https://git.io/JMPFq)）
    - スキーマ定義: [query.schema.json](query.schema.json)（ダウンロード: [https://git.io/JMPF7](https://git.io/JMPF7)）
    - IPFS CID: `QmWBf2hYCzFWtQp5w5NLD3oVSdV6nmdvFAdG4KGgsoBMNh`（ダウンロード: [IPFS Gateway](https://ipfs.io/ipfs/QmWBf2hYCzFWtQp5w5NLD3oVSdV6nmdvFAdG4KGgsoBMNh)）

## コントリビュート

- PR 先: `main` ブランチ
- イタズラの類いでなければ、[CI](./github/workflows/) さえ通ればマージ致します。
- `docker-compose` がインストールされていれば、ローカルでテストできます。

    ```bash
    docker-compose run validator ./query.schema.json ./query.json
    ```

## ライセンス

- MIT
- Copyright (c) [The QiiTask Query Contributors](https://github.com/KEINOS/QiiTaskQuery/graphs/contributors).

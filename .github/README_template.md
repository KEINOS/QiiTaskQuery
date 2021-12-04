[![IsValidJSON](https://github.com/KEINOS/QiiTaskQuery/actions/workflows/validate_json.yaml/badge.svg)](https://github.com/KEINOS/QiiTaskQuery/actions/workflows/validate_json.yaml)
[![](https://img.shields.io/badge/IPFS-%%IPFS_CID%%-blue?logo=ipfs)](https://ipfs.io/ipfs/%%IPFS_CID%% "IPFS Gateway")

# QiiTask Query

QiiTask コマンドに同梱される質問一覧専用のリポジトリです。追加したい質問があれば遠慮なく PR してください。

- 質問一覧ファイル
    - ソース: [query.json](query.json)（ダウンロード: [https://git.io/JMPFq](https://git.io/JMPFq)）
    - スキーマ定義: [query.schema.json](query.schema.json)（ダウンロード: [https://git.io/JMPF7](https://git.io/JMPF7)）
    - IPFS CID: `%%IPFS_CID%%`（ダウンロード: [IPFS Gateway](https://ipfs.io/ipfs/%%IPFS_CID%%)）

## コントリビュート

- PR 先: `main` ブランチ
- マージ:
    - イタズラの類いでなければ、[CI](./github/workflows/) さえ通ればマージ致します。
    - マージされても、別 PR で修正される可能性があります。
- 自動マージ:
    - `query.json` のみの変更で、CI をパスした PR は自動的にマージされます。
    - 複数ファイルの変更は、2 名以上の[コントリビュータ](https://github.com/KEINOS/QiiTaskQuery/graphs/contributors)の `approved` が付いた PR は自動的にマージされます。
- テスト:
    - `docker-compose` がインストールされていれば、事前にローカルでテストできます。

    ```bash
    docker-compose run validator ./query.schema.json ./query.json
    ```

## ライセンス

- MIT
- Copyright (c) [The QiiTask Query Contributors](https://github.com/KEINOS/QiiTaskQuery/graphs/contributors).

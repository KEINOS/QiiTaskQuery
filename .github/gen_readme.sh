#!/bin/sh
# =============================================================================
#  README Generator / IPFS CID Updator
# =============================================================================
#  このシェル・スクリプトは、トップページの README を更新します。query.json の
#  IFPS CID を現在のものに置き換えます。このスクリプトは ../docker-compose.yml
#  から呼び出されることを前提としています。

NAME_FILE_TARGET="query.json"
PATH_FILE_TEMPLATE="./.github/README_template.md"
PATH_FILE_README="./README.md"

ipfs init 2>/dev/null 1>/dev/null

HASH_CID="$(ipfs add --offline --only-hash --quieter "${NAME_FILE_TARGET}")"

sed "s/%%IPFS_CID%%/${HASH_CID}/g" <"$PATH_FILE_TEMPLATE" >"$PATH_FILE_README"

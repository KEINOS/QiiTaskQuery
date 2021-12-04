#!/bin/sh
# =============================================================================
#  README Generator / IPFS CID Updator
# =============================================================================
#  This script must be called from ../docker-compose.yml
#  このスクリプトは ../docker-compose.yml から呼び出されることを前提としています。
#
#  This script will update the ../README.md with IPFS CID of current query.json.
#  このシェル・スクリプトは、トップページの README を更新します。query.json の
#  CID (IFPS CID) を現在のものに置き換えます。

NAME_FILE_TARGET="query.json"
PATH_FILE_TEMPLATE="./.github/README_template.md"
PATH_FILE_README="./README.md"

# Get IPFS CID
ipfs init 2>/dev/null 1>/dev/null
HASH_CID="$(ipfs add --offline --only-hash --quieter "${NAME_FILE_TARGET}")"

# Add header
echo '<!-- Code generated using ./.github/gen_readme.sh on release; DO NOT EDIT. -->' >"$PATH_FILE_README"

# Append template contents
sed "s/%%IPFS_CID%%/${HASH_CID}/g" <"$PATH_FILE_TEMPLATE" >>"$PATH_FILE_README"

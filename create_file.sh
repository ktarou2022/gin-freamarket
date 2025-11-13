#!/bin/bash

# スクリプトの使用方法を表示する関数
usage() {
    echo "Usage: $0 "
    echo "Example: $0 item"
    exit 1
}

# 引数が1つ未満の場合は使い方を表示して終了
if [ "$#" -lt 1 ]; then
    usage
fi

ENTITY_NAME=$1
EXT=".go" # ファイルの拡張子を設定 (例: .go, .cs, .py)
BASE_DIR="./"

# 作成するファイルとディレクトリのリスト
FILES=(
    "controllers/${ENTITY_NAME}_controller${EXT}"
    "services/${ENTITY_NAME}_service${EXT}"
    "repositories/${ENTITY_NAME}_repository${EXT}"
)

echo "Creating files for entity: $ENTITY_NAME"

# ファイルをループで作成
for REL_PATH in "${FILES[@]}"; do
    FILE_PATH="${BASE_DIR}/${REL_PATH}"
    DIR_PATH=$(dirname "$FILE_PATH")

    # ディレクトリが存在しなければ作成
    if [ ! -d "$DIR_PATH" ]; then
        mkdir -p "$DIR_PATH"
        echo "Created directory: $DIR_PATH"
    fi

    # ファイルが既に存在するか確認
    if [ -f "$FILE_PATH" ]; then
        echo "Error: File '$FILE_PATH' already exists. Skipping."
    else
        # 空のファイルを作成
        touch "$FILE_PATH"
        echo "Created file: $FILE_PATH"
    fi
done

echo "Done."

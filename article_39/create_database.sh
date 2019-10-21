#!/bin/sh

DATABASE=test.db

SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"

cat "${SCRIPT_DIR}/schema.sql" | sqlite3 "${SCRIPT_DIR}/${DATABASE}"
cat "${SCRIPT_DIR}/test_data.sql" | sqlite3 "${SCRIPT_DIR}/${DATABASE}"

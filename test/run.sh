#!/bin/bash
set -e

export ALLURE_OUTPUT_PATH="$(pwd)"       
export ALLURE_OUTPUT_FOLDER="allure-results"

# Clean old information
rm -rf "$ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER"

# Tests running
go test -v ./...

# Report generation
allure generate "$ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER" --clean -o "$ALLURE_OUTPUT_PATH/allure-report"

echo "Allure results collected in $ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER"
echo "HTML report generated in $ALLURE_OUTPUT_PATH/allure-report"

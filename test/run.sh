#!/bin/bash
set -e

CLEAN=false
REPORT=false
SERVER=false
while [[ $# -gt 0 ]]; do
    case "$1" in
        -c|--clean)
            CLEAN=true
            shift
            ;;
        -r|--report)
            REPORT=true
            shift
            ;;
        -s|--server)
            SERVER=true
            shift
            ;;
        *)
            echo "Unknown option: $1"
            echo "Usage: $0 [-c|--clean] [-s|--server]"
            exit 1
        ;;
    esac
done

export ALLURE_OUTPUT_PATH="$(pwd)"       
export ALLURE_OUTPUT_FOLDER="allure-results"

# --clean -c
# Clean old information
if [ "$CLEAN" = true ]; then
    rm -rf "$ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER"
    echo "Old results have been cleaned."
fi

# Tests running
go test -v ./...
echo "Allure results collected in $ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER."

# --report -r
# Report generation
if [ "$REPORT" = true ]; then
    allure generate "$ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER" --clean -o "$ALLURE_OUTPUT_PATH/allure-report"
    echo "HTML report generated in $ALLURE_OUTPUT_PATH/allure-report."
fi

# --server -s
# Starting Allure host
if [ "$SERVER" = true ]; then
    echo "Starting Allure server..."
    allure serve "$ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER"
fi

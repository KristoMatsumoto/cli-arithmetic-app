#!/bin/bash
set -e

CLEAN=false
REPORT=false
SERVER=false
TIMEOUT_FLAG=false
TIMEOUT_VALUE=""
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
        -t|--timeout)
            TIMEOUT_FLAG=true
            if [[ -n "$2" && ! "$2" =~ ^- ]]; then
                TIMEOUT_VALUE="$2"
                shift 2
            else
                TIMEOUT_VALUE="30s"
                shift
            fi
            ;;
        *)
            echo "Unknown option: $1"
            echo "Usage: $0 [-c|--clean] [-r|--report] [-s|--server] [-t|--timeout [value]]"
            exit 1
        ;;
    esac
done

export ALLURE_OUTPUT_PATH="$(pwd)"       
export ALLURE_OUTPUT_FOLDER="allure-results"
export SECRET_KEY="example123456789"

# --clean -c
# Clean old information
if [ "$CLEAN" = true ]; then
    rm -rf "$ALLURE_OUTPUT_PATH/$ALLURE_OUTPUT_FOLDER"
    echo "Old results have been cleaned."
fi

# --timeout [x] -t [x]
# Add timeout for test running (default: x = 30s)
GO_TEST_CMD="go test -v"
if [ "$TIMEOUT_FLAG" = true ]; then
    GO_TEST_CMD+=" -timeout $TIMEOUT_VALUE"
fi
GO_TEST_CMD+=" ./..."

# Tests running
eval $GO_TEST_CMD
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

# ---- VARIABLES ----
ALLURE_OUTPUT_PATH ?= $(CURDIR)
ALLURE_OUTPUT_FOLDER ?= allure-results
ALLURE_REPORT_FOLDER ?= allure-report

# ALLURE_LOCALE ?= en

RESULTS_PATH := $(ALLURE_OUTPUT_PATH)/$(ALLURE_OUTPUT_FOLDER)
REPORT_PATH := $(ALLURE_OUTPUT_PATH)/$(ALLURE_REPORT_FOLDER)

TEST_ENV_FILE ?= test.env

# ---- TARGETS ----

.PHONY: env help test test-app test-cli test-rest-api test-soap-api clean report serve open allure-tests

env:
	envsubst < $(TEST_ENV_FILE)

help:
	@printf "\n\033[1;36mAvailable Make targets:\033[0m\n\n"
	@printf "  \033[1;33mmake test\033[0m           	Run Go tests for all modules\n"
	@printf "  \033[1;33mmake test-app\033[0m       	Run Go tests for APP\n"
	@printf "  \033[1;33mmake test-cli\033[0m       	Run Go tests for CLI\n"
	@printf "  \033[1;33mmake test-rest-api\033[0m  	Run Go tests for REST API\n"
	@printf "  \033[1;33mmake test-soap-api\033[0m  	Run Go tests for SOAP API\n"
	@printf "  \033[1;33mmake test-allure\033[0m    	Run Go tests with clean and Allure report\n"
	@printf "  \033[1;33mmake report\033[0m         	Generate Allure report\n"
	@printf "  \033[1;33mmake serve\033[0m          	Start Allure server\n"
	@printf "  \033[1;33mmake clean\033[0m          	Clean allure-results and allure-report\n"
	@printf "  \033[1;33mmake open\033[0m           	Open existing Allure report\n\n"

## allure-tests — full cycle
allure-tests: clean test report open

## test — run all tests
test: env test-app test-cli test-rest-api test-soap-api
	@echo All tests have been completed.

## test-app — run app tests
test-app: env
	@echo Run tests for APP...
	ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH) \
	ALLURE_OUTPUT_FOLDER=$(ALLURE_OUTPUT_FOLDER) \
	go test -v -count=1 ./app/test/... || TEST_FAILED=$$?; \
	echo "Tests finished with code $$TEST_FAILED"; \
	exit 0
	@echo Tests for APP have been completed.

## test-cli — run cli tests
test-cli: env
	@echo Run tests for CLI...
	ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH) \
	ALLURE_OUTPUT_FOLDER=$(ALLURE_OUTPUT_FOLDER) \
	go test -v -count=1 ./cli/test/... || TEST_FAILED=$$?; \
	echo "Tests finished with code $$TEST_FAILED"; \
	exit 0
	@echo Tests for CLI have been completed.

## test-rest-api — run rest-api tests
test-rest-api: env
	@echo Run tests for REST-API...
	ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH) \
	ALLURE_OUTPUT_FOLDER=$(ALLURE_OUTPUT_FOLDER) \
	go test -v -count=1 ./rest-api/test/... || TEST_FAILED=$$?; \
	echo "Tests finished with code $$TEST_FAILED"; \
	exit 0
	@echo Tests for REST API have been completed.

## test-soap-api — run soap-api tests
test-soap-api: env
	@echo Run tests for SOAP-API...
	ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH) \
	ALLURE_OUTPUT_FOLDER=$(ALLURE_OUTPUT_FOLDER) \
	go test -v -count=1 ./soap-api/test/... || TEST_FAILED=$$?; \
	echo "Tests finished with code $$TEST_FAILED"; \
	exit 0
	@echo Tests for SOAP API have been completed.

## clean — clean old information
clean:
	@echo Cleaning Allure data...
	rm -rf $(ALLURE_OUTPUT_PATH)/$(ALLURE_OUTPUT_FOLDER) $(ALLURE_OUTPUT_PATH)/$(ALLURE_REPORT_FOLDER)
	@echo Old results have been cleaned.

## report — generate HTML-report
report:
	@echo Generation Allure report...
	allure generate $(ALLURE_OUTPUT_PATH)/$(ALLURE_OUTPUT_FOLDER) --clean -o $(ALLURE_OUTPUT_PATH)/$(ALLURE_REPORT_FOLDER)

## serve — starting Allure server
serve:
	@echo Starting Allure server...
	allure serve $(ALLURE_OUTPUT_PATH)/$(ALLURE_OUTPUT_FOLDER)

open:
	allure open 


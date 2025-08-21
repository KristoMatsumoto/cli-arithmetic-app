$env:ALLURE_OUTPUT_PATH = (Get-Location).Path
$env:ALLURE_OUTPUT_FOLDER = "allure-results"

# Clean old information
Remove-Item -Recurse -Force "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER" -ErrorAction SilentlyContinue

# Tests running
go test -v ./...

# Report generation
allure generate "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER" --clean -o "$env:ALLURE_OUTPUT_PATH\allure-report"

Write-Host "Allure results collected in $env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER"
Write-Host "HTML report generated in $env:ALLURE_OUTPUT_PATH\allure-report"

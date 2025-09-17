param (
    [switch]$clean,
    [switch]$report,
    [switch]$server,
    [Parameter(Mandatory=$false)] [string]$timeout
)

$env:ALLURE_OUTPUT_PATH = (Get-Location).Path
$env:ALLURE_OUTPUT_FOLDER = "allure-results"

# -clean 
# Clean old information
if ($clean) {    
    Remove-Item -Recurse -Force "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER" -ErrorAction SilentlyContinue
    Write-Host "Old results have been cleaned."
}

# -timeout [x]
# Add timeout for test running (default: x = 30s)
$goTestCommand = "go test -v"
if ($PSBoundParameters.ContainsKey("timeout")) {
    if ([string]::IsNullOrEmpty($timeout)) {
        $goTestCommand += " -timeout 30s"
    } else {
        $goTestCommand += " -timeout $timeout"
    }
}
$goTestCommand += " ./..."

# Tests running
Invoke-Expression $goTestCommand 
Write-Host "Allure results collected in $env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER."

# -report
# Report generation
if ($report) { 
    allure generate "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER" --clean -o "$env:ALLURE_OUTPUT_PATH\allure-report"
    Write-Host "HTML report generated in $env:ALLURE_OUTPUT_PATH\allure-report."
}

# -server
# Starting Allure host
if ($server) {
    Write-Host "Starting Allure server..."
    allure serve "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER"
}

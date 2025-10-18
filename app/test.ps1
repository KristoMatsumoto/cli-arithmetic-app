param (
    [switch]$clean,
    [switch]$report,
    [switch]$server,
    [Parameter(Mandatory=$false)] [string]$timeout
)

# Read environment variables
if (Test-Path "test.env") {
    Get-Content "test.env" | ForEach-Object {
        if ($_ -match "^(.*?)=(.*)$") {
            $name = $matches[1]
            $value = $matches[2]
            [System.Environment]::SetEnvironmentVariable($name, $value, "Process")
        }
    }
}

# Properties of allure files path
# ALLURE_OUTPUT_PATH not specified, ALLURE_OUTPUT_RELATIVE_PATH not specified    
#       - called in the current directory
# ALLURE_OUTPUT_PATH not specified, ALLURE_OUTPUT_RELATIVE_PATH specified        
#       - absolute path is {current_path}/{ALLURE_OUTPUT_RELATIVE_PATH}
# ALLURE_OUTPUT_PATH specified                                                   
#       - ignore ALLURE_OUTPUT_RELATIVE_PATH and as result absolute path is as ALLURE_OUTPUT_PATH
if (-not $env:ALLURE_OUTPUT_PATH) {
    if ($env:ALLURE_OUTPUT_RELATIVE_PATH) {
        $env:ALLURE_OUTPUT_PATH = Join-Path (Get-Location).Path $env:ALLURE_OUTPUT_RELATIVE_PATH
    } else {
        $env:ALLURE_OUTPUT_PATH = (Get-Location).Path
    }
}
# ...and report path
if (-not $env:ALLURE_REPORT_FOLDER) {
    $env:ALLURE_REPORT_FOLDER = "allure-report"
}

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
$goTestCommand += " ./test/..."

# Tests running
Invoke-Expression $goTestCommand 
Write-Host "Allure results collected in $env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER."

# Locale (default EN)
if (-not $env:ALLURE_LOCALE) {
    $env:ALLURE_LOCALE = "en"
}

# -report
# Report generation
if ($report) { 
    allure generate "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER" --clean -o "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_REPORT_FOLDER" --locale $env:ALLURE_LOCALE
    Write-Host "HTML report generated in $env:ALLURE_OUTPUT_PATH\$env:ALLURE_REPORT_FOLDER."
}

# -server
# Starting Allure host
if ($server) {
    Write-Host "Starting Allure server..."
    allure serve "$env:ALLURE_OUTPUT_PATH\$env:ALLURE_OUTPUT_FOLDER" --locale $env:ALLURE_LOCALE
}

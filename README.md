# CLI-ARITHMETIC-APP

Test app to try Go-langruage skills.
At the moment, a small console application for processing arithmetic operations in the file.

---

## Usage

Start with

```go run main.go [OPTIONS]```

Options:

`--input=[VALUE] || -i [VALUE]` Indicating the path to the input file (default: "test/inputs/input.txt")

`--output=[VALUE] || -o [VALUE]` Indicating the name to create the output file (default: "test/outputs/output")

`--format=[VALUE] || -f [VALUE]` The processed format (default: "txt")

`--processor-version=[VALUE] || -p [VALUE]` Processor version (1 for naive, 2 for regex processor) (default: "1")

`--version=[VALUE] || -v [VALUE]` The logic version (only "1" now)

---

## Documentation

In progress...

---

## Testing

To test project we use allure ([ozontech/allure-go](https://github.com/ozontech/allure-go)).

To run test and create html-report page use:

- for Bash / Linux / macOS: **./run.sh** (from /test)

- for Windows: **./run.ps1** (from /test)

Options:

```
[ Windows  |    Bash   |    Bash ]
        option                          description
-clean      --clean         -c      clean old test information
-report     --report        -r      create allure html-report
-server     --server        -s      start allure server

```

---

Kristo Matsumoto

July 2025

In progress...

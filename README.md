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

`--processor-version=[VALUE] || -p [VALUE]` Processor version (1 for naive, 2 for regex, 3 for library processor) (default: "1")

`--version=[VALUE] || -v [VALUE]` The logic version (only "1" now)

---

## Documentation (RU)

Временное решение для ознакомления. Это далеко от документации на данный момент, но все мы ведь понимаем, что это учебный проект и никто кроме меня его просматривать для расширения не собирается... да и лишь временно ;)

Общая информация:

- Поддержка стандартных бинарных операторов: +, -, *, /.
- Все процессоры выдают результат с точностью до сотых.
- Все процессоры возвращают NaN в случаях деления на 0 и\или взятия остатка от деления на ноль.

### Naive Processor (99%)

- Поддержка дополнительных операторов: %, ^.
- Поддерживаются унарные плюсы и минусы.
- Поддержка чисел с плавающей точкой.
- IP-адреса или числа с несколькими точками не считаются выражениями.
- Поддержка многократно вложенных скобок.

Баги и недоработки

- Текст не сохраняется в полной мере: могут "съедаться" пробелы после выражения
- Возвращается NaN в случае синтаксической ошибки внутри выражения

### Regex Processor (88.76%)

Используется регулярное выражение для поиска кандидатов на арифметическое выражение.

- Поддержка ополнительных операторов: %.
- Поддержка чисел с плавающей точкой.
- IP-адреса или числа с несколькими точками не считаются выражениями.
- Сохраняется текст, не являющийся выражением, без изменений.
- Поддержка многократно вложенных скобок.
- Возвращается NaN в случае синтаксической ошибки внутри выражения

Баги и недоработки

- Унарные плюсы и минусы пока не поддерживаются в полной мере.
- Возведение в степень не поддерживается в полной мере.

### Library Processor (100%)

Используется ([mna/pigeon](https://github.com/mna/pigeon)) для разбития строк на текстовые и вычисляемые части, regex - для нормализации выражений под библиотеку-вычислитель ([Knetic/govaluate](https://github.com/Knetic/govaluate)).

- Поддержка ополнительных операторов: %, ^.
- Поддерживаются унарные плюсы и минусы.
- Поддержка чисел с плавающей точкой.
- IP-адреса или числа с несколькими точками не считаются выражениями.
- Сохраняется текст, не являющийся выражением, без изменений.
- Поддержка многократно вложенных скобок.

Особенности

- Невалидные выражения не вычисляются, оставаясь текстом, засчет полного контроля написаний выражений через псевдоязык.

P.S. Процент готовности основан на текущих тестовых кейсах, доступных в репозитории. Цифры могут быть неточными, а тестовые случаи дополнены в будущем.

---

## Testing

To test project we use allure ([ozontech/allure-go](https://github.com/ozontech/allure-go)).

To run test and create html-report page use:

- for Bash / Linux / macOS: **./run.sh** (from /test)

- for Windows: **./run.ps1** (from /test)

Options:

```
        option                            description
[ Windows    |   Bash       |  Bash   ]
-clean        --clean         -c        clean old test information
-report       --report        -r        create allure html-report
-server       --server        -s        start allure server
-timeout [x]  --timeout [x]   -t [x]    add timeout option for *go test*

```

---

Kristo Matsumoto

July 2025

In progress...

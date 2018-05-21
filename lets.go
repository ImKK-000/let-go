package main

import (
    "fmt"
)

func main() {
    var number1, number2 int
    var operator uint8
    sum := 0

    fmt.Scanf("%d %c %d", &number1, &operator, &number2)

    if (string(operator) == "+") {
        sum = number1 + number2
    } else if (string(operator) == "-") {
        sum = number1 - number2
    } else if (string(operator) == "*") {
        sum = number1 * number2
    } else if (string(operator) == "/") {
        if (number2 == 0) {
            number2 = 1
        }

        sum = number1 / number2
    }

    fmt.Printf("%d %c %d = %d\n", number1, operator, number2, sum)
}

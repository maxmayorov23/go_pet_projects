package main

import (
	"fmt"
	"regexp"
	"strings"
)

func checkString(s string) (bool, error) {

	if len(s) == 0 {
		return false, fmt.Errorf("%s", "An empty expression has been entered!")
	}

	// Строка содержит только цифры и разрешенные знаки
	valid,_:=regexp.Compile("^[0-9*/()+-.]+$")
	if !valid.MatchString(s) {
		return false, fmt.Errorf("%s","The expression contains invalid characters!")
	}
	// проверка, что есть числа
	hasNumber,_ := regexp.Compile("[0-9]")
	if !hasNumber.MatchString(s) {
		return false, fmt.Errorf("%s","Error! The expression must contain the digits 0-9")
	}
 
	lastWasOperator := true    // проверка нескольких операторов */+-
	brackets := 0     			   // проверка количества скобок
	hasOperator:= false        // проверка, что есть оператор
	lastChar := ' '

	for _,char := range(s) {

		// Проверка открытия и закрытия скобок
		if char == '(' {

			// если уже была закрывающая скобка, то ошибка
			if brackets < 0 {
				return false, fmt.Errorf("%s", "Ошибка! Не правильно расставлены скобки")
			} 

			brackets++

			// Если перед скобкой число, то ошибка. Должен быть оператор
			if (lastChar >= '0' && lastChar <= '9') {
				return false, fmt.Errorf("%s", "Error! There is no operator before the '(' character.")
			}

		} else if char == ')' {
			brackets--
		}
		
		isOperator := strings.ContainsRune("*/+-", char)

		// если символ оператор
		if isOperator {		

			if lastWasOperator {
				return false, fmt.Errorf("%s","Error using operators! The line starts with an operator or two operators in a row.")
			}
			lastWasOperator = true
			hasOperator = true

		} else {
			// последний символ не оператор
			lastWasOperator = false

			// проверка, что после ')' идет оператор
			if lastChar == ')' {
				return false, fmt.Errorf("%s","Error! The ')' character should be followed by the operator")
			}
		}

		lastChar = char
	}

	// если последний символ строки оператор */+-
	if lastWasOperator {
		return false, fmt.Errorf("%s","Error! The line ends with the operator.")
	}

	// нету операторов */+- 
	if !hasOperator {
		return false, fmt.Errorf("%s","Error! The operator is missing.")
	}

	// Если количество открытых и закрытых скобок разное, то ошибка
	if brackets != 0 {
		return false, fmt.Errorf("%s","Error in the use of signs '(' and ')'")
	}

  return true, nil
}

func main () {

	var task, unwanted string
	fmt.Println("Enter a string without spaces to calculate: ")
	fmt.Scanf("%s %s", &task, &unwanted)
	
	if (unwanted != "") {
		fmt.Printf("\n%s\n", "Do not use spaces when entering a task!")
		return
	}

	_, err := checkString(task)
	if err != nil {
		fmt.Printf("\n%s\n", err)
		return
	}
}


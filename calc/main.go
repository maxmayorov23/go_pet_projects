package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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


func calculate(nums []float64, ops []rune) ([]float64, []rune) {
	// Достаём последнее число и оператор
	right := nums[len(nums)-1]
	left := nums[len(nums)-2]
	op := ops[len(ops)-1]

	// Убираем их из стека
	nums = nums[:len(nums)-2]
	ops = ops[:len(ops)-1]

	var result float64
	switch op {
	case '+':
		result = left + right
	case '-':
		result = left - right
	case '*':
		result = left * right
	case '/':
		if right == 0 {
			panic("деление на ноль")
		}
		result = left / right
	}

	// Возвращаем результат в стек чисел
	nums = append(nums, result)
	return nums, ops
}

func precedence(op rune) int {
	if op == '+' || op == '-' {
		return 1
	}
	if op == '*' || op == '/' {
		return 2
	}
	return 0
}

func Calc(expression string) (float64, error) {

	_, err := checkString(expression)
	if err != nil {
		return 0, fmt.Errorf("%s",err) 
	}
    
	nums := []float64{}
	ops := []rune{}
	i := 0

	for i < len(expression) {
		
		ch := rune(expression[i])

		if unicode.IsDigit(ch) {
			j := i
			for j < len(expression) && unicode.IsDigit(rune(expression[j])) {
				j++
			}
			num, _ := strconv.ParseFloat(expression[i:j], 64)

			nums = append(nums, num)  // Добавлено число в стек
			i = j - 1

		} else if ch == '(' { 
		
			ops = append(ops, ch)

		} else if ch == ')' {
		

			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				nums, ops = calculate(nums, ops)
			}
			ops = ops[:len(ops)-1]

		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {

			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				nums, ops = calculate(nums, ops)
			}
			ops = append(ops, ch)
		}
		i++
	}

	// Выполняем оставшиеся операции
	for len(ops) > 0 {
		nums, ops = calculate(nums, ops)
	}

	return nums[0], nil
}
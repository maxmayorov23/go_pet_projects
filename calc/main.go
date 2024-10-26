package main

import (
	"fmt"
	"regexp"
)

func checkString(s string) (bool, error) {

	if len(s) == 0 {
		return false, fmt.Errorf("%s", "An empty expression has been entered!")
	}

	valid,_:=regexp.Compile("^[0-9*/()+-.]+$")
	if !valid.MatchString(s) {
		return false, fmt.Errorf("%s","The expression contains invalid characters!")
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


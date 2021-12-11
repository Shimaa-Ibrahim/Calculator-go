package calculator

func Calculate(num1 float64, operator string, num2 float64) (result float64) {

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2

	}

	return
}

func CheckOperator(operatorInput string) (isOperator bool) {
	isOperator = false
	operators := [4]string{"+", "-", "*", "/"}
	for _, operator := range operators {
		if operatorInput == operator {
			isOperator = true
			break
		}
	}
	return
}

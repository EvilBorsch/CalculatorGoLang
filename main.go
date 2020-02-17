package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string, error) {
	l := len(s)
	if l == 0 {
		return s, "", errors.New("Stack is Empty")
	}
	return s[:l-1], s[l-1], nil
}

func (s stack) CheckTop() (string, error) {
	if len(s) == 0 {
		return "", errors.New("Stack is Empty")
	}
	return s[len(s)-1], nil
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

func isDigit(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func factorize(s string) []string {
	i := 0
	resArr := make([]string, 0)
	number := ""
	for i < len(s) {
		if isDigit(string(s[i])) {
			for i < len(s) && isDigit(string(s[i])) {

				number += string(s[i])
				i++
			}
			resArr = append(resArr, number)
			number = ""

		}
		if i < len(s) {
			resArr = append(resArr, string(s[i]))
		}
		i++
	}
	return resArr
}

func getPriority(ch string) int {
	const lowPriority = 1
	const highPriority = 2
	const openBrackPriority = 500
	const closeBrackPriority = -2
	const errorCode = 404
	switch ch {
	case "+":
		return lowPriority
	case "-":
		return lowPriority
	case "*":
		return highPriority
	case "/":
		return highPriority
	case "(":
		return openBrackPriority
	case ")":
		return closeBrackPriority
	default:
		return errorCode
	}
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func StringToFloat(st string) float64 {
	num, _ := strconv.ParseFloat(strings.TrimSpace(st), 64)
	return num
}

func calcExpr(expr []string) float64 {
	operationStack := make(stack, 0)
	numbersStack := make(stack, 0)
	for _, el := range expr {
		if isDigit(el) {
			numbersStack = numbersStack.Push(el)
		} else {
			if !operationStack.isEmpty() {
				lastEl, err := operationStack.CheckTop()
				if getPriority(el) > getPriority(lastEl) || lastEl == "(" {
					operationStack = operationStack.Push(el)
				} else {
					if el == ")" {
						for lastEl != "(" {
							operationStack, numbersStack = calcElement(operationStack, numbersStack)
							lastEl, _ = operationStack.CheckTop()
						}
						operationStack, _, _ = operationStack.Pop()
					} else {
						for getPriority(el) <= getPriority(lastEl) && lastEl != "(" {
							operationStack, numbersStack = calcElement(operationStack, numbersStack)
							lastEl, err = operationStack.CheckTop()
							if err != nil {
								break
							}
						}
						operationStack = operationStack.Push(el)
					}
				}
			} else {
				operationStack = operationStack.Push(el)
			}
		}
	}
	for !operationStack.isEmpty() {
		operationStack, numbersStack = calcElement(operationStack, numbersStack)
	}
	_, answer, _ := numbersStack.Pop()
	return StringToFloat(answer)
}

func calcElement(operationStack stack, numbersStack stack) (stack, stack) {
	numbersStack, firstNumStr, _ := numbersStack.Pop()
	firstNum := StringToFloat(firstNumStr)

	numbersStack, secondNumStr, _ := numbersStack.Pop()
	secondNum := StringToFloat(secondNumStr)

	operationStack, operation, _ := operationStack.Pop()

	switch operation {
	case "+":
		numbersStack = numbersStack.Push(FloatToString(float64(secondNum + firstNum)))
	case "-":
		numbersStack = numbersStack.Push(FloatToString(float64(secondNum - firstNum)))
	case "*":
		numbersStack = numbersStack.Push(FloatToString(float64(secondNum * firstNum)))
	case "/":
		numbersStack = numbersStack.Push(FloatToString(float64(secondNum / firstNum)))

	}
	return operationStack, numbersStack
}

func calc(s string) float64 {
	factorizedData := factorize(s)
	return calcExpr(factorizedData)
}

func main() {
	testData := "1+2*(3+4/2-(1+2))*2+1"
	fmt.Println(calc(testData))
}

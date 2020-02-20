package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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
	resArr := make([]string, 0)
	number := ""

	for i := range s {
		if isDigit(string(s[i])) {
			number += string(s[i])
			continue
		}
		if number != "" {
			resArr = append(resArr, number)
			number = ""
		}
		resArr = append(resArr, string(s[i]))
	}
	if number != "" {
		resArr = append(resArr, number)
	}
	return resArr
}

func getPriority(ch string) int {
	const lowPriority = 1
	const highPriority = 2
	const openBrackPriority = 500
	const closeBrackPriority = -2
	const anotherSymbolPriority = 404
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
		return anotherSymbolPriority
	}
}

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', -1, 64)
}

func StringToFloat(st string) (float64, error) {
	num, err := strconv.ParseFloat(strings.TrimSpace(st), 64)
	return num, err
}

func calcExpr(expr []string) (float64, error) {
	operationStack := make(stack, 0)
	numbersStack := make(stack, 0)
	for _, el := range expr {
		if isDigit(el) {
			numbersStack = numbersStack.Push(el)
			continue
		}
		if !operationStack.isEmpty() {
			lastEl, err := operationStack.CheckTop()
			if err != nil {
				return 0, err
			}
			if getPriority(el) > getPriority(lastEl) || lastEl == "(" {
				operationStack = operationStack.Push(el)
				continue
			}
			if el == ")" {
				for lastEl != "(" {
					operationStack, numbersStack, err = calcElement(operationStack, numbersStack)
					if err != nil {
						return 0, err
					}
					lastEl, err = operationStack.CheckTop()
					if err != nil {
						return 0, err
					}
				}
				operationStack, _, err = operationStack.Pop()
				if err != nil {
					return 0, err
				}
				continue
			}
			for getPriority(el) <= getPriority(lastEl) && lastEl != "(" {
				operationStack, numbersStack, err = calcElement(operationStack, numbersStack)
				if err != nil {
					return 0, err
				}
				lastEl, _ = operationStack.CheckTop()
				break
			}
			operationStack = operationStack.Push(el)
			continue
		}
		operationStack = operationStack.Push(el)
	}

	for !operationStack.isEmpty() {
		var err error
		operationStack, numbersStack, err = calcElement(operationStack, numbersStack)
		if err != nil {
			return 0, err
		}
	}
	_, answerStr, err := numbersStack.Pop()
	if err != nil {
		return 0, err
	}
	return StringToFloat(answerStr)
}
func calcElement(operationStack stack, numbersStack stack) (operaionStack stack, numberStack stack, err error) {
	numbersStack, firstNumStr, err := numbersStack.Pop()
	if err != nil {
		return operationStack, numbersStack, err
	}
	firstNum, err := StringToFloat(firstNumStr)
	if err != nil {
		return operationStack, numbersStack, err
	}

	numbersStack, secondNumStr, err := numbersStack.Pop()
	if err != nil {
		return operationStack, numbersStack, err
	}

	secondNum, err := StringToFloat(secondNumStr)
	if err != nil {
		return operationStack, numbersStack, err
	}

	operationStack, operation, err := operationStack.Pop()
	if err != nil {
		return operationStack, numbersStack, err
	}

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
	return operationStack, numbersStack, nil
}

func calc(s string) (float64, error) {
	factorizedData := factorize(s)
	ans, err := calcExpr(factorizedData)
	return ans, err
}

func main() {
	testData := os.Args[len(os.Args)-1]
	ans, err := calc(testData)
	if err != nil {
		log.Fatal("Неккоректные данные")
	}
	fmt.Println(ans)
}

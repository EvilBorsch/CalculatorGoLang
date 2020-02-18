package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	assert := assert.New(t)

	bigData := "1+2*(3+4/2-(1+2))*2+1"
	res := calc(bigData)
	assert.Equal(res, 10.0, "Результат равен 10")

	easyData := "2+2*2"
	res = calc(easyData)
	assert.Equal(res, 6.0)

	easyData = "(2+2)*2"
	res = calc(easyData)
	assert.Equal(res, 8.0)

	easyData = "20+30"
	res = calc(easyData)
	assert.Equal(res, 50.0)

	easyData = "(20+30)"
	res2 := calc(easyData)
	assert.Equal(res2, res)

	hardToParseData := "(20+50-10)/400"
	res = calc(hardToParseData)
	assert.Equal(res, 0.15)

	twoBrackets := "(200-500)/40+(30-40)*2"
	result := calc(twoBrackets)
	assert.Equal(result, -27.5)

	twoBrackets = "(30-40)*2+(200-500)/40"
	result2 := calc(twoBrackets)
	assert.Equal(result, result2)

	twoBrackets = "1/0"
	result2 = calc(twoBrackets)
	assert.Equal(math.Inf(1), result2)

}

func TestIsDigit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(isDigit("2"), true)
	assert.Equal(isDigit("a"), false)
}

func TestFactorize(t *testing.T) {
	assert := assert.New(t)
	var res = []string{"2", "+", "3", "*", "4"}
	assert.Equal(factorize("2+3*4"), res)

	res = []string{"20", "+", "30", "*", "40"}
	assert.Equal(factorize("20+30*40"), res)
}

func TestConvert(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(FloatToString(2.05), "2.050000")
	assert.Equal(StringToFloat("3.0"), float64(3.0))
	assert.Equal(StringToFloat("a"), float64(0))
	assert.Equal(StringToFloat("qweajskhdkjashdkjas"), float64(0))
	assert.Equal(FloatToString(123), "123.000000")

}

func TestElementCalc(t *testing.T) {
	assert := assert.New(t)

	operationStack := make(stack, 0)
	numbersStack := make(stack, 0)

	operationStack = operationStack.Push("+")
	numbersStack = numbersStack.Push("2").Push("2")

	operationStack, numbersStack = calcElement(operationStack, numbersStack)

	numbersStack, res, _ := numbersStack.Pop()
	assert.Equal(res, FloatToString(float64(4)))
	assert.Equal(operationStack.isEmpty(), true)

	operationStack2 := make(stack, 0)
	numbersStack2 := make(stack, 0)

	operationStack2 = operationStack.Push("+").Push("*")
	numbersStack2 = numbersStack.Push("2").Push("3").Push("4").Push("5")

	operationStack2, numbersStack2 = calcElement(operationStack2, numbersStack2)

	numbersStack2, lastNum, _ := numbersStack2.Pop()
	assert.Equal(lastNum, FloatToString(float64(20)))

	operationStack2, lastOperation, _ := operationStack2.Pop()
	assert.Equal(lastOperation, "+")

}

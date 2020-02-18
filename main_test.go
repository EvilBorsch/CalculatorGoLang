package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
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

}

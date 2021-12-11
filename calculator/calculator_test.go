package calculator

import "testing"

type calculatorTest struct {
	numOne, numTwo, expected float64
	operator                 string
}

type operatorTest struct {
	operator string
	expected bool
}

var calculatorTests = []calculatorTest{
	{1, 2, 3, "+"},
	{10, 2, 8, "-"},
	{2, 2, 1, "/"},
	{5, 5, 25, "*"},
}

var operatorTests = []operatorTest{
	{"+", true},
	{"-", true},
	{"/", true},
	{"*", true},
	{"s", false},
}

func TestCalculate(t *testing.T) {
	for _, test := range calculatorTests {
		if output := Calculate(test.numOne, test.operator, test.numTwo); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestCheckOperator(t *testing.T) {
	for _, test := range operatorTests {
		if output := CheckOperator(test.operator); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

package main

import (
	"testing"
)

func InitProduct() Products {
	var products Products
	products = append(products, Product{
		Category: "Smartphone",
		InstallmentFreePeriod: InstallmentPeriod{
			From: 3,
			To:   9,
		},
		Percentage: 3,
	})

	products = append(products, Product{
		Category: "PC",
		InstallmentFreePeriod: InstallmentPeriod{
			From: 3,
			To:   12,
		},
		Percentage: 4,
	})

	products = append(products, Product{
		Category: "TV",
		InstallmentFreePeriod: InstallmentPeriod{
			From: 3,
			To:   18,
		},
		Percentage: 5,
	})

	return products
}

func TestCalculation(t *testing.T) {
	products := InitProduct()
	intervals := []int{3, 6, 9, 12, 18, 24}

	calculator, err := NewCalculator(products, intervals)
	if err != nil {
		t.Errorf("Initialization error: %s", err)
	}

	type TestCase struct {
		Category       string
		Sum            int
		Period         int
		ExpectedAmount int
		ExpectError    string
	}

	testCases := []TestCase{
		{
			Category:       "TV",
			Sum:            1000,
			Period:         19,
			ExpectedAmount: 1050,
		},
		{
			Category:       "Smartphone",
			Sum:            1000,
			Period:         10,
			ExpectedAmount: 1030,
		},
	}

	for _, testCase := range testCases {
		s, err := calculator.GetAmount(testCase.Category, testCase.Sum, testCase.Period)

		if err != nil {
			t.Errorf("Failed: %s", err)
		}

		if s != testCase.ExpectedAmount {
			t.Errorf("Wrong answer. Expected %d, got %d", testCase.ExpectedAmount, s)
		}
	}
}

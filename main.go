package main

import (
	"errors"
)

type InstallmentPeriod struct {
	From int
	To   int
}

type Product struct {
	Category              string
	InstallmentFreePeriod InstallmentPeriod
	Percentage            int
}

type Products []Product

type Calculator struct {
	products  Products
	intervals []int
}

func NewCalculator(products Products, interval []int) (Calculator, error) {
	if len(interval) == 0 {
		return Calculator{}, errors.New("interval must contain at least one value")
	}

	return Calculator{
		products,
		interval,
	}, nil
}

func (c *Calculator) GetAmount(category string, sum, period int) (int, error) {
	if sum < 0 {
		return 0, errors.New("sum must be positive")
	}

	var product *Product
	for _, v := range c.products {
		if v.Category == category {
			product = &v
			break
		}
	}

	if product == nil {
		return 0, errors.New("product not found")
	}

	lastIntervalElement := c.intervals[len(c.intervals)-1]
	if period > lastIntervalElement {
		return 0, errors.New("period does not exist")
	}

	if period < product.InstallmentFreePeriod.From {
		return sum, nil
	}

	if product.InstallmentFreePeriod.From <= period && product.InstallmentFreePeriod.To >= period {
		return sum, nil
	}

	// [3, 6, 9, 12, 18, 24]
	first := 0
	for product.InstallmentFreePeriod.To > c.intervals[first] {
		if product.InstallmentFreePeriod.To == c.intervals[first] {
			break
		}
		first++
	}

	// [3, 6, 9, 12, 18, 24]
	second := 0
	for period > c.intervals[second] {
		if period == c.intervals[second] {
			break
		}
		second++
	}

	distance := second - first
	percentage := distance * product.Percentage

	return sum + sum*percentage/100, nil
}

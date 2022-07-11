package main

import (
	"errors"
	"fmt"
	"log"
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

	return products
}

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

	fmt.Println(first, second)

	distance := second - first
	percentage := distance * product.Percentage

	return sum + sum*percentage/100, nil
}

func main() {
	products := InitProduct()
	intervals := []int{3, 6, 9, 12, 18, 24}

	calculator, err := NewCalculator(products, intervals)
	if err != nil {
		log.Fatalln(err)
	}

	s, err := calculator.GetAmount("Smartphone", 1000, 24)
	fmt.Println(s, err)
}

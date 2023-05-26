package main

import "fmt"

type store struct {
	products []product
}
type product struct {
	productType string
	name        string
	price       int
}

type ProductInterface interface {
	CalculateCost() int
}
type Ecommerce interface {
	Total() int
	Add()
}

func (p product) CalculateCost() int {
	if p.productType == "small" {
		return p.price
	} else if p.productType == "medium" {
		return int(float64(p.price) * 1.03)
	} else if p.productType == "large" {
		return int(float64(p.price)*1.06 + 2500)
	}
	return 0

}
func (s store) Total() int {
	total := 0
	for _, v := range s.products {
		total += v.CalculateCost()
	}
	return total
}

func (s *store) Add(product product) {
	s.products = append(s.products, product)
}

func NewProduct(productType string, name string, price int) product {
	newP := product{productType: productType, name: name, price: price}
	return newP
}

func main() {

	product1 := NewProduct("large", "roupa", 25)

	product2 := NewProduct("medium", "tenis", 250)
	var s store
	s.Add(product1)
	total1 := s.Total()
	fmt.Println(s, total1)

	s.Add(product2)
	total2 := s.Total()
	fmt.Println(s, total2)

}

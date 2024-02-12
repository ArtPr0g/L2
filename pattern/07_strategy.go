package pattern

/*
	Этот паттерн позволяет задавать инкапсулированный алгоритм поведения и взаимодействия с объектами в рамках определенного контекста.
*/


import (
	"fmt"
)

// Интерфейс стратегии
type DiscountStrategy interface {
	ApplyDiscount(price float64) float64
}

// Конкретная стретегия для скидки в 10%
type TenPercentDiscount struct{}

func (t *TenPercentDiscount) ApplyDiscount(price float64) float64 {
	return 0.9 * price
}

// Конкретная стратегия для скидки в 20%
type TwentyPercentDiscount struct{}

func (t *TwentyPercentDiscount) ApplyDiscount(price float64) float64 {
	return 0.8 * price
}

// Контекст, который использует стратегию
type Product struct {
	price    float64
	discount DiscountStrategy
}

func NewProduct(price float64, discount DiscountStrategy) *Product {
	return &Product{
		price:    price,
		discount: discount,
	}
}

func (p *Product) SetDiscount(discount DiscountStrategy) {
	p.discount = discount
}

func (p *Product) ApplyDiscount() {
	p.price = p.discount.ApplyDiscount(p.price)
}

func main() {
	product1 := NewProduct(100.0, &TenPercentDiscount{})
	fmt.Printf("Initial price: $%.2f\n", product1.price)

	product1.ApplyDiscount()
	fmt.Printf("Price after 10%% discount: $%.2f\n", product1.price)

	product1.SetDiscount(&TwentyPercentDiscount{})
	product1.ApplyDiscount()
	fmt.Printf("Price after 20%% discount: $%.2f\n", product1.price)
}

package main

import (
 "errors"
 "fmt"
)

type Product struct {
 ID       int
 Name     string
 Price    float64
 Quantity int
}

type Inventory struct {
 products map[int]Product
}

func (inv *Inventory) AddProduct(p Product) {
 if inv.products == nil {
  inv.products = make(map[int]Product)
 }
 inv.products[p.ID] = p
}

func (inv *Inventory) WriteOff(productID int, quantity int) error {
 p, ok := inv.products[productID]
 if !ok {
  return errors.New("товар не найден")
 }
 if p.Quantity < quantity {
  return errors.New("недостаточно товара")
 }
 p.Quantity -= quantity
 inv.products[productID] = p
 return nil
}

func (inv *Inventory) RemoveProduct(productID int) error {
 _, ok := inv.products[productID]
 if !ok {
  return errors.New("товар не найден")
 }
 delete(inv.products, productID)
 return nil
}

func (inv *Inventory) GetTotalValue() float64 {
 total := 0.0
 for _, p := range inv.products {
  total += p.Price * float64(p.Quantity)
 }
 return total
}

func main() {
 inv := Inventory{}
 inv.AddProduct(Product{1, "Ноутбук", 50000, 3})
 inv.AddProduct(Product{2, "Мышь", 1000, 3})
 fmt.Println("Общая сумма:", inv.GetTotalValue())
}

package main

import "fmt"

type OrderItem struct {
 ProductName string
 Price       float64
 Quantity    int
}

type Customer struct {
 Name string
}

type Order struct {
 Customer Customer
 Items    []OrderItem
 Status   string
}

func (o *Order) AddItem(item OrderItem) {
 o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(name string) {
 newItems := []OrderItem{}
 for _, item := range o.Items {
  if item.ProductName != name {
   newItems = append(newItems, item)
  }
 }
 o.Items = newItems
}

func (o *Order) GetTotal() float64 {
 total := 0.0
 for _, item := range o.Items {
  total += item.Price * float64(item.Quantity)
 }
 return total
}

func (o *Order) ChangeStatus(status string) {
 o.Status = status
}

func main() {
 order := Order{Customer: Customer{"Матвей"}}
 order.AddItem(OrderItem{"Кофе", 200, 2})
 order.AddItem(OrderItem{"Круасан", 300, 1})
 fmt.Println("Сумма заказа:", order.GetTotal())
}

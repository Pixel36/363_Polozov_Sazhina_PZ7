package main

import "fmt"

type EventBus struct {
 handlers map[string][]func(interface{})
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
 if eb.handlers == nil {
  eb.handlers = make(map[string][]func(interface{}))
 }
 eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
 if handlers, ok := eb.handlers[event]; ok {
  for _, h := range handlers {
   h(data)
  }
 }
}

func main() {
 bus := EventBus{}
 bus.Subscribe("greet", func(data interface{}) {
  fmt.Println("Привет,", data)
 })
 bus.Publish("greet", "Карина")
}

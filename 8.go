package main

import "fmt"

type Room struct {
 Number int
 Price  float64
 Booked bool
}

type Hotel struct {
 Name  string
 Rooms []Room
}

type Reservation struct {
 RoomNumber int
 Days       int
}

func (h *Hotel) CheckAvailability() []Room {
 var free []Room
 for _, r := range h.Rooms {
  if !r.Booked {
   free = append(free, r)
  }
 }
 return free
}

func (h *Hotel) BookRoom(number int, days int) *Reservation {
 for i := range h.Rooms {
  if h.Rooms[i].Number == number && !h.Rooms[i].Booked {
   h.Rooms[i].Booked = true
   return &Reservation{number, days}
  }
 }
 return nil
}

func (r *Reservation) GetCost(h *Hotel) float64 {
 for _, room := range h.Rooms {
  if room.Number == r.RoomNumber {
   return float64(r.Days) * room.Price
  }
 }
 return 0
}

func main() {
 h := Hotel{"Гостиница", []Room{{1, 1000, false}, {2, 1500, false}, {3, 3000, false}}}
 res := h.BookRoom(1, 3)
 fmt.Println("Стоимость брони:", res.GetCost(&h))
}

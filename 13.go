package main

import (
 "fmt"
 "time"
)

type Event struct {
 ID           int
 Title        string
 Description  string
 Date         time.Time
 Location     string
 MaxAttendees int
 Attendees    []string
}

type EventManager struct {
 Events []Event
 NextID int
}

func (em *EventManager) CreateEvent(title, desc, loc string, startDate time.Time, max int) *Event {
 e := Event{
  ID:           em.NextID,
  Title:        title,
  Description:  desc,
  Date:         startDate,
  Location:     loc,
  MaxAttendees: max,
 }
 em.Events = append(em.Events, e)
 em.NextID++
 return &em.Events[len(em.Events)-1]
}

func (em *EventManager) Register(eventID int, name string) bool {
 for i := range em.Events {
  e := &em.Events[i]
  if e.ID == eventID && len(e.Attendees) < e.MaxAttendees {
   e.Attendees = append(e.Attendees, name)
   return true
  }
 }
 return false
}

func (em *EventManager) Cancel(eventID int, name string) {
 for i := range em.Events {
  e := &em.Events[i]
  newList := []string{}
  for _, a := range e.Attendees {
   if a != name {
    newList = append(newList, a)
   }
  }
  e.Attendees = newList
 }
}

func (em *EventManager) Upcoming() []Event {
 var res []Event
 now := time.Now()
 for _, e := range em.Events {
  if e.Date.After(now) {
   res = append(res, e)
  }
 }
 return res
}

func main() {
 em := EventManager{}

 start := time.Date(2025, 10, 5, 19, 0, 0, 0, time.Local)

 event := em.CreateEvent(
  "Концерт",
  "Живая музыка",
  "Курган, клуб Опиум",
  start,
  2,
 )

 em.Register(event.ID, "Матвей")
 em.Register(event.ID, "Карина")

 fmt.Println("Будущие мероприятия:", em.Upcoming())
}1

package main

import "fmt"

type ContactInfo struct {
 Type  string
 Value string
}

type Contact struct {
 Name  string
 Infos []ContactInfo
}

type ContactManager struct {
 Contacts []Contact
}

func (cm *ContactManager) AddContact(c Contact) {
 cm.Contacts = append(cm.Contacts, c)
}

func (cm *ContactManager) FindByName(name string) *Contact {
 for _, c := range cm.Contacts {
  if c.Name == name {
   return &c
  }
 }
 return nil
}

func main() {
 cm := ContactManager{}
 cm.AddContact(Contact{"Карина", []ContactInfo{{"phone", "88005553535"}}})
 fmt.Println(cm.FindByName("Карина"))
}

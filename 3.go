package main

import (
 "crypto/sha256"
 "encoding/hex"
 "fmt"
)

type User struct {
 Username string
 Email    string
 Password string
}

func (u *User) SetPassword(password string) {
 hash := sha256.Sum256([]byte(password))
 u.Password = hex.EncodeToString(hash[:])
}

func (u *User) VerifyPassword(password string) bool {
 hash := sha256.Sum256([]byte(password))
 return u.Password == hex.EncodeToString(hash[:])
}

func main() {
 user := User{Username: "user", Email: "test@gmail.com"}
 user.SetPassword("parol")
 fmt.Println("Пароль верный?", user.VerifyPassword("parol"))
 fmt.Println("Пароль верный?", user.VerifyPassword("12345"))
}

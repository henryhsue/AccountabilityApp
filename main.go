package main

import "fmt"

type BalanceError struct {
  Err error
}

func (e *BalanceError) Error() string {
  return e.Err.Error()
}

type User struct {
  Name string
  Balance int
}

func (u User) getBalance() int {
  return u.Balance
}

func (u User) getName() string {
  return u.Name
}

func (u User) incrementBalance(amount int) error {
  if u.Balance + amount > 0 {
    u.Balance += amount 
    return nil
  }
  return BalanceError{}
  
}

func main() {
  u := User{"George", 0}
  fmt.Println("Hello, " + u.getName() + "!")
  fmt.Println("Your balance is:", u.getBalance())
}

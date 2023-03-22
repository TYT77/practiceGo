package main

import (
    "fmt"

    "bank"
)

var accounts = map[float64]*bank.Account{}

func main() {
    
    accounts[1001] = &bank.Account{
        Customer: bank.Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number: 1001,
    }

    
}
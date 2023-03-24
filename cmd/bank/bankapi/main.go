package main

import (
    "fmt"
    "net/http"
    "strconv"
    "log"

    "bank"
)

var accounts = map[float64]*EBank{}

func main() {
    
    accounts[1001] = &EBank{
        Account: &bank.Account{
            Customer: bank.Customer{
                Name:    "John",
                Address: "Los Angeles, California",
                Phone:   "(213) 555 0147",
            },
            Number: 1001,
        }
    }

    accounts[1002] = &EBank{
        Account: &bank.Account{
            Customer: bank.Customer{
                Name:    "yukke",
                Address: "Los Angeles, California",
                Phone:   "(123) 456 7890",
            },
            Number: 1002,
        }
    }

    // エンドポイントの設定
    http.HandleFunc("/statement", statement)
    http.HandleFunc("/deposit", deposit)
    http.HandleFunc("/withdraw", withdraw)
    
    // apiの公開
    fmt.Print("Listening...")
    log.Fatal(http.ListenAndServe(":8080", nil))

}

func statement(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            // json形式のデータを返すように修正
            fmt.Fprintf(w, account.Statement())
        }
    }
}

func deposit(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.Deposit(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

func withdraw(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.Withdraw(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {

                // json形式で返却する
                json.NewEncoder(w).Encode(bank.Statement(account))
                
                // 独自形式で返却する
                // fmt.Fprintf(w, account.Statement())
            }
        }
    }
}


type EBank struct {
    *bank.Account
}

func (eb *EBank) Statement() string {
    
    json, err := json.Marshal(eb)
	if err != nil {
		return err.Error()
	}

	return string(json)

}

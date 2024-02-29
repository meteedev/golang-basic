package main

import (
	"fmt"
	"github.com/meteedev/struct/account"
	"github.com/shopspring/decimal"
)





func main(){
	
	balance , err := init_Balance("1000")

	if(err != nil){
		return
	}

	account_1:=createAccount("Lieang","007",balance)
	account_1.Info()
	
}


func createAccount(name string,id string ,balance decimal.Decimal)account.Account{
	return account.Account {
		Account_name: name,
		Account_id:id,
		Balance:balance,
	}
}

func init_Balance(balance string) (decimal.Decimal ,error){
	
	balance_decimal , err := decimal.NewFromString("1000.0")
	
	if err != nil {
		fmt.Printf("Error converting string '%s' to Decimal: %s\n", balance, err)
		return decimal.Zero, err
	}

	return balance_decimal , nil
}


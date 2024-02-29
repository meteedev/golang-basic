package account

import(
	"fmt"
	"github.com/shopspring/decimal"
)

type Account struct{
	Account_name 	string
	Account_id 		string
	Balance			decimal.Decimal
}


func (a Account)  Info(){
	fmt.Println(a.Account_name)
	fmt.Println(a.Account_id)
	fmt.Println(a.Balance)
}

func (a Account)  view_balance(){
	fmt.Println(a.Balance)
}
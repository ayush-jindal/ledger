package main

import (
	acc "ledger/account"
)


func main() {
	sa := acc.OpenSavingsAccount("Ayush")
	acc.Transfer(75000, nil, &sa, "Salary")
	acc.Transfer(25000, nil, &sa, "Salary")
	acc.Transfer(55000, nil, &sa, "Salary")
	acc.Transfer(95000, nil, &sa, "Salary")
	acc.Transfer(5000, nil, &sa, "Salary")
	acc.Transfer(75000, &sa, nil, "Expenses")

	ta := acc.OpenTradingAccount("Ayush")
	acc.Transfer(5000, &sa, &ta, "Initial investment")
	acc.Transfer(200, &ta, &sa, "Returns")
	sa.Statement()
	ta.Statement()
}

package account

import (
	"fmt"
	"ledger/equity"
)

type TradingAccount struct {
	id int32
	holderName string
	dateOfOpening string
	interestRate float32
	transactions []Transaction
	balance float32
	demat DematAccount
}

func OpenTradingAccount(holderName string) TradingAccount {
	return TradingAccount{0, holderName, "now", 0, make([]Transaction, 0), 0, OpenDematAccount()}
}

func (ta *TradingAccount) deposit(amt float32, from Account, note string) {
	ta.transactions = append(ta.transactions, DepositTransaction(amt, from, ta, note))
	ta.balance += amt
}

func (ta *TradingAccount) withdraw(amt float32, to Account, note string) {
	ta.transactions = append(ta.transactions, WithdrawTransaction(amt, ta, to, note))
	ta.balance -= amt
}

func (ta *TradingAccount) SharesBuy(qty int32, share equity.Share, xchng equity.Exchange) {
	price :=  share.GetPrice(xchng)
	ta.SharesBought(qty, share, price)
}

func (ta *TradingAccount) SharesBought(qty int32, share equity.Share, price float32) {
	amt :=  float32(qty)*price
	ta.withdraw(amt, globalAccount, fmt.Sprintf("%d %s shares @ %g", qty, share.GetDetails(), price))
	ta.balance -= amt
	ta.demat.Deposit(qty, share)
}

func (ta *TradingAccount) SharesSell(qty int32, share equity.Share, xchng equity.Exchange) {
	price := share.GetPrice(xchng)
	ta.SharesSold(qty, share, price)
}

func (ta *TradingAccount) SharesSold(qty int32, share equity.Share, price float32) {
	amt := float32(qty)*price
	ta.deposit(amt, globalAccount, fmt.Sprintf("%d %s shares @ %g", qty, share.GetDetails(), price))
	ta.balance += amt
	ta.demat.Withdraw(qty, share)
}

func (ta *TradingAccount) accountHolderSummary() string {
	return "Trading " + ta.holderName
}

func (ta *TradingAccount) Statement() {
	fmt.Println("***********Trading Account*********")
	for _, v := range ta.transactions {
		v.PrintTransaction()
	}
	fmt.Printf("Balance: %g\n", ta.balance)
}

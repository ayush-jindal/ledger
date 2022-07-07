package account

import "fmt"

type SavingsAccount struct {
	holderName string
	dateOfOpening string
	interestRate float32
	transactions []Transaction
	balance float32
}

func OpenSavingsAccount(holderName string) SavingsAccount {
	return SavingsAccount{holderName, "now", 0, make([]Transaction, 5), 0}
}

func (sa *SavingsAccount) Deposit(amt float32, from Account, note string) {
	sa.transactions = append(sa.transactions, DepositTransaction(amt, from, sa, note))
	sa.balance += amt
}

func (sa *SavingsAccount) Withdraw(amt float32, to Account, note string) {
	sa.transactions = append(sa.transactions, WithdrawTransaction(amt, sa, to, note))
	sa.balance += amt
}

func (sa *SavingsAccount) Statement() {
	fmt.Println("Savings Account")
}

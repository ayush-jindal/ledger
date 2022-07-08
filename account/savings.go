package account

import "fmt"

type SavingsAccount struct {
	id int32
	holderName string
	dateOfOpening string
	interestRate float32
	transactions []Transaction
	balance float32
}

func OpenSavingsAccount(holderName string) SavingsAccount {
	return SavingsAccount{0, holderName, "now", 0, make([]Transaction, 0), 0}
}

func (sa *SavingsAccount) deposit(amt float32, from Account, note string) {
	if from == nil {
		from = globalAccount
	}
	sa.transactions = append(sa.transactions, DepositTransaction(amt, from, sa, note))
	sa.balance += amt
}

func (sa *SavingsAccount) withdraw(amt float32, to Account, note string) {
	if to == nil {
		to = globalAccount
	}
	sa.transactions = append(sa.transactions, WithdrawTransaction(amt, sa, to, note))
	sa.balance -= amt
}

func (sa *SavingsAccount) accountHolderSummary() string {
	return "Savings " + sa.holderName
}

func (sa *SavingsAccount) Statement() {
	fmt.Println("***********Savings Account***********")
	for _, v := range sa.transactions {
		v.PrintTransaction()
	}
	fmt.Printf("Balance: %g\n", sa.balance)
}

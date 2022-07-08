package account

import (
	"fmt"
)

type Transaction struct {
	id string
	from Account
	to Account
	debit bool
	amt float32
	note string
}

func DepositTransaction(amt float32, from, to Account, note string) Transaction {
	return newTransaction(amt, from, to, note, false)
}

func WithdrawTransaction(amt float32, from, to Account, note string) Transaction {
	return newTransaction(amt, from, to, note, true)
}

func newTransaction(amt float32, from, to Account, note string, debit bool) Transaction {
	return Transaction{"1", from, to, debit, amt, note}
}

func (t Transaction) PrintTransaction() {
	if t.debit {
		fmt.Printf("Withdrew %g for %s", t.amt, t.to.accountHolderSummary())
		fmt.Printf("\nRemarks: '%s'\n", t.note)
	} else {
		fmt.Printf("Received %g from %s", t.amt, t.from.accountHolderSummary())
		fmt.Printf("\nRemarks: '%s'\n", t.note)
	}
}

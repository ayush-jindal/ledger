package account

import (
	"fmt"
	"ledger/equity"
)

type DematAccount struct {
	shares map[equity.Share]int32
}


func OpenDematAccount() DematAccount {
	return DematAccount{make(map[equity.Share]int32)}
}

func (da *DematAccount) Deposit(qty int32, share equity.Share) {
	val := da.shares[share]
	da.shares[share] = val+qty
}

func (da *DematAccount) Withdraw(qty int32, share equity.Share) {
	val := da.shares[share]
	da.shares[share] = val-qty
}

func (da *DematAccount) Statement() {
	fmt.Println("Demat Account")
}

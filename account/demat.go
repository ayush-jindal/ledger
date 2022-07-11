package account

import (
	"fmt"
	"ledger/equity"
)

type DematAccount struct {
	shares map[equity.ISIN]int32
}


func OpenDematAccount() DematAccount {
	return DematAccount{make(map[equity.ISIN]int32)}
}

func (da *DematAccount) Deposit(qty int32, share equity.Share) {
	val := da.shares[share.GetID()]
	da.shares[share.GetID()] = val+qty
}

func (da *DematAccount) Withdraw(qty int32, share equity.Share) {
	val := da.shares[share.GetID()]
	da.shares[share.GetID()] = val-qty
}

func (da *DematAccount) Statement() {
	fmt.Println("Demat Account")
}

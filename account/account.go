package account

type Account interface {
	withdraw(amt float32, to Account, note string)
	deposit(amt float32, from Account, note string)
	accountHolderSummary() string
	Statement()
}

func Transfer(amt float32, from, to Account, note string) {
	if to == nil {
		to = globalAccount
	}
	if from == nil {
		from = globalAccount
	}
	from.withdraw(amt, to, note)
	to.deposit(amt, from, note)
}

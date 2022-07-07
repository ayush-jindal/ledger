package account

type Account interface {
	Withdraw(amt float32, to Account, note string)
	Deposit(amt float32, to Account, note string)
	Statement()
}

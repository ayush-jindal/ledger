package account

type gAccount struct {}

func (ga gAccount) withdraw(amt float32, to Account, note string) {}
func (ga gAccount) deposit(amt float32, to Account, note string) {}
func (ga gAccount) accountHolderSummary() string {
	return "VOID"
}
func (ga gAccount) Statement() {}

var globalAccount = gAccount{}

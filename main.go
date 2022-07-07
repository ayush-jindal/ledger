package main

import (
	acc "ledger/account"
)


func main() {
	sa := acc.OpenSavingsAccount("Ayush")
	sa.Statement()
}

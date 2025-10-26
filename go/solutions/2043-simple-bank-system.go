package solutions

type Bank struct {
	balance []int64
	n       int
}

func BankConstructor(balance []int64) Bank {
	return Bank{
		balance,
		len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !(this.isValid(account1) && this.isValid(account2)) {
		return false
	}

	if this.Withdraw(account1, money) {
		return this.Deposit(account2, money)
	}

	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.isValid(account) {
		this.balance[account-1] += money
		return true
	}

	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isValid(account) {
		return false
	}

	if money <= this.balance[account-1] {
		this.balance[account-1] -= money
		return true
	}

	return false
}

func (this *Bank) isValid(account int) bool {
	return account >= 1 && account <= this.n
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

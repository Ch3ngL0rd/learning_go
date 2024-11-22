package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func( t*testing.T) {
		wallet := Wallet{balance : Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, nil)
	})
	t.Run("withdraw insufficient funds", func (t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance : startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, balance Bitcoin) {
	if wallet.Balance() != balance {
		t.Errorf("Expected %s, got %s", balance, wallet.Balance())
	}
}
func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got %q, but got %q", err, want)
	}
}

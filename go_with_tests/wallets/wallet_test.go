package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, balance Bitcoin) {
		if wallet.Balance() != balance {
			t.Errorf("Expected %s, got %s", balance, wallet.Balance())
		}
	}
	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	}

	t.Run("deposit", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func( t*testing.T) {
		wallet := Wallet{balance : Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func (t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance : startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}

package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("didn't get error but wanted one")
		}
		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
 	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
 	}

	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, w, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{fortune: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, w, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		w := Wallet{fortune: Bitcoin(10)}
		err := w.Withdraw(Bitcoin(20))

		want := Bitcoin(10)
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, w, want)
	})
}

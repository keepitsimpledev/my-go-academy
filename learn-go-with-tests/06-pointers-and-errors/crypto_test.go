package crypto

import (
	"fmt"
	"testing"
)

func assertBitcoin(tb testing.TB, want int, got Bitcoin) {
	tb.Helper()

	wantBitcoin := Bitcoin(want)
	if got != wantBitcoin {
		tb.Errorf("got %s want %s", got, wantBitcoin)
	}
}

func assertError(tb testing.TB, got error, want error) {
	tb.Helper()

	if got == nil {
		tb.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		tb.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(tb testing.TB, err error) {
	tb.Helper()

	if err != nil {
		tb.Error("got an error but didn't expect one")
	}
}

func TestWallet(t *testing.T) {
	t.Run("balance", func(t *testing.T) {
		// arrange
		wallet := Wallet{balance: 9}

		// act
		got := wallet.Balance()

		// assert
		assertBitcoin(t, 9, got)
	})

	t.Run("deposit", func(t *testing.T) {
		// arrange
		wallet := Wallet{balance: 0}

		// act
		wallet.Deposit(10)
		got := wallet.Balance()
		fmt.Printf("address of balance in test is %p \n", &wallet.balance) // this line is for learning usage of &

		// assert
		assertBitcoin(t, 10, got)
	})

	t.Run("withdraw some", func(t *testing.T) {
		// arrange
		wallet := Wallet{balance: 100}

		// act
		err := wallet.Withdraw(90)

		// assert
		got := wallet.Balance()
		assertBitcoin(t, 10, got)
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		// arrange
		wallet := Wallet{balance: 20}

		// act
		err := wallet.Withdraw(Bitcoin(100))

		// assert
		assertBitcoin(t, 20, wallet.balance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

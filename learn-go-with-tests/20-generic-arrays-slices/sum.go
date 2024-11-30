package genericlists

type Transaction struct {
	From, To string
	Sum      float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  amount,
	}
}

func NewBalanceFor(acc Account, transactions []Transaction) Account {
	transact := func(acc Account, transaction Transaction) Account {
		if transaction.From == acc.Name {
			acc.Balance -= transaction.Sum
		}

		if transaction.To == acc.Name {
			acc.Balance += transaction.Sum
		}

		return acc
	}

	return Reduce(transactions, transact, acc)
}

func BalanceFor(transactions []Transaction, name string) float64 {
	transact := func(balance float64, transaction Transaction) float64 {
		if transaction.From == name {
			balance -= transaction.Sum
		}

		if transaction.To == name {
			balance += transaction.Sum
		}

		return balance
	}

	return Reduce(transactions, transact, 0)
}

func Reduce[InType, OutType any](items []InType, fn func(OutType, InType) OutType, initialValue OutType) OutType {
	result := initialValue
	for _, item := range items {
		result = fn(result, item)
	}

	return result
}

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := func(a, b int) int { return a + b }
	return Reduce(numbers, sum, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	tailsSum := func(result, current []int) []int {
		if len(current) == 0 {
			result = append(result, 0)
		} else {
			tail := current[1:]
			result = append(result, Sum(tail))
		}

		return result
	}

	return Reduce(numbersToSum, tailsSum, []int{})
}

package sum

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}

type Person struct {
	Name string
}

// func BalanceFor(transactions []Transaction, name string) float64 {
// 	adjustBalance := func(currentBalance float64, t Transaction) float64 {
// 		if t.From == name {
// 			return currentBalance - t.Sum
// 		}
// 		if t.To == name {
// 			return currentBalance + t.Sum
// 		}
// 		return currentBalance
// 	}

// 	return Reduce(transactions, adjustBalance, 0.0)

// 	// var balance float64
// 	// for _, t := range transactions {
// 	// 	if t.From == name {
// 	// 		balance -= t.Sum
// 	// 	}
// 	// 	if t.To == name {
// 	// 		balance += t.Sum
// 	// 	}
// 	// }
// 	// return balance
// }

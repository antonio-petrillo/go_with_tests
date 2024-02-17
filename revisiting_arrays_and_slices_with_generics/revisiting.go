package revisiting

// func Reduce[A any] (collection []A, accumulator func(A, A) A, initialValue A) A {
// 	var result = initialValue
// 	for _, x := range collection {
// 		result = accumulator(result, x)
// 	}
// 	return result
// }

func Reduce[A, B any] (collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

type Transaction struct {
	From string
	To string
	Sum float64
}

type Account struct {
	Name string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, t Transaction) Account{
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

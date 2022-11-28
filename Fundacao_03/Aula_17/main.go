package main

import "fmt"

type BankAccount struct {
	balance     float64
}

// Sem o * (indicador de ponteiro), eu só executo o método sem alterar valores fora da func
// com o * (indicador de ponteiro), eu estou alterando dentro do ponteiro, de uma forma mais 'Global'
func (account *BankAccount) simulation(value float64) float64{
	account.balance += value
	fmt.Printf("O valor da simulação é %v\n", account.balance)
	return account.balance
}

func main() {
	account := BankAccount {
		balance: 100,
	}
	account.simulation(200)
	println(account.balance)
	fmt.Printf("O valor da conta é %v", account.balance)
}

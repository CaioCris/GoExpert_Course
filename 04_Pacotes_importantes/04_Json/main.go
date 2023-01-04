package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	// Marshal guarda o valor em uma variavel
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Como encoder, você já envia o valor para algum lugar (web server, arquivo, stdout)
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	// Transformando Json em struct
	jsonInBytes := []byte(`{"Number":2,"Balance":200}`)
	var account_02 Account
	err = json.Unmarshal(jsonInBytes, &account_02)
	if err != nil {
		println(err)
	}
	println(account_02.Balance)

	// Testando as Tags na struct
	jsonInBytesWithTags := []byte(`{"number":3,"balance":300}`)
	var account_03 Account
	err = json.Unmarshal(jsonInBytesWithTags, &account_03)
	if err != nil {
		println(err)
	}
	println(account_03.Balance)

}

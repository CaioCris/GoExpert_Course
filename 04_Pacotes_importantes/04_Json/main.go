package main

import (
	"encoding/json"
	"os"
)

type account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	account01 := account{Number: 1, Balance: 100}
	// Marshal guarda o valor em uma variavel
	res, err := json.Marshal(account01)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Como encoder, você já envia o valor para algum lugar (web server, arquivo, stdout)
	err = json.NewEncoder(os.Stdout).Encode(account01)
	if err != nil {
		println(err)
	}

	// Transformando Json em struct
	jsonInBytes := []byte(`{"Number":2,"Balance":200}`)
	var account02 account
	err = json.Unmarshal(jsonInBytes, &account02)
	if err != nil {
		println(err)
	}
	println(account02.Balance)

	// Testando as Tags na struct
	jsonInBytesWithTags := []byte(`{"number":3,"balance":300}`)
	var account03 account
	err = json.Unmarshal(jsonInBytesWithTags, &account03)
	if err != nil {
		println(err)
	}
	println(account03.Balance)

}

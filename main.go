package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("alphabet.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data)) // conversion de byte en string
}

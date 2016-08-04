package main

import "github.com/andrewdruzhinin/go-russianpost/russianpost"

func main() {
	client := russianpost.NewClient("login", "password")
	client.GetOperationHistory("RA644000001RU", "0", "RUS")
}

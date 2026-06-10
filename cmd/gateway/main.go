package main

import (
	"fmt"
	"mart-gateway/internal/bank"
)

func main() {
	err := bank.SendAuthorizationRequestToBank()
	fmt.Println("error o")
	fmt.Println(err)
	
}
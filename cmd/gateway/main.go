package main

import (
	"fmt"
	"mart-gateway/internal/bank"
)

func main() {
	// err := bank.SendAuthorizationRequestToBank()
	err := bank.SendCaptureRequest()
	fmt.Println(err, "captured")
}
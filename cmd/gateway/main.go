package main

import (
	"fmt"
	"mart-gateway/internal/bank"
)

func main() {
	// err := bank.SendAuthorizationRequestToBank()
	// err := bank.SendVoidRequest()
	// err := bank.SendCaptureRequest()
	err := bank.SendRefundRequest()
	fmt.Println(err, "refund")
}
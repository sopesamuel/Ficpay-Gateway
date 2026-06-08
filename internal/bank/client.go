package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"mart-gateway/internal/models"
)


func SendAuthorizationRequestToBank(){
	posturl := "http://localhost:8787/"

	req := models.Bankauthrequest{
			Amount: 10,
			Carddetails : models.Card{ Card_number : "4111111111111111", Cvv : "123", Expiry_month : 12, Expiry_year: 2030,},
		}


	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error message for json marshalling our structs for now!")
	}

	_, err = http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error message for creating post request for authorization for now!")
	}


}
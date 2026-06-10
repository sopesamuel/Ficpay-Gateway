package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mart-gateway/internal/models"
	"net/http"
)


func SendAuthorizationRequestToBank() (error){
	posturl := "http://localhost:8787/api/v1/authorizations"

	req := models.Bankauthrequest{
			Amount: 10,
			Carddetails : models.Card{ Card_number : "4111111111111111", Cvv : "123", Expiry_month : 12, Expiry_year: 2030,},
		}


	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error message for json marshalling our structs for now!")
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error message for creating post request for authorization for now!")
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Idempotency-Key", "2")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil{
		return fmt.Errorf("Error in http client creating request -%s", err.Error())
	}

	defer res.Body.Close()

	errormessage := &models.Errorresponse{}
	if res.StatusCode != http.StatusCreated {
		err = json.NewDecoder(res.Body).Decode(errormessage)
		if err != nil {
			return fmt.Errorf("Error decoding error response body -%s", string(err.Error()))
		}
	}

	post := &models.Bankauthresponse{}
	err = json.NewDecoder(res.Body).Decode(post)
	if err != nil{
		return fmt.Errorf("Error decoding message response body -%s", string(err.Error()))
		}

	fmt.Println(post.Amount)
	fmt.Println(post.Authorization_id)
	fmt.Println(post.Created_at)
	fmt.Println(post.Expires_at)
	fmt.Println(post.Currency)
	fmt.Println(post.Status)

	return nil
}
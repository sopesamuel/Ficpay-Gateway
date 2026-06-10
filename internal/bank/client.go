package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mart-gateway/internal/models"
	"net/http"
)
type BankInterface interface {
	SendAuthorizationRequestToBank() 
}

type Bankclientstruct struct {
	baseURL string
	httpClient *http.Client
}


func SendAuthorizationRequestToBank() (error){
	posturl := "http://localhost:8787/api/v1/authorizations"

	req := models.Bankauthrequest{
			 Amount: 10, Card_number : "4111111111111111", Cvv : "123", Expiry_month : 12, Expiry_year: 2030,
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
	r.Header.Add("Idempotency-Key", "4")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil{
		return fmt.Errorf("Error in http client creating request -%s", err.Error())
	}

	defer res.Body.Close()

	errormessage := &models.Errorresponse{}
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := &models.Bankauthresponse{}
	err = json.NewDecoder(res.Body).Decode(post)
	if err != nil{
		return fmt.Errorf("Error decoding message response body -%v", res.Status)
		}

	fmt.Println(post.Amount)
	fmt.Println(post.Authorization_id)
	fmt.Println(post.Created_at)
	fmt.Println(post.Expires_at)
	fmt.Println(post.Currency)
	fmt.Println(post.Status)

	return nil
}

func SendCaptureRequest() (error){

	posturl := "http://localhost:8787/api/v1/captures"

	capreq := models.Bankcapturerequest{
		Amount: 10,
		Authorization_id : "auth_c0ceb55e-f8b0-4109-82b9-283b0c73587e",
	}

	body, err := json.Marshal(capreq)
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
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := &models.Captureresponse{}
	err = json.NewDecoder(res.Body).Decode(post)
	if err != nil{
		return fmt.Errorf("Error decoding message response body -%v", res.Status)
		}

	

	fmt.Println(post.Amount)
	fmt.Println(post.Authorization_id)
	fmt.Println(post.Capture_id)
	fmt.Println(post.Captured_at)
	fmt.Println(post.Currency)
	fmt.Println(post.Status)

	return nil

}

func SendVoidRequest() (error){


	posturl := "http://localhost:8787/api/v1/voids"

	capreq := models.Bankvoidrequest{
		Authorization_id : "auth_13bb859d-c4b8-4ed4-aafc-efc1899591e2",
	}

	body, err := json.Marshal(capreq)
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
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := &models.Voidresponse{}
	err = json.NewDecoder(res.Body).Decode(post)
	if err != nil{
		return fmt.Errorf("Error decoding message response body -%v", res.Status)
		}

	
	fmt.Println(post.Authorization_id)
	fmt.Println(post.Void_id)
	fmt.Println(post.Voided_At)
	fmt.Println(post.Status)

	return nil
}


func SendRefundRequest() (error){


	posturl := "http://localhost:8787/api/v1/refunds"

	capreq := models.Bankrefundrequest{
		Amount: 10,
		Capture_id: "cap_da145645-203d-4087-8754-8fe458b2a59a",	
	}

	body, err := json.Marshal(capreq)
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
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := &models.Refundresponse{}
	err = json.NewDecoder(res.Body).Decode(post)
	if err != nil{
		return fmt.Errorf("Error decoding message response body -%v", res.Status)
		}

	
	fmt.Println(post.Amount)
	fmt.Println(post.Refund_id)
	fmt.Println(post.Refunded_at)
	fmt.Println(post.Status)

	return nil
}
package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mart-gateway/internal/models"
	"net/http"
)

//This holds the shared state used by all these methods.
type Bankclientstruct struct {
	baseURL string
	httpClient *http.Client
}

func NewClient(base_url string) *Bankclientstruct{
	return &Bankclientstruct{
		baseURL: base_url,
		httpClient: &http.Client{},
	}
}

func (cfg *Bankclientstruct) Authorize(bank_auth models.Bankauthrequest, key string) (models.Bankauthresponse, error){

	body, err := json.Marshal(bank_auth)
	if err != nil {
		return models.Bankauthresponse{}, fmt.Errorf("Error json marshalling our bank authorization request struct! %w", err)
	}

	r, err := http.NewRequest("POST", cfg.baseURL + "/api/v1/authorizations", bytes.NewBuffer(body))
	if err != nil {
		return models.Bankauthresponse{}, fmt.Errorf("Error creating new post request to the bank for authorization! %w", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Idempotency-Key", key)

	res, err := cfg.httpClient.Do(r)
	if err != nil{
		return models.Bankauthresponse{}, fmt.Errorf("Error configuring new client to the bank for authorization!- %w", err)
	}

	defer res.Body.Close()

	errormessage := &models.Errorresponse{}
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return models.Bankauthresponse{}, fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := models.Bankauthresponse{}
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil{
		return models.Bankauthresponse{}, fmt.Errorf("Error decoding message response body for bank authorization -%v, %w", res.Status, err)
	}

	return post,nil
}


func (cfg *Bankclientstruct) Capture(bankcapture_req models.Bankcapturerequest, key string) (models.Captureresponse,error){

	body, err := json.Marshal(bankcapture_req)
	if err != nil {
		return models.Captureresponse{}, fmt.Errorf("Error json marshalling our bank capture request struct! %w", err)
	}

	r, err := http.NewRequest("POST", cfg.baseURL + "/api/v1/captures", bytes.NewBuffer(body))
	if err != nil {
		return models.Captureresponse{}, fmt.Errorf("Error creating new post request to the bank for capture request! %w", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Idempotency-Key", key)

	
	res, err := cfg.httpClient.Do(r)
	if err != nil{
		return models.Captureresponse{}, fmt.Errorf("Error configuring new client to the bank for capture request! -%w", err)
	}

	defer res.Body.Close()

	errormessage := &models.Errorresponse{}
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return models.Captureresponse{}, fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := models.Captureresponse{}
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil{
		return models.Captureresponse{},  fmt.Errorf("Error decoding message response body for bank capture request -%v, %w", res.Status, err)
		}

	return post, nil
}

func (cfg *Bankclientstruct) Void(bank_void models.Bankvoidrequest, key string) (models.Voidresponse, error){

	body, err := json.Marshal(bank_void)
	if err != nil {
		return models.Voidresponse{}, fmt.Errorf("Error json marshalling our bank void request struct! %w", err)
	}

	r, err := http.NewRequest("POST", cfg.baseURL + "/api/v1/voids", bytes.NewBuffer(body))
	if err != nil {
		return models.Voidresponse{}, fmt.Errorf("Error creating new post request to the bank for void request! %w", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Idempotency-Key", key)


	
	res, err := cfg.httpClient.Do(r)
	if err != nil{
		return models.Voidresponse{} , fmt.Errorf("Error configuring new client to the bank for void request! -%w", err)
	}

	defer res.Body.Close()

	errormessage := &models.Errorresponse{}
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return models.Voidresponse{},fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := models.Voidresponse{}
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil{
		return models.Voidresponse{}, fmt.Errorf("Error decoding message response body response for bank void request-%w", err)
	}

	return post,nil
}


func (cfg *Bankclientstruct) Refund(refund_req models.Bankrefundrequest, key string) (models.Refundresponse , error){

	body, err := json.Marshal(refund_req)
	if err != nil {
		return models.Refundresponse{} , fmt.Errorf("Error json marshalling our bank refund request struct! %w", err)
	}

	r, err := http.NewRequest("POST", cfg.baseURL + "/api/v1/refunds", bytes.NewBuffer(body))
	if err != nil {
		return models.Refundresponse{} , fmt.Errorf("Error creating new post request to the bank for refund request! %w", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Idempotency-Key", key)

	res, err := cfg.httpClient.Do(r)
	if err != nil{
		return models.Refundresponse{} , fmt.Errorf("Error configuring new client to the bank for refund request! -%w", err)
	}

	defer res.Body.Close()

	errormessage := &models.Errorresponse{}
	if res.StatusCode != http.StatusOK {
		json.NewDecoder(res.Body).Decode(errormessage)
		return models.Refundresponse{},fmt.Errorf("%s: with %s %v", errormessage.Message, errormessage.Error, res.Status)
	}

	post := models.Refundresponse{}
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil{
		return models.Refundresponse{},fmt.Errorf("Error decoding message response body response for bank refund request-%w", err)
		}

	return post, nil
}
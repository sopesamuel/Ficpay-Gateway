package service

import (
	"mart-gateway/internal/models"
)


type BankInterface interface {
	Authorize(req models.Bankauthrequest, idempotencyKey string) (models.Bankauthresponse, error)
	Capture(req models.Bankcapturerequest, idempotencyKey string) (models.Captureresponse, error)
	Void(req models.Bankvoidrequest, idempotencyKey string) (models.Voidresponse, error)
	Refund(req models.Bankrefundrequest, idempotencyKey string) (models.Refundresponse, error)
}
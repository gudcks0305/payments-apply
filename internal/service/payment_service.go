package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gudcks0305/payments-apply/internal/dto"
	"github.com/gudcks0305/payments-apply/internal/model"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/internal/repository"
)

type PaymentService struct {
	repository    *repository.PaymentRepository
	portoneClient *portone.Client
}

func NewPaymentService(repository *repository.PaymentRepository, portoneClient *portone.Client) *PaymentService {
	return &PaymentService{repository: repository, portoneClient: portoneClient}
}

func (s *PaymentService) CreatePayment(d *dto.PaymentCreateRequest) (*dto.IdResponse[uuid.UUID], error) {
	payment := &model.Payment{
		Amount:      d.Amount,
		PayMethod:   d.PayMethod,
		ProductName: d.ProductName,
		Status:      model.StatusPending,
	}
	err := s.repository.CreatePayment(payment)
	if err != nil {
		return nil, err
	}
	return &dto.IdResponse[uuid.UUID]{ID: payment.ID}, nil
}

func (s *PaymentService) ConfirmWithCompletePayment(p *portone.PaymentData) (interface{}, error) {
	var result = &portone.PaymentData{}
	err := s.portoneClient.GetPayment(p.ImpUID, result)
	if err != nil {
		return nil, err
	}

	err = validate(p, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "paid" {
		cancelReq := portone.PaymentCancelRequest{
			ImpUID:      p.ImpUID,
			MerchantUID: result.MerchantUID,
			Amount:      result.PaidAmount,
			Reason:      "TEST",
		}
		err := s.portoneClient.CancelPayment(cancelReq, nil)
		if err != nil {
			return nil, err
		}
		s.UpdatePaymentModel(p)
	}

	return result, nil
}

func (s *PaymentService) UpdatePaymentModel(p *portone.PaymentData) {
	tx := s.repository.DB.Begin()
	defer tx.Rollback()

	id, _ := uuid.FromBytes([]byte(p.MerchantUID))
	paymentModel, _ := s.repository.GetPaymentByID(id)
	paymentModel.Status = model.StatusCancelled
	paymentModel.ImpUID = p.ImpUID

	tx.Save(paymentModel)
	tx.Commit()
}

func (s *PaymentService) GetPaymentByIMPUID(impUID string) (interface{}, error) {
	var res *portone.PaymentData
	err := s.portoneClient.GetPayment(impUID, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func validate(d, res *portone.PaymentData) error {
	if d.PaidAmount != res.PaidAmount {
		return fmt.Errorf("invalid paid amount")
	}
	return nil
}

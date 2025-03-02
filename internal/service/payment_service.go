package service

import (
	"github.com/gudcks0305/payments-apply/internal/errors"
	errors2 "github.com/pkg/errors"

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

func (s *PaymentService) ConfirmWithCompletePayment(p *portone.PaymentClientResponse) (interface{}, error) {
	var result = &portone.APIResponse[portone.PaymentData]{}
	err := s.portoneClient.GetPayment(p.ImpUid, result)
	if err != nil {
		return nil, err
	}

	res := result.Response
	err = validate(p, &res)
	if err != nil {
		return nil, err
	}
	if res.Status == "paid" {
		cancelReq := portone.PaymentCancelRequest{
			ImpUID:      p.ImpUid,
			MerchantUID: res.MerchantUID,
			Amount:      res.Amount,
			Reason:      "TEST",
		}
		resp := &portone.APIResponse[portone.PaymentData]{}
		err := s.portoneClient.CancelPayment(cancelReq, resp)
		if err != nil {
			return nil, err
		}
		err = s.UpdatePaymentModel(&resp.Response)
		if err != nil {
			return nil, err
		}
		// 취소 Resp 가 제대로 내려가지 않으니 확인 필요
		return resp.Response, nil
	}

	return res, nil
}

func (s *PaymentService) UpdatePaymentModel(p *portone.PaymentData) error {
	tx := s.repository.DB.Begin()

	id, _ := uuid.Parse(p.MerchantUID)
	paymentModel, _ := s.repository.GetPaymentByID(id)
	paymentModel.Status = model.StatusCancelled
	paymentModel.ImpUID = p.ImpUID

	err := tx.Save(paymentModel).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *PaymentService) GetPaymentByIMPUID(impUID string) (interface{}, error) {
	var res = &portone.APIResponse[portone.PaymentData]{}
	err := s.portoneClient.GetPayment(impUID, res)
	if err != nil {
		return nil, errors2.Wrap(errors.ErrPortOneError, err.Error())
	}
	return res.Response, nil
}

func validate(d *portone.PaymentClientResponse, res *portone.PaymentData) error {
	if d.PaidAmount != res.Amount {
		return errors.ErrInvalidAmount
	}
	return nil
}

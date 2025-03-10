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
	portoneClient portone.POClient
}

func NewPaymentService(repository *repository.PaymentRepository, portoneClient portone.POClient) *PaymentService {
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

func (s *PaymentService) ConfirmWithCompletePayment(id string, p *portone.PaymentClientResponse) (interface{}, error) {
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
	err = s.UpdatePaymentModel(id, &res, model.StatusPending)

	if res.Status == "paid" {
		cancelReq := portone.PaymentCancelRequest{
			ImpUID:      p.ImpUid,
			MerchantUID: &p.MerchantUid,
			Reason:      nil,
		}
		resp := &portone.APIResponse[portone.PaymentData]{}
		err := s.portoneClient.CancelPayment(cancelReq, resp)
		if err != nil {
			return nil, err
		}
		err = s.UpdatePaymentModel(id, &resp.Response, model.StatusCancelled)
		if err != nil {
			return nil, err
		}
		// 취소 Resp 가 제대로 내려가지 않으니 확인 필요
		return resp.Response, nil
	}
	err = s.UpdatePaymentModel(id, &res, model.StatusCompleted)
	return res, nil
}

func (s *PaymentService) ConfirmWithCompletePaymentBasic(p *dto.PaymentBasicConfirmRequest) (portone.PaymentData, error) {
	var result = &portone.APIResponse[portone.PaymentData]{}
	err := s.portoneClient.GetPayment(p.ImpUID, result)
	if err != nil {
		return portone.PaymentData{}, err
	}

	res := result.Response

	if res.Status == "paid" {
		cancelReq := portone.PaymentCancelRequest{
			ImpUID:      res.ImpUID,
			MerchantUID: &res.MerchantUID,
			Reason:      nil,
		}
		resp := &portone.APIResponse[portone.PaymentData]{}
		err := s.portoneClient.CancelPayment(cancelReq, resp)
		if err != nil {
			return portone.PaymentData{}, err
		}
		return resp.Response, nil
	}
	return res, nil
}

func (s *PaymentService) UpdatePaymentModel(id string, p *portone.PaymentData, statusType model.PaymentStatusType) error {
	tx := s.repository.DB.Begin()

	uuids, _ := uuid.Parse(id)
	paymentModel, _ := s.repository.GetPaymentByID(uuids)
	paymentModel.Status = statusType
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

func (s *PaymentService) CancelPaymentByIMPUID(uid string) (interface{}, error) {
	var res = &portone.APIResponse[portone.PaymentData]{}
	err := s.portoneClient.GetPayment(uid, res)
	if err != nil {
		return nil, errors2.Wrap(errors.ErrPortOneError, err.Error())
	}
	reason := "TEST"
	cancelReq := portone.PaymentCancelRequest{
		ImpUID:      uid,
		MerchantUID: &res.Response.MerchantUID,
		Reason:      &reason,
	}
	resp := &portone.APIResponse[portone.PaymentData]{}
	err = s.portoneClient.CancelPayment(cancelReq, resp)
	if err != nil {
		return nil, errors2.Wrap(errors.ErrPortOneError, err.Error())
	}
	return resp.Response, nil
}

func validate(d *portone.PaymentClientResponse, res *portone.PaymentData) error {
	if d.PaidAmount != res.Amount {
		return errors.ErrInvalidAmount
	}
	return nil
}

package dto

import "github.com/gudcks0305/payments-apply/internal/errors"

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type IdResponse[T any] struct {
	ID T `json:"id"`
}

func APIResponseSuccess[T any](data T) APIResponse[T] {
	return APIResponse[T]{Code: 200, Message: "Success", Data: data}
}

func APIResponseCreated[T any](response T) APIResponse[T] {
	return APIResponse[T]{Code: 201, Message: "Created", Data: response}
}

func APIResponseError[T any](appError errors.AppError) APIResponse[T] {
	return APIResponse[T]{Code: appError.StatusCode, Message: appError.Message}
}

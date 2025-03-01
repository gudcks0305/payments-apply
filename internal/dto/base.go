package dto

type APIResponse[T any] struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response T      `json:"response"`
}

type IdResponse[T any] struct {
	ID T `json:"id"`
}

func APIResponseSuccess[T any](response T) APIResponse[T] {
	return APIResponse[T]{Code: 200, Message: "Success", Response: response}
}

func APIResponseCreated[T any](response T) APIResponse[T] {
	return APIResponse[T]{Code: 201, Message: "Created", Response: response}
}

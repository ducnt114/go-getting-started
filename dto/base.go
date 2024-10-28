package dto

import "net/http"

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	SuccessMeta = &Meta{
		Code:    http.StatusOK,
		Message: "success",
	}
	InternalServerErrorMeta = &Meta{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
	BadRequestMeta = &Meta{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
)

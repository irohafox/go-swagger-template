package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"golang-rest-api-template/app/gen/restapi/factory"
)

func Test(p factory.TestParams) middleware.Responder {
	payload := "test"
	return factory.NewTestOK().WithPayload(payload)
}

package test

import (
	"goyave.dev/goyave/v3"
	"net/http/httptest"
)

func GetReqAndRes() (*goyave.Request, *goyave.Response) {
	suite := goyave.TestSuite{}
	request := suite.CreateTestRequest(nil)
	response := suite.CreateTestResponse(httptest.NewRecorder())
	return request, response
}

// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package handlers

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsHandler

type MockTransactionResultsHandler struct {
	mock.Mock
}

func (s *MockTransactionResultsHandler) HandleTransactionResults(input *HandleTransactionResultsInput) (*HandleTransactionResultsOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleTransactionResultsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}


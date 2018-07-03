// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service VirtualMachine

type MockVirtualMachine struct {
	mock.Mock
}

func (s *MockVirtualMachine) ProcessTransactionSet(input *ProcessTransactionSetInput) (*ProcessTransactionSetOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ProcessTransactionSetOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) RunLocalMethod(input *RunLocalMethodInput) (*RunLocalMethodOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*RunLocalMethodOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) TransactionSetPreOrder(input *TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*TransactionSetPreOrderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) SdkCall(input *SdkCallInput) (*SdkCallOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SdkCallOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}


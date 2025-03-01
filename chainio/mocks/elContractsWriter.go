// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts (interfaces: Writer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/elContractsWriter.go -package=mocks -mock_names=Writer=MockELWriter github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts Writer
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	contractIRewardsCoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	types "github.com/Layr-Labs/eigensdk-go/types"
	common "github.com/ethereum/go-ethereum/common"
	types0 "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockELWriter is a mock of Writer interface.
type MockELWriter struct {
	ctrl     *gomock.Controller
	recorder *MockELWriterMockRecorder
}

// MockELWriterMockRecorder is the mock recorder for MockELWriter.
type MockELWriterMockRecorder struct {
	mock *MockELWriter
}

// NewMockELWriter creates a new mock instance.
func NewMockELWriter(ctrl *gomock.Controller) *MockELWriter {
	mock := &MockELWriter{ctrl: ctrl}
	mock.recorder = &MockELWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockELWriter) EXPECT() *MockELWriterMockRecorder {
	return m.recorder
}

// DepositERC20IntoStrategy mocks base method.
func (m *MockELWriter) DepositERC20IntoStrategy(arg0 context.Context, arg1 common.Address, arg2 *big.Int) (*types0.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositERC20IntoStrategy", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DepositERC20IntoStrategy indicates an expected call of DepositERC20IntoStrategy.
func (mr *MockELWriterMockRecorder) DepositERC20IntoStrategy(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositERC20IntoStrategy", reflect.TypeOf((*MockELWriter)(nil).DepositERC20IntoStrategy), arg0, arg1, arg2)
}

// ProcessClaim mocks base method.
func (m *MockELWriter) ProcessClaim(arg0 context.Context, arg1 contractIRewardsCoordinator.IRewardsCoordinatorRewardsMerkleClaim, arg2 common.Address) (*types0.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessClaim", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessClaim indicates an expected call of ProcessClaim.
func (mr *MockELWriterMockRecorder) ProcessClaim(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessClaim", reflect.TypeOf((*MockELWriter)(nil).ProcessClaim), arg0, arg1, arg2)
}

// RegisterAsOperator mocks base method.
func (m *MockELWriter) RegisterAsOperator(arg0 context.Context, arg1 types.Operator) (*types0.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAsOperator", arg0, arg1)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAsOperator indicates an expected call of RegisterAsOperator.
func (mr *MockELWriterMockRecorder) RegisterAsOperator(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAsOperator", reflect.TypeOf((*MockELWriter)(nil).RegisterAsOperator), arg0, arg1)
}

// SetClaimerFor mocks base method.
func (m *MockELWriter) SetClaimerFor(arg0 context.Context, arg1 common.Address) (*types0.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetClaimerFor", arg0, arg1)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetClaimerFor indicates an expected call of SetClaimerFor.
func (mr *MockELWriterMockRecorder) SetClaimerFor(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClaimerFor", reflect.TypeOf((*MockELWriter)(nil).SetClaimerFor), arg0, arg1)
}

// UpdateMetadataURI mocks base method.
func (m *MockELWriter) UpdateMetadataURI(arg0 context.Context, arg1 string) (*types0.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetadataURI", arg0, arg1)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMetadataURI indicates an expected call of UpdateMetadataURI.
func (mr *MockELWriterMockRecorder) UpdateMetadataURI(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetadataURI", reflect.TypeOf((*MockELWriter)(nil).UpdateMetadataURI), arg0, arg1)
}

// UpdateOperatorDetails mocks base method.
func (m *MockELWriter) UpdateOperatorDetails(arg0 context.Context, arg1 types.Operator) (*types0.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOperatorDetails", arg0, arg1)
	ret0, _ := ret[0].(*types0.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOperatorDetails indicates an expected call of UpdateOperatorDetails.
func (mr *MockELWriterMockRecorder) UpdateOperatorDetails(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOperatorDetails", reflect.TypeOf((*MockELWriter)(nil).UpdateOperatorDetails), arg0, arg1)
}

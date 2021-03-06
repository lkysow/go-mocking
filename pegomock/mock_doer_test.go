// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/lkysow/go-mocking/pegomock (interfaces: Doer)

package pegomock_test

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockDoer struct {
	fail func(message string, callerSkip ...int)
}

func NewMockDoer() *MockDoer {
	return &MockDoer{fail: pegomock.GlobalFailHandler}
}

func (mock *MockDoer) DoIt() error {
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DoIt", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockDoer) VerifyWasCalledOnce() *VerifierDoer {
	return &VerifierDoer{mock, pegomock.Times(1), nil}
}

func (mock *MockDoer) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierDoer {
	return &VerifierDoer{mock, invocationCountMatcher, nil}
}

func (mock *MockDoer) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierDoer {
	return &VerifierDoer{mock, invocationCountMatcher, inOrderContext}
}

type VerifierDoer struct {
	mock                   *MockDoer
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierDoer) DoIt() *Doer_DoIt_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DoIt", params)
	return &Doer_DoIt_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Doer_DoIt_OngoingVerification struct {
	mock              *MockDoer
	methodInvocations []pegomock.MethodInvocation
}

func (c *Doer_DoIt_OngoingVerification) GetCapturedArguments() {
}

func (c *Doer_DoIt_OngoingVerification) GetAllCapturedArguments() {
}

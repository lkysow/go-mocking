package pegomock_test

import (
	"testing"
	. "github.com/petergtz/pegomock"
	"github.com/lkysow/go-mocking/pegomock"
	"errors"
)

func TestRunSuccess(t *testing.T) {
	RegisterMockTestingT(t)
	d := NewMockDoer()
	When(d.DoIt()).ThenReturn(nil)

	c := pegomock.User{
		D: d,
	}
	err := c.Run()
	if err != nil {
		t.Errorf("expected err to be nil, was: %v", err)
	}
}

func TestRunError(t *testing.T) {
	RegisterMockTestingT(t)
	d := NewMockDoer()
	When(d.DoIt()).ThenReturn(errors.New("err"))

	c := pegomock.User{
		D: d,
	}
	err := c.Run()
	if err.Error() != "err" {
		t.Errorf("expected error to be \"err\", was: %v", err.Error())
	}
}

package domain

import (
	"testing"
	"time"
)

func TestAccount_SetName(t *testing.T) {
	var a = NewAccount("123", "zxie", 10000, time.Now())
	if a.GetAccountName() != "zxie" {
		t.Errorf("Expected zxie but got %s", a.GetAccountName())
	}
}

func TestAccount_Withdraw(t *testing.T) {
	//GIVEN
	a := NewAccount("123", "zxie", 10000, time.Now())
	//WHEN
	err := a.Withdraw(20000)
	//THEN
	if err != nil && err.Error() != ErrInsufficientBalance.Error() {
		t.Errorf("Expected ErrInsufficientBalance but got other error %s", err.Error())
	}
	if err == nil {
		t.Errorf("Expected ErrInsufficientBalance but got no error")
	}
}

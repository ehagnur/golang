package utils_test

import (
	"examples/imports/utils"
	"testing"
)

func TestAuth(t *testing.T) {
	testLogin := utils.Login()
	if testLogin != "Login Successful" {
		t.Error("Login Failed")
	}

	pageLoadTest := utils.PageTest()
	if pageLoadTest != "Page loaded successfully" {
		t.Error("Loading home page failed")
	}
}

package controllers

import "examples/imports/utils"

//CheckLogin checks the login functionality
func CheckLogin() string {
	return utils.Login()
}

//CheckPageLoading checks page loads
func CheckPageLoading() string {
	return utils.PageTest()
}

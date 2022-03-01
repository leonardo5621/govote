package user_service

import (
	"testing"
)

func TestCreationvalidation(t *testing.T) {
	user := User{
		FirstName: "tester",
		LastName: "tester last",
		Email: "testemail",
	}
	validationErrEmail := user.Validate()
	if validationErrEmail == nil {
		t.Error("Email validation failed")
	}
	user2 := User{
		FirstName: "",
		LastName: "tester last",
		Email: "test@email.com",
	}
	validationErrName := user2.Validate()
	if validationErrName == nil {
		t.Error("First name validation failed")
	}
}

func TestGetvalidation(t *testing.T) {
	userReq := &GetUserRequest{
		UserId: "12323@$!!",
	}
	idValidation := userReq.Validate()
	if idValidation == nil {
		t.Error("Id validation failed, special characters passed")
	}
	userReqCorrect := &GetUserRequest{
		UserId: "123232wrsedSASF",
	}
	idValidationErr := userReqCorrect.Validate()
	if idValidationErr != nil {
		t.Errorf("Id validation failed, allowed characters were not recognized, %v", idValidationErr)
	}
}
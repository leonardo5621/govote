package utilities

import (
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func ReturnInternalError(err error) (error){
	return status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal error: %v", err),
	)
}

func ReturnValidationError(err error) (error){
	return status.Errorf(
		codes.InvalidArgument,
		fmt.Sprintf("Invalid parameters: %v", err),
	)
}
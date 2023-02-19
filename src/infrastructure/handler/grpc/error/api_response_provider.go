package errorGrpc

import (
	domainError "github.com/lumialvarez/go-grpc-profile-service/src/internal/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type APIResponseProvider struct {
}

func NewAPIResponseProvider() *APIResponseProvider {
	return &APIResponseProvider{}
}

func (a *APIResponseProvider) ToAPIResponse(err error) error {
	log.Print(err.Error())
	switch err.(type) {
	case domainError.AlreadyExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case domainError.NotFound:
		return status.Error(codes.NotFound, err.Error())
	case domainError.InvalidData:
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}

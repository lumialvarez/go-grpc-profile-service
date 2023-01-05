package profile

import (
	"context"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/mapper"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
)

type UseCaseList interface {
	Execute(ctx context.Context, id int64, language string) ([]*profile.Profile, error)
}

type UseCaseSave interface {
	Execute(ctx context.Context, domainUser *profile.Profile) (*profile.Profile, error)
}

type UseCaseUpdate interface {
	Execute(ctx context.Context, domainUser *profile.Profile) (*profile.Profile, error)
}

type ApiResponseProvider interface {
	ToAPIResponse(err error) error
}

type Handler struct {
	useCaseList         UseCaseList
	useCaseSave         UseCaseSave
	useCaseUpdate       UseCaseUpdate
	apiResponseProvider ApiResponseProvider
	mapper.Mapper
	pb.UnimplementedProfileServiceServer
}

func NewHandler(useCaseList UseCaseList, useCaseSave UseCaseSave, useCaseUpdate UseCaseUpdate, apiResponseProvider ApiResponseProvider) Handler {
	return Handler{useCaseList: useCaseList, useCaseSave: useCaseSave, useCaseUpdate: useCaseUpdate, apiResponseProvider: apiResponseProvider}
}

func (h *Handler) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	domainProfiles, err := h.useCaseList.Execute(ctx, request.GetProfileId(), request.GetProfileLanguage())
	if err != nil {
		return nil, h.apiResponseProvider.ToAPIResponse(err)
	}
	response := h.ToDTOListResponse(domainProfiles)
	return response, nil
}

func (h *Handler) Save(ctx context.Context, request *pb.SaveRequest) (*pb.SaveResponse, error) {
	domainProfile := h.ToDomainSaveRequest(request)
	profileCreated, err := h.useCaseSave.Execute(ctx, domainProfile)
	if err != nil {
		return nil, h.apiResponseProvider.ToAPIResponse(err)
	}
	response := h.ToDTOSaveResponse(profileCreated)
	return response, nil
}

func (h *Handler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	domainProfile := h.ToDomainUpdateRequest(request)
	profileCreated, err := h.useCaseUpdate.Execute(ctx, domainProfile)
	if err != nil {
		return nil, h.apiResponseProvider.ToAPIResponse(err)
	}
	response := h.ToDTOUpdateResponse(profileCreated)
	return response, nil
}

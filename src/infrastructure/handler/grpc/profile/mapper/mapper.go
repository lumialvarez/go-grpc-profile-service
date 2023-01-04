package mapper

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
	"time"
)

type Mapper struct {
}

func (m Mapper) ToDTOListResponse(domainProfiles []*profile.Profile) *pb.ListResponse {
	var dtoProfiles []*pb.ListResponse_Profile
	for _, domainProfile := range domainProfiles {
		domainProfileData := domainProfile.Data()

		dtoProfileData := toDtoProfileData(domainProfileData)

		dtoProfile := &pb.ListResponse_Profile{
			ProfileId:       domainProfile.Id(),
			ProfileLanguage: domainProfile.Language(),
			ProfileData:     dtoProfileData,
			LastUpdate:      domainProfile.LastUpdate().Format(time.ANSIC),
			Status:          domainProfile.Status(),
		}
		dtoProfiles = append(dtoProfiles, dtoProfile)
	}

	return &pb.ListResponse{Profile: dtoProfiles}
}

func (m Mapper) ToDomainSaveRequest(registerReq *pb.SaveRequest) *profile.Profile {
	dtoProfileData := registerReq.GetProfileData()

	domainProfileData := toDomainProfileData(dtoProfileData)

	domainProfile := profile.NewProfile(0, registerReq.ProfileLanguage, domainProfileData, time.Now(), true)

	return domainProfile
}

func (m Mapper) ToDTOSaveResponse(domainProfile *profile.Profile) *pb.SaveResponse {
	domainProfileData := domainProfile.Data()

	dtoProfileData := toDtoProfileData(domainProfileData)

	dtoProfile := pb.SaveResponse{
		ProfileId:       domainProfile.Id(),
		ProfileLanguage: domainProfile.Language(),
		ProfileData:     dtoProfileData,
		LastUpdate:      domainProfile.LastUpdate().Format(time.ANSIC),
		Status:          domainProfile.Status(),
	}

	return &dtoProfile
}

func toDomainProfileData(dtoProfileData *pb.ProfileData) profile.Data {
	dtoProjects := dtoProfileData.Projects
	dtoKnowledges := dtoProfileData.Knowledges

	var domainProjects []profile.Project
	for _, dtoProject := range dtoProjects {
		domainProject := profile.Project{
			Id:          dtoProject.Id,
			Name:        dtoProject.Name,
			Description: dtoProject.Description,
			DetailHtml:  dtoProject.DetailHtml,
			MainImage:   dtoProject.MainImage,
		}
		domainProjects = append(domainProjects, domainProject)
	}

	var domainKnowledges []profile.Knowledge
	for _, dtoKnowledge := range dtoKnowledges {
		domainKnowledge := profile.Knowledge{
			Id:          dtoKnowledge.Id,
			Name:        dtoKnowledge.Name,
			Type:        dtoKnowledge.Type,
			Level:       dtoKnowledge.Level,
			Description: dtoKnowledge.Description,
			Categories:  dtoKnowledge.Categories,
		}
		domainKnowledges = append(domainKnowledges, domainKnowledge)
	}

	profileData := profile.Data{
		Name:                dtoProfileData.Name,
		Profession:          dtoProfileData.Profession,
		ProfessionalProfile: dtoProfileData.ProfessionalProfile,
		PersonalProfile:     dtoProfileData.PersonalProfile,
		Projects:            domainProjects,
		Knowledges:          domainKnowledges,
	}
	return profileData
}

func toDtoProfileData(domainProfileData profile.Data) *pb.ProfileData {
	domainProjects := domainProfileData.Projects
	domainKnowledges := domainProfileData.Knowledges

	var dtoProjects []*pb.Project
	for _, domainProject := range domainProjects {
		dtoProject := &pb.Project{
			Id:          domainProject.Id,
			Name:        domainProject.Name,
			Description: domainProject.Description,
			DetailHtml:  domainProject.DetailHtml,
			MainImage:   domainProject.MainImage,
		}
		dtoProjects = append(dtoProjects, dtoProject)
	}

	var dtoKnowledges []*pb.Knowledge
	for _, domainKnowledge := range domainKnowledges {
		dtoKnowledge := &pb.Knowledge{
			Id:          domainKnowledge.Id,
			Name:        domainKnowledge.Name,
			Type:        domainKnowledge.Type,
			Level:       domainKnowledge.Level,
			Description: domainKnowledge.Description,
			Categories:  domainKnowledge.Categories,
		}
		dtoKnowledges = append(dtoKnowledges, dtoKnowledge)
	}

	data := domainProfileData
	dtoProfileData := &pb.ProfileData{
		Name:                data.Name,
		Profession:          data.Profession,
		ProfessionalProfile: data.ProfessionalProfile,
		PersonalProfile:     data.PersonalProfile,
		Projects:            dtoProjects,
		Knowledges:          dtoKnowledges,
	}
	return dtoProfileData
}

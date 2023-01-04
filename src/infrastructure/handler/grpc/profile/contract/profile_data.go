package contract

type ProfileData struct {
	Name                string `json:"name" binding:"required"`
	Profession          string `json:"profession" binding:"required"`
	ProfessionalProfile string `json:"professional_profile" binding:"required"`
	PersonalProfile     string `json:"personal_profile" binding:"required"`
}

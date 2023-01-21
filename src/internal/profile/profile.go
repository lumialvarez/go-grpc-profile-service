package profile

import "time"

type Profile struct {
	id         int64
	language   string
	data       Data
	lastUpdate time.Time
	status     bool
}

func NewProfile(id int64, language string, data Data, lastUpdate time.Time, status bool) *Profile {
	return &Profile{id: id, language: language, data: data, lastUpdate: lastUpdate, status: status}
}

func (p *Profile) Id() int64 {
	return p.id
}

func (p *Profile) SetId(id int64) {
	p.id = id
}

func (p *Profile) Language() string {
	return p.language
}

func (p *Profile) SetLanguage(language string) {
	p.language = language
}

func (p *Profile) Data() Data {
	return p.data
}

func (p *Profile) SetData(data Data) {
	p.data = data
}

func (p *Profile) LastUpdate() time.Time {
	return p.lastUpdate
}

func (p *Profile) SetLastUpdate(lastUpdate time.Time) {
	p.lastUpdate = lastUpdate
}

func (p *Profile) Status() bool {
	return p.status
}

func (p *Profile) SetStatus(status bool) {
	p.status = status
}

type Data struct {
	Name                string
	Profession          string
	ProfessionalProfile string
	PersonalProfile     string
	Projects            []Project
	Knowledges          []Knowledge
}

type Project struct {
	Id          string
	Name        string
	Description string
	DetailHtml  string
	MainImage   string
	Order       int32
}

type Knowledge struct {
	Id          string
	Name        string
	Type        string
	Level       int32
	Description string
	Categories  []string
}

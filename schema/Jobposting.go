package schema

import (
	"encoding/json"
)

//JobPosting is a listing that describes a job opening in a certain organization.
type JobPosting struct {
	MetaContext            string         `json:"MetaContext"`
	MetaType               string         `json:"@type"`
	BaseSalary             MonetaryAmount `json:"baseSalary,omitempty"`
	DatePosted             string         `json:"datePosted"`
	EducationRequirements  string         `json:"educationRequirements,omitempty"`
	EmploymentType         string         `json:"employmentType"`
	EstimatedSalary        MonetaryAmount `json:"estimatedSalary,omitempty"`
	ExperienceRequirements string         `json:"experienceRequirements,omitempty"`
	HiringOrganization     Organization   `json:"hiringOrganization,omitempty"`
	IncentiveCompensation  string         `json:"incentiveCompensation,omitempty"`
	Industry               string         `json:"industry,omitempty"`
	JobBenefits            string         `json:"jobBenefits,omitempty"`
	JobLocation            Place          `json:"jobLocation"`
	OccupationalCategory   string         `json:"occupationalCategory,omitempty"`
	Qualifications         string         `json:"qualifications,omitempty"`
	RelevantOccupation     Occupation     `json:"relevantOccupation,omitempty"`
	Responsibilities       string         `json:"responsibilities,omitempty"`
	SalaryCurrency         string         `json:"salaryCurrency,omitempty"`
	Skills                 string         `json:"skills,omitempty"`
	SpecialCommitments     string         `json:"specialCommitments,omitempty"`
	Title                  string         `json:"title"`
	ValidThrough           string         `json:"validThrough,omitempty"`
	WorkHours              string         `json:"workHours,omitempty"`
}

//Occupation is a profession, may involve prolonged training and/or a formal qualification.
type Occupation struct {
	MetaType               string             `json:"@type"`
	EducationRequirements  string             `json:"educationRequirements,omitempty"`
	EstimatedSalary        MonetaryAmount     `json:"estimatedSalary,omitempty"`
	ExperienceRequirements string             `json:"experienceRequirements,omitempty"`
	OccupationLocation     AdministrativeArea `json:"occupationLocation,omitempty"`
	OccupationalCategory   string             `json:"occupationalCategory,omitempty"`
	Qualifications         string             `json:"qualifications,omitempty"`
	Responsibilities       string             `json:"responsibilities,omitempty"`
	Skills                 string             `json:"skills,omitempty"`
}

//MonetaryAmount is a monetary value or range
type MonetaryAmount struct {
	MetaType     string `json:"@type"`
	Currency     string `json:"currency,omitempty"`
	MaxValue     int64  `json:"maxValue,omitempty"`
	MinValue     int64  `json:"minValue,omitempty"`
	ValidFrom    string `json:"validFrom,omitempty"`
	ValidThrough string `json:"validThrough,omitempty"`
	Value        string `json:"value,omitempty"`
}

//NewJobPosting returns a new instance of JobPosting with compulsory attributes set
func NewJobPosting(datePosted, employmentType, title, jobLocation string) JobPosting {
	return JobPosting{
		MetaContext:    context,
		MetaType:       "JobPosting",
		DatePosted:     datePosted,
		EmploymentType: employmentType,
		Title:          title,
		JobLocation: Place{
			Address: jobLocation,
		},
	}
}

func (jp *JobPosting) String() (string, error) {
	b, err := json.Marshal(jp)
	return string(b), err
}

package schema

//Jobposting is a listing that describes a job opening in a certain organization.
type Jobposting struct {
	MetaContext            string         `json:"MetaContext"`
	MetaType               string         `json:"@type"`
	BaseSalary             MonetaryAmount `json:"baseSalary"`
	DatePosted             string         `json:"datePosted"`
	EducationRequirements  string         `json:"educationRequirements"`
	EmploymentType         string         `json:"employmentType"`
	EstimatedSalary        MonetaryAmount `json:"estimatedSalary"`
	ExperienceRequirements string         `json:"experienceRequirements"`
	HiringOrganization     Organization   `json:"hiringOrganization"`
	IncentiveCompensation  string         `json:"incentiveCompensation"`
	Industry               string         `json:"industry"`
	JobBenefits            string         `json:"jobBenefits"`
	JobLocation            Place          `json:"jobLocation"`
	OccupationalCategory   string         `json:"occupationalCategory"`
	Qualifications         string         `json:"qualifications"`
	RelevantOccupation     Occupation     `json:"relevantOccupation"`
	Responsibilities       string         `json:"responsibilities"`
	SalaryCurrency         string         `json:"salaryCurrency"`
	Skills                 string         `json:"skills"`
	SpecialCommitments     string         `json:"specialCommitments"`
	Title                  string         `json:"title"`
	ValidThrough           string         `json:"validThrough"`
	WorkHours              string         `json:"workHours"`
}

//Occupation is a profession, may involve prolonged training and/or a formal qualification.
type Occupation struct {
	MetaType               string             `json:"@type"`
	EducationRequirements  string             `json:"educationRequirements"`
	EstimatedSalary        MonetaryAmount     `json:"estimatedSalary"`
	ExperienceRequirements string             `json:"experienceRequirements"`
	OccupationLocation     AdministrativeArea `json:"occupationLocation"`
	OccupationalCategory   string             `json:"occupationalCategory"`
	Qualifications         string             `json:"qualifications"`
	Responsibilities       string             `json:"responsibilities"`
	Skills                 string             `json:"skills"`
}

//MonetaryAmount is a monetary value or range
type MonetaryAmount struct {
	MetaType     string `json:"@type"`
	Currency     string `json:"currency"`
	MaxValue     int64  `json:"maxValue"`
	MinValue     int64  `json:"minValue"`
	ValidFrom    string `json:"validFrom"`
	ValidThrough string `json:"validThrough"`
	Value        string `json:"value"`
}

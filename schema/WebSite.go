package schema

// WebSite is a set of related web pages and other items typically served from a single web domain and accessible via URLs.
type WebSite struct {
	MetaContext string       `json:"@context"`
	MetaType    string       `json:"@type"`
	URL         string       `json:"url"`
	Issn        string       `json:"issn"`
	Creator     Organization `json:"creator"`
	InLanguage  string       `json:"inLanguage"`
	Keywords    string       `json:"keywords"`
	License     License      `json:"license"`
	Award       string       `json:"award"`
	Audience    Audience     `json:"audience"`
}

//License is a subset of CreativeWork
type License struct {
	Type            string        `json:"@type"`
	Context         string        `json:"@context"`
	CopyrightHolder OwnershipInfo `json:"copyrightHolder"`
	CopyrightYear   int           `json:"copyrightYear"`
}

//OwnershipInfo is a structured value providing information about when a certain organization or person owned a certain product.
type OwnershipInfo struct {
	Type               string       `json:"@type"`
	TypeOfGood         string       `json:"typeOfGood"`
	OwnedFrom          string       `json:"ownedFrom"`
	AcquiredFromOrg    Organization `json:"acquiredFromOrg"`
	AcquiredFromPerson Person       `json:"acquiredFromPerson"`
}

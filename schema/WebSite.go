package schema

import "encoding/json"

// WebSite is a set of related web pages and other items typically served from a single web domain and accessible via URLs.
type WebSite struct {
	MetaContext string       `json:"@context"`
	MetaType    string       `json:"@type"`
	Name        string       `json:"name"`
	URL         string       `json:"url"`
	Issn        string       `json:"issn,omitempty"`
	Creator     Organization `json:"creator,omitempty"`
	InLanguage  string       `json:"inLanguage,omitempty"`
	Keywords    string       `json:"keywords,omitempty"`
	License     License      `json:"license,omitempty"`
	Audience    Audience     `json:"audience,omitempty"`
}

//License is a subset of CreativeWork
type License struct {
	MetaType        string        `json:"@type"`
	CopyrightHolder OwnershipInfo `json:"copyrightHolder,omitempty"`
	CopyrightYear   int           `json:"copyrightYear,omitempty"`
}

//OwnershipInfo is a structured value providing information about when a certain organization or person owned a certain product.
type OwnershipInfo struct {
	MetaType           string       `json:"@type"`
	TypeOfGood         string       `json:"typeOfGood,omitempty"`
	OwnedFrom          string       `json:"ownedFrom,omitempty"`
	AcquiredFromOrg    Organization `json:"acquiredFromOrg,omitempty"`
	AcquiredFromPerson Person       `json:"acquiredFromPerson,omitempty"`
}

//NewWebSite returns a new instance of WebSite
func NewWebSite(name, url, issn, inLanguage, keywords string) WebSite {
	return WebSite{
		MetaContext: context,
		MetaType:    "WebSite",
		Name:        name,
		URL:         url,
		Issn:        issn,
		InLanguage:  inLanguage,
	}
}

//SetOrganization is a helper function
func (ws *WebPage) SetOrganization() {

}

func (ws *WebSite) String() string {
	b, _ := json.Marshal(ws)
	return string(b)
}

package schema

import (
	"encoding/json"
	"strings"
)

//MetaLocalBusiness is the metatype tag for LocalBusiness
const MetaLocalBusiness string = "LocalBusiness"

//LocalBusiness is a particular physical business or branch of an organization.
type LocalBusiness struct {
	MetaContext        string        `json:"MetaContext"`
	MetaType           string        `json:"@type"`
	Address            PostalAddress `json:"address,omitempty"`
	CurrenciesAccepted string        `json:"currenciesAccepted,omitempty"`
	Description        string        `json:"description,omitempty"`
	FaxNumber          string        `json:"faxNumber,omitempty"`
	Founder            Person        `json:"founder,omitempty"`
	FoundingDate       string        `json:"foundingDate,omitempty"`
	Image              string        `json:"Image,omitempty"`
	OpeningHours       string        `json:"openingHours,omitempty"`
	PaymentAccepted    string        `json:"paymentAccepted,omitempty"`
	PriceRange         string        `json:"priceRange,omitempty"`
	Name               string        `json:"name,omitempty"`
	TaxID              string        `json:"taxID,omitempty"`
	Telephone          string        `json:"telephone,omitempty"`
	VatID              string        `json:"vatID,omitempty"`
	URL                string        `json:"url,omitempty"`
}

//NewLocalBusiness returns a new instance of LocalBusiness
func NewLocalBusiness(name, image, url, streetaddress, locality, phone string, openinghours []string) LocalBusiness {
	return LocalBusiness{
		MetaContext: MetaContext,
		MetaType:    MetaLocalBusiness,
		Name:        name,
		Image:       image,
		URL:         url,
		Address: PostalAddress{
			MetaType:        MetaPostalAddress,
			StreetAddress:   streetaddress,
			AddressLocality: locality,
		},
		Telephone:    phone,
		OpeningHours: strings.Join(openinghours, ", "),
	}
}

func (lb LocalBusiness) String() string {
	b, _ := json.Marshal(lb)
	return string(b)
}

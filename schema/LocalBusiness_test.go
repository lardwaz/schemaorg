package schema

import (
	"strings"
	"testing"
)

func TestLocalBusiness_String(t *testing.T) {
	type fields struct {
		MetaContext        string
		MetaType           string
		Address            PostalAddress
		CurrenciesAccepted string
		Description        string
		FaxNumber          string
		Founder            Person
		FoundingDate       string
		Image              string
		OpeningHours       string
		PaymentAccepted    string
		PriceRange         string
		Name               string
		TaxID              string
		Telephone          string
		VatID              string
		URL                string
	}

	openingHrs := []string{
		"Mo-Sa 11:00-14:30",
		"Mo-Th 17:00-21:30",
		"Fr-Sa 17:00-22:00",
	}

	case1 := NewLocalBusiness("ABCD Alpha Ltd", "https://pixabay.com/en/urban-people-crowd-citizens-438393/",
		"http://localhost:8080/localbusiness/", "27 Ketch Harbour Street", "Mexico Beach", "(230) 499-9999", openingHrs)

	case2 := LocalBusiness{
		MetaContext:        context,
		MetaType:           MetaLocalBusiness,
		Address:            PostalAddress{},
		CurrenciesAccepted: "USD, MUR, INR",
		Description:        "We specialise in dealing with fruits",
		FaxNumber:          "499-8989",
		Founder: Person{
			GivenName:  "James",
			FamilyName: "Smith",
		},
		FoundingDate:    "2000-05-06",
		Image:           "https://pixabay.com/en/urban-people-crowd-citizens-438393/",
		OpeningHours:    strings.Join(openingHrs, ", "),
		PaymentAccepted: "Cash, Cheque, Bitcoin",
		PriceRange:      "20 USD - 3000 USD",
		Name:            "Romeo LTD",
		Telephone:       "5999-9999",
		URL:             "http://localhost:8080/romeo/",
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext: context,
				MetaType:    MetaLocalBusiness,
				Name:        "ABCD Alpha Ltd",
				Image:       "https://pixabay.com/en/urban-people-crowd-citizens-438393/",
				URL:         "http://localhost:8080/localbusiness/",
				Address: PostalAddress{
					MetaType:        MetaPostalAddress,
					AddressLocality: "Mexico Beach",
					StreetAddress:   "27 Ketch Harbour Street",
				},
				Telephone:    "(230) 499-9999",
				OpeningHours: strings.Join(openingHrs, ", "),
			},
			want: case1.String(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				MetaContext:        context,
				MetaType:           MetaLocalBusiness,
				Address:            PostalAddress{},
				CurrenciesAccepted: "USD, MUR, INR",
				Description:        "We specialise in dealing with fruits",
				FaxNumber:          "499-8989",
				Founder: Person{
					GivenName:  "James",
					FamilyName: "Smith",
				},
				FoundingDate:    "2000-05-06",
				Image:           "https://pixabay.com/en/urban-people-crowd-citizens-438393/",
				OpeningHours:    strings.Join(openingHrs, ", "),
				PaymentAccepted: "Cash, Cheque, Bitcoin",
				PriceRange:      "20 USD - 3000 USD",
				Name:            "Romeo LTD",
				Telephone:       "5999-9999",
				URL:             "http://localhost:8080/romeo/",
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := LocalBusiness{
				MetaContext:        tt.fields.MetaContext,
				MetaType:           tt.fields.MetaType,
				Address:            tt.fields.Address,
				CurrenciesAccepted: tt.fields.CurrenciesAccepted,
				Description:        tt.fields.Description,
				FaxNumber:          tt.fields.FaxNumber,
				Founder:            tt.fields.Founder,
				FoundingDate:       tt.fields.FoundingDate,
				Image:              tt.fields.Image,
				OpeningHours:       tt.fields.OpeningHours,
				PaymentAccepted:    tt.fields.PaymentAccepted,
				PriceRange:         tt.fields.PriceRange,
				Name:               tt.fields.Name,
				TaxID:              tt.fields.TaxID,
				Telephone:          tt.fields.Telephone,
				VatID:              tt.fields.VatID,
				URL:                tt.fields.URL,
			}
			if got := lb.String(); got != tt.want {
				t.Errorf("LocalBusiness.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

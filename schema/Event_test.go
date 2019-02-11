package schema

import "testing"

func TestEvent_String(t *testing.T) {
	type fields struct {
		MetaContext string
		MetaType    string
		Name        string
		StartDate   string
		EndDate     string
		Image       string
		URL         string
		Organizer   string
		Description string
		Location    EventLocation
	}

	case1 := Event{
		MetaContext: context,
		MetaType:    "Event",
		Name:        "Tuesday Sushi",
		StartDate:   "Janvier 29 (Mardi) 11:00 Am",
		EndDate:     "",
		Image:       "https://i0.wp.com/koze.mu/wp-content/uploads/2019/01/48356257_1949939745123408_1613789933762248704_o.jpg",
		URL:         "https://koze.mu/events/tuesday-sushi/",
		Organizer:   "Koze.mu",
		Description: "Embark on a culinary trip with sushi, sashimi, maki, temaki and our specialties, along with the indispensable wasabi and soy sauce, as your companions for any tasty Tuesday journey.",
		Location: EventLocation{
			MetaType:  MetaForLocation,
			Name:      "The Address Boutique Hotel",
			Telephone: "405 30 00",
			Address:   "The Address Boutique Hotel, Port Chambly, Balaclava",
		},
	}

	case2 := NewEvent("Le Market", "https://i1.wp.com/koze.mu/wp-content/uploads/2019/01/41408192_1204721073002153_3691893886948474880_o.jpg", "https://koze.mu/events/le-market/", "Mont Choisy Le Mall, Mont Choisy, 742CU001, Grand Baie, Riviere Du Rempart, Mauritius", "57 33 95 57")

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext: context,
				MetaType:    "Event",
				Name:        "Tuesday Sushi",
				StartDate:   "Janvier 29 (Mardi) 11:00 Am",
				EndDate:     "",
				Image:       "https://i0.wp.com/koze.mu/wp-content/uploads/2019/01/48356257_1949939745123408_1613789933762248704_o.jpg",
				URL:         "https://koze.mu/events/tuesday-sushi/",
				Organizer:   "Koze.mu",
				Description: "Embark on a culinary trip with sushi, sashimi, maki, temaki and our specialties, along with the indispensable wasabi and soy sauce, as your companions for any tasty Tuesday journey.",
				Location: EventLocation{
					MetaType:  MetaForLocation,
					Name:      "The Address Boutique Hotel",
					Telephone: "405 30 00",
					Address:   "The Address Boutique Hotel, Port Chambly, Balaclava",
				},
			},
			want: case1.String(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				MetaContext: context,
				MetaType:    "Event",
				Name:        "Le Market",
				Image:       "https://i1.wp.com/koze.mu/wp-content/uploads/2019/01/41408192_1204721073002153_3691893886948474880_o.jpg",
				URL:         "https://koze.mu/events/le-market/",
				Location: EventLocation{
					MetaType:  MetaForLocation,
					Telephone: "57 33 95 57",
					Address:   "Mont Choisy Le Mall, Mont Choisy, 742CU001, Grand Baie, Riviere Du Rempart, Mauritius",
				},
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Event{
				MetaContext: tt.fields.MetaContext,
				MetaType:    tt.fields.MetaType,
				Name:        tt.fields.Name,
				StartDate:   tt.fields.StartDate,
				EndDate:     tt.fields.EndDate,
				Image:       tt.fields.Image,
				URL:         tt.fields.URL,
				Organizer:   tt.fields.Organizer,
				Description: tt.fields.Description,
				Location:    tt.fields.Location,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Event.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

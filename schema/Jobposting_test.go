package schema

import "testing"

func TestJobPosting_String(t *testing.T) {
	type fields struct {
		MetaContext            string
		MetaType               string
		BaseSalary             MonetaryAmount
		DatePosted             string
		EducationRequirements  string
		EmploymentType         string
		EstimatedSalary        MonetaryAmount
		ExperienceRequirements string
		HiringOrganization     Organization
		IncentiveCompensation  string
		Industry               string
		JobBenefits            string
		JobLocation            Place
		OccupationalCategory   string
		Qualifications         string
		RelevantOccupation     Occupation
		Responsibilities       string
		SalaryCurrency         string
		Skills                 string
		SpecialCommitments     string
		Title                  string
		ValidThrough           string
		WorkHours              string
	}

	case1 := NewJobPosting("2019-01-03", "Full-Time", "Software Developer", "Ceridian Learning Center, Quatres Bornes")

	case2 := JobPosting{
		MetaContext: context,
		MetaType:    "JobPosting",
		BaseSalary: MonetaryAmount{
			MetaType: MetaMonetaryAmount,
			Value:    "Rs 25000",
		},
		DatePosted:             "2019-02-10",
		EducationRequirements:  "BSc (Hons) Computer Science",
		EmploymentType:         "Permanent",
		ExperienceRequirements: "2 yrs",
		Industry:               "Computing",
		JobLocation: Place{
			MetaType: MetaPlace,
			Address:  "Ceridian Learning Center, Quatres Bornes",
		},
		OccupationalCategory: "Devops",
		Title:                "Network Engineer",
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext:    context,
				MetaType:       "JobPosting",
				DatePosted:     "2019-01-03",
				EmploymentType: "Full-Time",
				Title:          "Software Developer",
				JobLocation: Place{
					Address: "Ceridian Learning Center, Quatres Bornes",
				},
			},
			want: case1.String(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				MetaContext: context,
				MetaType:    "JobPosting",
				BaseSalary: MonetaryAmount{
					MetaType: MetaMonetaryAmount,
					Value:    "Rs 25000",
				},
				DatePosted:             "2019-02-10",
				EducationRequirements:  "BSc (Hons) Computer Science",
				EmploymentType:         "Permanent",
				ExperienceRequirements: "2 yrs",
				Industry:               "Computing",
				JobLocation: Place{
					MetaType: MetaPlace,
					Address:  "Ceridian Learning Center, Quatres Bornes",
				},
				OccupationalCategory: "Devops",
				Title:                "Network Engineer",
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jp := JobPosting{
				MetaContext:            tt.fields.MetaContext,
				MetaType:               tt.fields.MetaType,
				BaseSalary:             tt.fields.BaseSalary,
				DatePosted:             tt.fields.DatePosted,
				EducationRequirements:  tt.fields.EducationRequirements,
				EmploymentType:         tt.fields.EmploymentType,
				EstimatedSalary:        tt.fields.EstimatedSalary,
				ExperienceRequirements: tt.fields.ExperienceRequirements,
				HiringOrganization:     tt.fields.HiringOrganization,
				IncentiveCompensation:  tt.fields.IncentiveCompensation,
				Industry:               tt.fields.Industry,
				JobBenefits:            tt.fields.JobBenefits,
				JobLocation:            tt.fields.JobLocation,
				OccupationalCategory:   tt.fields.OccupationalCategory,
				Qualifications:         tt.fields.Qualifications,
				RelevantOccupation:     tt.fields.RelevantOccupation,
				Responsibilities:       tt.fields.Responsibilities,
				SalaryCurrency:         tt.fields.SalaryCurrency,
				Skills:                 tt.fields.Skills,
				SpecialCommitments:     tt.fields.SpecialCommitments,
				Title:                  tt.fields.Title,
				ValidThrough:           tt.fields.ValidThrough,
				WorkHours:              tt.fields.WorkHours,
			}
			if got := jp.String(); got != tt.want {
				t.Errorf("JobPosting.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

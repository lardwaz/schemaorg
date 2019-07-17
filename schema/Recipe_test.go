package schema_test

import (
	"testing"

	"go.lsl.digital/lardwaz/schemaorg/schema"
)

func TestRecipe_String(t *testing.T) {
	type fields struct {
		MetaContext        string
		MetaType           string
		Name               string
		RecipeCuisine      string
		RecipeCategory     string
		Description        string
		Image              string
		Video              string
		Ingredients        []string
		RecipeInstructions []string
		RecipeYield        string
		Nutrition          schema.NutritionInformation
		CookTime           string
		PrepTime           string
		TotalTime          string
		DatePublished      string
		URL                string
	}

	ingredients := []string{
		"4	EGGS",
		"1/4 cup	milk",
		"salt and pepper, as desired",
		"2 tsp.	butter",
	}

	recipeInstructions := []string{
		"BEAT eggs, milk, salt and pepper in medium bowl until blended.",
		"HEAT butter in large nonstick skillet over medium heat until hot. POUR IN egg mixture. As eggs begin to set, GENTLY PULL the eggs across the pan with a spatula, forming large soft curds.",
		"CONTINUE cooking – pulling, lifting and folding eggs – until thickened and no visible liquid egg remains. Do not stir constantly. REMOVE from heat. SERVE immediately.",
	}

	case1 := schema.NewRecipe("Scrambled eggs", "https://x9wsr1khhgk5pxnq1f1r8kye-wpengine.netdna-ssl.com/wp-content/uploads/Scrambled-with-Milk-930x620.jpg",
		"https://www.incredibleegg.org/recipe/basic-scrambled-eggs/", ingredients, recipeInstructions)

	case2 := schema.Recipe{
		MetaContext:        schema.MetaContext,
		MetaType:           schema.MetaRecipe,
		Name:               "Basic Scrambled Eggs Recipe",
		RecipeCuisine:      "English",
		RecipeCategory:     "Breakfast",
		Description:        "Basic Scrambled Eggs Recipe is prepared by beating them with a little liquid and then cooking and stirring gently.",
		Image:              "https://x9wsr1khhgk5pxnq1f1r8kye-wpengine.netdna-ssl.com/wp-content/uploads/Scrambled-with-Milk-930x620.jpg",
		Ingredients:        ingredients,
		RecipeInstructions: recipeInstructions,
		RecipeYield:        "2 servings",
		Nutrition: schema.NutritionInformation{
			MetaType:           schema.MetaNutritionInformation,
			Calories:           "320 calories",
			CholesterolContent: "3%",
		},
		CookTime:      "4m",
		PrepTime:      "1m",
		TotalTime:     "5m",
		DatePublished: "2019-02-06",
		URL:           "https://www.incredibleegg.org/recipe/basic-scrambled-eggs/",
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext:        schema.MetaContext,
				MetaType:           schema.MetaRecipe,
				Name:               "Scrambled eggs",
				Image:              "https://x9wsr1khhgk5pxnq1f1r8kye-wpengine.netdna-ssl.com/wp-content/uploads/Scrambled-with-Milk-930x620.jpg",
				URL:                "https://www.incredibleegg.org/recipe/basic-scrambled-eggs/",
				Ingredients:        ingredients,
				RecipeInstructions: recipeInstructions,
			},
			want: case1.String(),
		},
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext:        schema.MetaContext,
				MetaType:           schema.MetaRecipe,
				Name:               "Basic Scrambled Eggs Recipe",
				RecipeCuisine:      "English",
				RecipeCategory:     "Breakfast",
				Description:        "Basic Scrambled Eggs Recipe is prepared by beating them with a little liquid and then cooking and stirring gently.",
				Image:              "https://x9wsr1khhgk5pxnq1f1r8kye-wpengine.netdna-ssl.com/wp-content/uploads/Scrambled-with-Milk-930x620.jpg",
				Ingredients:        ingredients,
				RecipeInstructions: recipeInstructions,
				RecipeYield:        "2 servings",
				Nutrition: schema.NutritionInformation{
					MetaType:           schema.MetaNutritionInformation,
					Calories:           "320 calories",
					CholesterolContent: "3%",
				},
				CookTime:      "4m",
				PrepTime:      "1m",
				TotalTime:     "5m",
				DatePublished: "2019-02-06",
				URL:           "https://www.incredibleegg.org/recipe/basic-scrambled-eggs/",
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := schema.Recipe{
				MetaContext:        tt.fields.MetaContext,
				MetaType:           tt.fields.MetaType,
				Name:               tt.fields.Name,
				RecipeCuisine:      tt.fields.RecipeCuisine,
				RecipeCategory:     tt.fields.RecipeCategory,
				Description:        tt.fields.Description,
				Image:              tt.fields.Image,
				Video:              tt.fields.Video,
				Ingredients:        tt.fields.Ingredients,
				RecipeInstructions: tt.fields.RecipeInstructions,
				RecipeYield:        tt.fields.RecipeYield,
				Nutrition:          tt.fields.Nutrition,
				CookTime:           tt.fields.CookTime,
				PrepTime:           tt.fields.PrepTime,
				TotalTime:          tt.fields.TotalTime,
				DatePublished:      tt.fields.DatePublished,
				URL:                tt.fields.URL,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Recipe.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

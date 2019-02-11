package schema

import "encoding/json"

//Recipe is a set of instructions for preparing a particular dish, including a list of the ingredients required.
type Recipe struct {
	MetaContext        string               `json:"@context"`
	MetaType           string               `json:"@type"`
	Name               string               `json:"name"`
	RecipeCuisine      string               `json:"recipeCuisine,omitempty"`
	RecipeCategory     string               `json:"recipeCategory,omitempty"`
	Description        string               `json:"description"`
	Image              string               `json:"image,omitempty"`
	Video              string               `json:"video,omitempty"`
	Ingredients        []string             `json:"ingredients"`
	RecipeInstructions []string             `json:"recipeInstructions"`
	RecipeYield        string               `json:"recipeYield,omitempty"`
	Nutrition          NutritionInformation `json:"nutrition,omitempty"`
	CookTime           string               `json:"cookTime,omitempty"`
	PrepTime           string               `json:"prepTime,omitempty"`
	TotalTime          string               `json:"totalTime,omitempty"`
	DatePublished      string               `json:"datePublished,omitempty"`
	URL                string               `json:"url"`
}

//NutritionInformation about the recipe.
type NutritionInformation struct {
	MetaType              string `json:"@type"`
	Calories              string `json:"calories,omitempty"`
	CarbohydrateContent   string `json:"carbohydrateContent,omitempty"`
	CholesterolContent    string `json:"cholesterolContent,omitempty"`
	FatContent            string `json:"fatContent,omitempty"`
	FiberContent          string `json:"fiberContent,omitempty"`
	ProteinContent        string `json:"proteinContent,omitempty"`
	SaturatedFatContent   string `json:"saturatedFatContent,omitempty"`
	ServingSize           string `json:"servingSize,omitempty"`
	SodiumContent         string `json:"sodiumContent,omitempty"`
	SugarContent          string `json:"sugarContent,omitempty"`
	TransFatContent       string `json:"transFatContent,omitempty"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent,omitempty"`
}

//NewRecipe returns a new instance of Recipe with compulsory attributes set
func NewRecipe(name, image, url string, ingredients, recipeInstructions []string) Recipe {
	return Recipe{
		MetaContext:        context,
		MetaType:           "Recipe",
		Name:               name,
		URL:                url,
		Ingredients:        ingredients,
		RecipeInstructions: recipeInstructions,
	}
}

func (r *Recipe) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

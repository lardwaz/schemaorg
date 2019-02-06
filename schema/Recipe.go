package schema

//Recipe is a set of instructions for preparing a particular dish, including a list of the ingredients required.
type Recipe struct {
	MetaContext        string               `json:"MetaContext"`
	MetaType           string               `json:"@type"`
	CookTime           string               `json:"cookTime"`
	CookingMethod      string               `json:"cookingMethod"`
	Nutrition          NutritionInformation `json:"nutrition"`
	RecipeCategory     string               `json:"recipeCategory"`
	RecipeCuisine      string               `json:"recipeCuisine"`
	RecipeIngredient   string               `json:"recipeIngredient"`
	RecipeInstructions string               `json:"recipeInstructions"`
	RecipeYield        string               `json:"recipeYield"`
	SuitableForDiet    string               `json:"suitableForDiet"`
}

//NutritionInformation about the recipe.
type NutritionInformation struct {
	MetaType              string `json:"@type"`
	Calories              string `json:"calories"`
	CarbohydrateContent   string `json:"carbohydrateContent"`
	CholesterolContent    string `json:"cholesterolContent"`
	FatContent            string `json:"fatContent"`
	FiberContent          string `json:"fiberContent"`
	ProteinContent        string `json:"proteinContent"`
	SaturatedFatContent   string `json:"saturatedFatContent"`
	ServingSize           string `json:"servingSize"`
	SodiumContent         string `json:"sodiumContent"`
	SugarContent          string `json:"sugarContent"`
	TransFatContent       string `json:"transFatContent"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent"`
}

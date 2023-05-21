package db

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RecipeRepository interface {
	GetRecipesFromSpoonacular(number int) ([]Recipe, error)
}

type recipeRepositoryImpl struct{}

func NewRecipeRepository() *recipeRepositoryImpl {
	return &recipeRepositoryImpl{}
}

type Recipe struct {
	ID           int64              `json:"id"`
	Aisle        string             `json:"aisle"`
	Image        string             `json:"image"`
	Consistency  string             `json:"consistency"`
	Name         string             `json:"name"`
	NameClean    string             `json:"nameClean"`
	Original     string             `json:"original"`
	OriginalName string             `json:"originalName"`
	Amount       float64            `json:"amount"`
	Unit         string             `json:"unit"`
	Meta         []string           `json:"meta"`
	Measures     map[string]Measure `json:"measures"`
}

type Measure struct {
	Amount    float64 `json:"amount"`
	UnitShort string  `json:"unitShort"`
	UnitLong  string  `json:"unitLong"`
}

type SpoonacularResponse struct {
	Recipes []SpoonacularRecipe `json:"recipes"`
}

type SpoonacularRecipe struct {
	ID                  int64                `json:"id"`
	Aisle               string               `json:"aisle"`
	Image               string               `json:"image"`
	Consistency         string               `json:"consistency"`
	Name                string               `json:"name"`
	NameClean           string               `json:"nameClean"`
	Original            string               `json:"original"`
	OriginalName        string               `json:"originalName"`
	Amount              float64              `json:"amount"`
	Unit                string               `json:"unit"`
	Meta                []string             `json:"meta"`
	ExtendedIngredients []ExtendedIngredient `json:"extendedIngredients"`
}

type ExtendedIngredient struct {
	ID           int64              `json:"id"`
	Aisle        string             `json:"aisle"`
	Image        string             `json:"image"`
	Consistency  string             `json:"consistency"`
	Name         string             `json:"name"`
	NameClean    string             `json:"nameClean"`
	Original     string             `json:"original"`
	OriginalName string             `json:"originalName"`
	Amount       float64            `json:"amount"`
	Unit         string             `json:"unit"`
	Meta         []string           `json:"meta"`
	Measures     map[string]Measure `json:"measures"`
}

func (*recipeRepositoryImpl) GetRecipesFromSpoonacular(number int) ([]Recipe, error) {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/random?number=%d", number)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request recipes to Spoonacular API: %v", err)
	}

	req.Header.Set("x-api-key", APIKey)
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error requesting recipes from Spoonacular API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Spoonacular API returned non-OK status code: %d", resp.StatusCode)
	}

	var spoonacularResponse SpoonacularResponse
	err = json.NewDecoder(resp.Body).Decode(&spoonacularResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding Spoonacular response: %v", err)
	}

	recipes := make([]Recipe, len(spoonacularResponse.Recipes))
	for i, spoonacularRecipe := range spoonacularResponse.Recipes {
		ingredients := make(map[string]Measure)
		for _, ingredient := range spoonacularRecipe.ExtendedIngredients {
			ingredients[ingredient.Name] = Measure{
				Amount:    ingredient.Amount,
				UnitShort: ingredient.Unit,
				UnitLong:  ingredient.Unit,
			}
		}

		recipe := Recipe{
			ID:           spoonacularRecipe.ID,
			Aisle:        spoonacularRecipe.Aisle,
			Image:        spoonacularRecipe.Image,
			Consistency:  spoonacularRecipe.Consistency,
			Name:         spoonacularRecipe.Name,
			NameClean:    spoonacularRecipe.NameClean,
			Original:     spoonacularRecipe.Original,
			OriginalName: spoonacularRecipe.OriginalName,
			Amount:       spoonacularRecipe.Amount,
			Unit:         spoonacularRecipe.Unit,
			Meta:         spoonacularRecipe.Meta,
			Measures:     ingredients,
		}
		recipes[i] = recipe
	}

	// Afficher les données JSON indentées
	indentedData, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling recipes data: %v", err)
	}

	// Afficher les données dans la sortie standard
	fmt.Println(string(indentedData))

	return recipes, nil
}

package main

//Lets make a Cook

type Cook struct {
	recipe RecipeMaker
}

//Recipe

type Recipe struct {
	name string
	ingredients string
	timeTaken int
}

//RecipeMaker interface

type RecipeMaker interface {
	SetName()
	SetIngredients()
	GetRecipe()
}

//Let's make Espresso

type Expresso struct {
	Recipe
}

func (e *Expresso) SetName() RecipeMaker {
	e.Recipe.name = "Expresso"
	return e
}


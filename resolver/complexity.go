package resolver

const (
	singleResourceCost   = 25
	multipleResourceCost = 50
)

func ComplexityLimit() int {
	return (singleResourceCost * 2) + (multipleResourceCost * 2)
}

func NewComplexityRoot() ComplexityRoot {
	c := ComplexityRoot{}

	c.Query.Film = singleResourceComplexity
	c.Query.Character = singleResourceComplexity
	c.Film.Characters = multipleResourceComplexity
	c.Person.Films = multipleResourceComplexity

	return c
}

func singleResourceComplexity(childComplexity int, id string) int {
	return singleResourceCost + childComplexity
}

func multipleResourceComplexity(childComplexity int) int {
	return multipleResourceCost + childComplexity
}

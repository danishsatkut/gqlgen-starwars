package resolver

const singleResourceCost = 25
const multipleResourceCost = 50

func NewComplexityRoot() ComplexityRoot {
	c := ComplexityRoot{}

	c.Query.Film = func(childComplexity int, id string) int {
		return singleResourceCost + childComplexity
	}

	c.Query.Character = func(childComplexity int, id string) int {
		return singleResourceCost + childComplexity
	}

	c.Film.Characters = func(childComplexity int) int {
		return multipleResourceCost + childComplexity
	}

	c.Person.Films = func(childComplexity int) int {
		return multipleResourceCost + childComplexity
	}

	return c
}

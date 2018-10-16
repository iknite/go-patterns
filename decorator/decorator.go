/*
	decorator pattern understanded as a specilzation around an object, can
	be done in golang through struct initialization.
*/
package decorator

import (
	"fmt"
	"math/rand"
)

// Person is the struct to be decorated.
type Person struct {
	title string
	name  string

	facts []string
}

// Fact returns a random or concrete tip around person. It's used to illustrate
// that baseClass function can be accessed transparent in decorated objects.
func (f *Person) fact(i int) string {
	if i == -1 {
		i = rand.Intn(len(f.facts))
	}
	return f.facts[i]
}

// Stringer implementation, we use it as the overridable method.
func (p *Person) String() string {
	return fmt.Sprintf("%s %s %s", p.title, p.name, p.fact(0))
}

// EarthPresident is a decorator around Person.
type EarthPresident struct {
	Person
}

// Stringer implemantation, the specialized method.
// With access to the base method through the promoted field `Person`
func (s *EarthPresident) String() string {
	return fmt.Sprintf(
		"The Chief Executive Officer of the human race, %s",
		s.Person.String(),
	)
}

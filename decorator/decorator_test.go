package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecotaror(t *testing.T) {
	p := Person{"Doctor", "Who",
		[]string{
			"has a swift tardis",
			"fights like a rockstar",
			"makes daleks crazy",
		},
	}

	assert.Equal(t, p.fact(2), "makes daleks crazy")
	assert.Equal(t, p.String(), "Doctor Who has a swift tardis")

	ceo := EarthPresident{p}

	// still accessing the lower methods of Person struct.
	assert.Equal(t, ceo.fact(2), "makes daleks crazy")

	// It decorates over the struct to compose more especialized methods.
	assert.Equal(t,
		ceo.String(),
		"The Chief Executive Officer of the human race, Doctor Who has a swift tardis")
}

package demolib_test

import (
	"testing"

	"github.com/fmelosilva/demolib"
	"github.com/stretchr/testify/require"
)

func TestDemoLibCreation(t *testing.T) {
	assert := require.New(t)

	person := demolib.NewPerson("john")

	assert.Equal("john", person.GetName())
}

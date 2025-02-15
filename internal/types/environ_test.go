package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnviron(t *testing.T) {

	environ, err := Environ(Spec{
		Envfile: Envfile{"testdata/spec.env"},
		Env: EnvVars{
			"BAR": "2",
		},
	}, Task{
		Envfile: Envfile{"testdata/task.env"},
		Env: EnvVars{
			"QUX": "4",
			"FUZ": "5",
		},
	})

	assert.NoError(t, err)

	assert.ElementsMatch(t, []string{"FOO=1", "BAR=2", "BAZ=3", "QUX=4", "FUZ=5"}, environ)

}

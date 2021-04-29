package playground

import (
	"github.com/denil/promethee/src/config/generate"
)

func playground() string {
	var PLAYGROUND_TEST string = generate.GenerateNomor()
	return PLAYGROUND_TEST
}

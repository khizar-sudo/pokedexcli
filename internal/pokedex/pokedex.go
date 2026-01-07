package pokedex

import "github.com/khizar-sudo/pokedexcli/internal/pokeapi"

type Pokedex map[string]pokeapi.Pokemon

func NewPokedex() Pokedex {
	return make(Pokedex)
}

module github.com/khizar-sudo/pokedexcli

replace github.com/khizar-sudo/pokedexcli/internal/pokecache v0.0.0 => ./internal/pokecache
replace github.com/khizar-sudo/pokedexcli/internal/pokeapi v0.0.0 => ./internal/pokeapi
replace github.com/khizar-sudo/pokedexcli/internal/commands v0.0.0 => ./internal/commands
replace github.com/khizar-sudo/pokedexcli/internal/utils v0.0.0 => ./internal/utils
replace github.com/khizar-sudo/pokedexcli/internal/constants v0.0.0 => ./internal/constants

go 1.25.5

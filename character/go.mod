module Mucklands/character

go 1.19

replace characterClasses => ./characterClasses

replace equipment => ./equipment

replace homelands => ./homelands

replace species => ./species

replace dice => ../dice

require golang.org/x/exp v0.0.0-20220907003533-145caa8ea1d0 // indirect

require	characterClasses v0.0.0-00010101000000-000000000000
require	dice v0.0.0-00010101000000-000000000000
require	equipment v0.0.0-00010101000000-000000000000

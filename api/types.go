package api

import (
    "encoding/json"
    "net/http"
)

type Type struct {
    name string
}

func (attackType *Type) GetMultiplier(defenderType *Type) float32 {
    return 1.0
}

func GetTypes(w http.ResponseWriter, req *http.Request) {
    types := [18]string{"normal", "fire", "water", "electric", "grass", "flying", "ground", "fighting", "bug", "poison", "psychic", "rock", "ghost", "ice", "dragon", "dark", "steel", "fairy"}
    json.NewEncoder(w).Encode(types)
}

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
    types := [18]string{"Normal", "Fire", "Water", "Electric", "Grass", "Flying", "Ground", "Fighting", "Bug", "Poison", "Psychic", "Rock", "Ghost", "Ice", "Dragon", "Dark", "Steel", "Fairy"}
    json.NewEncoder(w).Encode(types)
}

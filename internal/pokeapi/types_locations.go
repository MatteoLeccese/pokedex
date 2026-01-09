package pokeapi

// Location area structs

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

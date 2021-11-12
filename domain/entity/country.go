package entity


type CountryStruct struct {
	Name      string `json:"name"`
	CountryCode      string `json:"code"`
	Francophone bool `json:"francophone"`
	Anglophone bool `json:"anglophone"`
}


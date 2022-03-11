package entity

type Request struct{}

type Response struct {
	Hackatons *[]HackatonResponse `json:"hackaton"`
}

type HackatonResponse struct {
	ID   uint                `json:"id"`
	Name string              `json:"name"`
	Devs []DeveloperResponse `json:"devs" `
}

type DeveloperResponse struct {
	ID       uint   `json:"id"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

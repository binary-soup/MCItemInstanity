package data

type Info struct {
	Name    string
	Display InfoDisplay `json:"display"`
}

type InfoDisplay struct {
	Item        string `json:"item"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Background  string `json:"background"`
}

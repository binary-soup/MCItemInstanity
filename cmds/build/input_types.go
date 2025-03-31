package build_cmd

type display struct {
	Item        string `json:"item"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Background  string `json:"background"`
}

type author struct {
	Creator string  `json:"creator"`
	Socials socials `json:"socials"`
}

type socials struct {
	Github  string `json:"github"`
	Youtube string `json:"youtube"`
}

type pack struct {
	Description string  `json:"description"`
	Author      author  `json:"author"`
	Root        display `json:"root"`
}

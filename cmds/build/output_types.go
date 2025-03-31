package build_cmd

type mcMeta struct {
	Pack mcMetaPack `json:"pack"`
}

type mcMetaPack struct {
	Description string `json:"description"`
	PackFormat  int    `json:"pack_format"`
}

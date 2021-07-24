package main

type Target struct {
	Bin         string   `json:"bin"`
	OutputDir   string   `json:"outputDir"`
	SourceDir   string   `json:"sourceDir"`
	Flags       []string `json:"flags"`
}

type Build struct {
	Project        string   `json:"project"`
	InstallModules bool     `json:"installModules"`
	Targets        []Target `json:"targets"`
}

package main

type Target struct {
	Bin         string   `json:"bin"`
	OutputDir   string   `json:"outputDir"`
	SourceDir   string   `json:"sourceDir"`
	Flags       []string `json:"flags"`
    Vendor      bool     `json:"vendor"`
    Release     bool     `json:"release"`
}

type Build struct {
	Project        string   `json:"project"`
	InstallModules bool     `json:"installModules"`
	Targets        []Target `json:"targets"`
}

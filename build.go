package main

type Build struct {
	Project        string   `json:"project"`
	InstallModules bool     `json:"installModules"`
	Targets        []Target `json:"targets"`
}
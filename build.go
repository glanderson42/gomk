package main

type Build struct {
	Project        string   `json:"project"`
	InstallModules bool     `json:"installModules"`
	Targets        []Target `json:"targets"`
	Before         Step     `json:"before"`
	After          Step     `json:"after"`
	BeforeEach     Step     `json:"beforeEach"`
	AfterEach      Step     `json:"afterEach"`
}

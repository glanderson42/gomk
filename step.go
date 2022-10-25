package main

import "strings"

type Step struct {
	Command string   `json:"command"`
	Params  []string `json:"params"`
}

func (s *Step) GetCommand() string {
	return s.Command
}

func (s *Step) GetParams() string {
	return strings.Join(s.Params, " ")
}

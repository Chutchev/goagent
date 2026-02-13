package agent

import (
	"log"
)

type Agent struct {
	systemPrompt string
	userPrompt   string
	name         string
	mode         string
}

type AgentInterface interface {
	RunInteractive()
	RunGRPC()
	RunHTTP()
	Do(string)
	GetSystemPrompt() string
	GetUserPrompt() string
	GetMode() string
	GetName() string
}

func NewAgent(systemPrompt, userPrompt, name, mode string) AgentInterface {

	return &Agent{
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
		name:         name,
		mode:         mode,
	}
}

func (a *Agent) RunInteractive()         { log.Fatal("Agent interactive not implemented.") }
func (a *Agent) RunGRPC()                { log.Fatal("Agent grpc not implemented.") }
func (a *Agent) RunHTTP()                { log.Fatal("Agent http not implemented.") }
func (a *Agent) GetSystemPrompt() string { return a.systemPrompt }
func (a *Agent) GetUserPrompt() string   { return a.userPrompt }
func (a *Agent) Do(s string)             { log.Print("Agent Do not implemented") }
func (a *Agent) GetMode() string         { return a.mode }
func (a *Agent) GetName() string         { return a.name }

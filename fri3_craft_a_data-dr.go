package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type AutomationScript struct {
	ID      string `json:"id"`
	Triggers []Trigger `json:"triggers"`
	Actions []Action `json:"actions"`
}

type Trigger struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Action struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (a *AutomationScript) Run() {
	fmt.Printf("Running script %s...\n", a.ID)
	for _, trigger := range a.Triggers {
		fmt.Printf("Evaluating trigger %s...\n", trigger.Type)
		if trigger.Type == "timer" {
			fmt.Println("Timer trigger fired!")
			for _, action := range a.Actions {
				fmt.Printf("Executing action %s...\n", action.Type)
				if action.Type == "email" {
					fmt.Println("Sending email...")
				} else if action.Type == "api" {
					fmt.Println("Calling API...")
				}
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	script := AutomationScript{
		ID: "script-1",
		Triggers: []Trigger{
			{
				Type: "timer",
				Data: "5m",
			},
		},
		Actions: []Action{
			{
				Type: "email",
				Data: "user@example.com",
			},
			{
				Type: "api",
				Data: "https://api.example.com/endpoint",
			},
		},
	}

	scriptJson, _ := json.MarshalIndent(script, "", "  ")
	fmt.Println(string(scriptJson))

	script.Run()
}
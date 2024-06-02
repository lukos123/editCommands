package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	// "github.com/emirpasic/gods/lists/arraylist"
	// "github.com/go-git/go-git/v5/plumbing/object"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// type Commands = []Command

// type Command struct {
// 	Command                     string   `json:"c"`
// 	Help                        string   `json:"h"`
// 	Compile_function            string   `json:"cf"`
// 	Addon_with_compile_function Commands `json:"awcf"`
// 	Next_always                 bool     `json:"na"`
// 	Next                        Commands `json:"n"`
// }

// Greet returns a greeting for the given name

func (a *App) GetJson() []interface{} {
	var commands []interface{}
	file, err := os.Open("command.json")
	if len(os.Args) > 1 {

		file, err = os.Open(os.Args[1] + "/command.json")
	}
	if err != nil {
		fmt.Println(err)

		return commands

	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&commands)
	if err != nil {
		fmt.Println(err)
		return commands
	}
	return commands
}
func (a *App) SaveJson(commands interface{}) bool {
	file, err := os.Create("command.json")
	if len(os.Args) > 1 {

		file, err = os.Create(os.Args[1] + "/command.json")
	}

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(commands)
	if err != nil {
		fmt.Println(err)
		return false

	}
	return true

}

package gocli

import (
	"errors"
	"fmt"
)

var (
	ErrCommandNotFound = errors.New("go-cli: command not found")
)

type CLI struct {
	commands map[string]CLICommandHandler
	running  bool
}

func NewCLI() *CLI {
	return new(CLI)
}

type CLICommandHandler = func(cli *CLI, args []string)

var defaultCLI CLI = CLI{commands: make(map[string]CLICommandHandler), running: false}
var DefaultCLI = &defaultCLI

func (cli *CLI) RegisterCommandHandler(command string, handler CLICommandHandler) {
	if handler == nil {
		panic("go-cli: nil handler")
	}
	cli.commands[command] = handler
}

func RegisterCommandHandler(command string, handler CLICommandHandler) {
	DefaultCLI.RegisterCommandHandler(command, handler)
}

func (cli *CLI) executeCommand(command string, args []string) {
	if command == "exit" {
		cli.running = false
		return
	}
	function, ok := cli.commands[command]
	if !ok {
		fmt.Printf("go-cli: command %s not found\n", command)
		return
	}
	function(cli, args)
}

func (cli *CLI) PrintError(errorMessage string) {
	fmt.Printf("go-cli: %s", errorMessage)
}

func GetArgs() []string {
	args := []string{}
	for {
		var currentArg string
		_, err := fmt.Scan(&currentArg)
		if err != nil || currentArg == "" {
			break
		}
		args = append(args, currentArg)
	}
	return args
}

func (cli *CLI) Run() {
	cli.running = true
	for cli.running {
		var command string
		var arg string
		fmt.Print("Enter command: ")
		fmt.Scan(&command)
		fmt.Scan(&arg)
		cli.executeCommand(command, []string{arg})
	}
}

func Run() {
	DefaultCLI.Run()
}

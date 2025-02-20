package main

import (
	"bufio"
	"cliTaskManager/handler"
	"cliTaskManager/model"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Welcome To task Manager")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		color.Blue("\nTerminating Task Manager, Goodbye!!!")
		os.Exit(0)
	}()

	db, err := model.StartServer()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			color.Blue("\nExiting Task Manager, Goodbye!!")
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "quit" || input == "exit" {
			color.Blue("Exiting Task Manager, Goodbye!!")
			break
		}

		fields := strings.Fields(input)

		if fields[0] != "task" {
			colorRed := color.New(color.FgRed).SprintFunc()
			colorCyan := color.New(color.FgCyan).SprintFunc()
			fmt.Printf("%s \"%s\" %s \n", colorRed("Incorrect command entered, enter"), colorCyan("task <command>"), colorRed("to run a command"))
			continue
		}

		handler.HandleCommands(db, fields)
	}

}

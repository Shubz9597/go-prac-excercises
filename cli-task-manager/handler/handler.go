package handler

import (
	"cliTaskManager/model"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

func printTaskCommand() {
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	cyanCommand := color.New(color.FgCyan, color.Bold).SprintFunc()
	color.Yellow("\ntask is a CLI for managing your TODOs\n\n")
	fmt.Println("Usage:")
	fmt.Fprintf(w, "\t%s\n\n", cyanCommand("task [command]"))
	fmt.Fprintln(w, "Available Commands:")
	fmt.Fprintf(w, "\t%s\t\t%s\n", cyanCommand("add"), "Add a new task to your TODO list")
	fmt.Fprintf(w, "\t%s\t\t%s\n", cyanCommand("do"), "Mark a task on your TODO list as complete")
	fmt.Fprintf(w, "\t%s\t\t%s\n", cyanCommand("list"), "List all of your incomplete tasks")
	fmt.Fprintf(w, "\n\nUse \"%s\" for more information about a command.\n\n", cyanCommand("task [command] --help"))

	w.Flush()
}

func printList(tasks []model.Task) {
	yellowCommand := color.New(color.FgYellow, color.Bold).SprintFunc()
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
	fmt.Print("\n")
	fmt.Printf("%s\n", yellowCommand("You have the following tasks:"))
	for index, val := range tasks {
		fmt.Fprintf(w, "%d.\t%s\n", index+1, val.Task)
	}
	w.Flush()
	fmt.Print("\n")
}

func HandleCommands(ts *model.TaskStore, input []string) {

	if len(input) < 2 {
		printTaskCommand()
		return
	}

	subCommand := input[1]

	switch subCommand {
	case "add":
		if len(input) < 3 {
			color.Red("Error: No Task is provided with the add Command")
			return
		}

		task := strings.Join(input[2:], " ")
		err := ts.AddItem(task)
		if err != nil {
			color.Red("error %s", err)
			return
		}

		color.Green("Task: \"%s\" added successfully\n", task)
		return
	case "do":
		if len(input) < 3 {
			color.Red("Error: No input is provided for marking the task as complete")
			return
		}

		index, err := strconv.Atoi(input[2])
		if err != nil {
			color.Red("Error: the parameter for the todo do should be an integer")
			return
		}

		err = ts.GetIncompleteList(index)
		if err != nil {
			color.Red("Error: failed to update the task - %s", err)
		}

		color.Green("Tasked updated successfully")
		return

	case "list":
		if len(input) > 2 {
			color.Red("Error: there is no input after list")
		}

		tasks, err := ts.ListTasks()

		if err != nil {
			color.Red("Error: there is some error fetching the tasks - %s", err)
			return
		}

		if len(tasks) == 0 {
			color.Yellow("There are no current tasks left to be completed")
			return
		}
		printList(tasks)
		return
	case "rm":
		if len(input) < 3 {
			color.Red("Error: index is required to delete a task")
			return
		}
		index, err := strconv.Atoi(input[2])
		if err != nil {
			color.Red("Error: the parameter for the todo do should be an integer")
			return
		}

		task, err := ts.DeleteTask(index)
		if err != nil {
			color.Red("Error: there is some problem deleting - %s", err)
			return
		}
		color.Green("You have deleted the \"%s\" task.", task.Task)
		return
	default:
		color.Red("Error: Wrong Subcommand entered")
		return
	}

}

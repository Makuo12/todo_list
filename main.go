package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments to perform action!")
		return
	}
	// os.Args[1] should be the action
	switch os.Args[1] {
		case CREATE:
			if task, err := CreateTask(os.Args[2], os.Args[3]); err == nil {
				writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
				format := "%v\t%v\t%v\n"
				fmt.Println("Task successfully created")
				fmt.Fprintf(writer, format, "ID", "Task", "Completed")
				fmt.Fprintf(writer, format, task.ID, task.Text, task.Completed)
				writer.Flush()
			}
		case ADD:
			if task, err := AddTask(os.Args[2], os.Args[3]); err == nil {
				writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
				format := "%v\t%v\t%v\n"
				fmt.Println("Task successfully added")
				fmt.Fprintf(writer, format, "ID", "Task", "Completed")
				fmt.Fprintf(writer, format, task.ID, task.Text, task.Completed)
				writer.Flush()
			}
		
	}
}


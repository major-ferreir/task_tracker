package main

import (
	"fmt"
	"os"
	"github.com/major-ferreir/task_tracker/task"
	"strconv"
)

func main(){
	args := os.Args[1:]

	//	ARGUMENT NUMBER CHECK
	if len(args) > 3 {
		fmt.Println("Too many arguments")
		return
	}

	//	OPEN TASK FILE
	var taskList []task.Task
	task.TaskFile(&taskList)

	//	SELECTING COMMAND		
	switch args[0] {

		case "add":
			if len(args) != 2 {
				fmt.Println("Syntax error!!")
				fmt.Println("Correct usage: ./task-cli add <task-to-be-added>")
				return 
			} 
			task.CreateTask(args[1], &taskList)

		case "update":
			if len(args) != 3 {
				fmt.Println("Syntax error!!")
				fmt.Println("Correct usage: ./task-cli update <task-id> <task-to-be-updated>")
				return 
			}
			taskId, _:= strconv.Atoi(args[1])
			task.UpdateTask(taskId, args[2], &taskList)

		case "delete":
			if len(args) != 2 {
				fmt.Println("Syntax error!!")
				fmt.Println("Correct usage: ./task-cli delete <task-id-to-be-deleted>")
				return 
			} 
			taskId, err:= strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			task.DeleteTask(taskId, &taskList)

		case "mark-done":
			if len(args) != 2 {
				fmt.Println("Syntax error!!")
				fmt.Println("Correct usage: ./task-cli mark-done <task-id-to-be-marked>")
				return 
			} 
		
		case "mark-in-progress":
			if len(args) != 2 {
				fmt.Println("Syntax error!!")
				fmt.Println("Correct usage: ./task-cli mark-in-progress <task-id-to-be-marked>")
				return 
			} 
			fmt.Println(args[0])
		
		case "list":
			lenArgs := len(args)
			if lenArgs == 1{
				task.ListAllTasks(taskList)
			} else if lenArgs == 2{
				task.ListTaskByStatus(args[1], taskList)
			} else {
				fmt.Println("Syntax error!!")
				fmt.Println("Correct usage: ./task-cli list [status]")
				return 
			}

		default:
			fmt.Println("Command " + args[0] + " not found!")
	}

}
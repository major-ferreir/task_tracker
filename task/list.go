package task

import (
	"fmt"
	"strings"
)

func ListAllTasks(taskList []Task) {

	// data, err := json.MarshalIndent(taskList, "", " ")
// 	if err != nil {
// 		panic(err)
// 	}

	if len(taskList) == 0 {
		fmt.Println("No task were added!")
		return
	}

	for i, task := range taskList {

		if i > 0 {
			fmt.Println(strings.Repeat("-", 40))
		}

		fmt.Println("Task Id:", task.Id)
		fmt.Println("Task Name: " + task.TaskName)
		fmt.Println("Task Description:")
		fmt.Println(task.Description)
		fmt.Print("Task Status: ")

		switch task.Status {
		case 0:
			fmt.Println("To-do")
		case 1:
			fmt.Println("In Progress")
		case 2:
			fmt.Println("Done")
		}
	}

	fmt.Println(strings.Repeat("-", 40))
}

func ListTaskByStatus(status string, taskList []Task){

	var localStatus Status
	switch status {
	case "todo":
		localStatus = Todo
	case "in-progress":
		localStatus = InProgress
	case "done":
		localStatus = Done
	default:
		fmt.Println("Incorrect task status")
		fmt.Println("Status avaliable: 'todo' 'in-progress' 'done'")
		return 
	}

	notFound := true

	for i, task := range taskList {
		if task.Status == localStatus {
			if i > 0 {
				fmt.Println(strings.Repeat("-", 40))
				notFound = false
			}
			fmt.Println("Task Id:", task.Id)
			fmt.Println("Task name: " + task.TaskName)
			fmt.Println("Description: " + task.Description)
		}
	}

	if notFound == false {
		fmt.Println(strings.Repeat("-", 40))
	}

}

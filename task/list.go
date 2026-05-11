package task

import (
	"fmt"
	"encoding/json"
)

func ListAllTasks(taskList []Task) {

	data, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		panic(err)
	}

	if len(taskList) == 0 {
		fmt.Println("No task were added!")
		return
	}

	fmt.Println(string(data)) 

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

	for _, task := range taskList {
		if task.Status == localStatus {
			fmt.Println("Task name: " + task.TaskName)
			fmt.Println("Description: " + task.Description)
		}
	}
}

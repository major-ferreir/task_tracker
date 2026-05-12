package task

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)

func UpdateTask(taskId int, newTaskName string, taskList *[]Task) {
	
	for i, task := range *taskList {
		if task.Id == taskId {
			(*taskList)[i].TaskName = newTaskName
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Description: ")
			scanner.Scan()
			description := scanner.Text()

			if description != "" {
				(*taskList)[i].Description = description
			}

			newData, err := json.MarshalIndent(taskList, "", " ")

			if err != nil {
				panic(err)
			}

			err = os.WriteFile(".task-cli.json", newData, 0777)

			if err != nil {
				panic(err)
			}

			fmt.Println("Task updated!")
			break 
		}
	}
}

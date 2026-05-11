package task

import "fmt"
import "os"
import "encoding/json"

func DeleteTask(id int, taskList *[]Task) {

	var indexToDelete int

	for index, task := range *taskList {
		if task.Id == id {
			fmt.Println("id founded")
			indexToDelete = index
			break
		} else {
			fmt.Println("id not fouded")
			return
		}
	}

	*taskList = append((*taskList)[:indexToDelete], (*taskList)[indexToDelete + 1:]...)
	newData, err := json.MarshalIndent(*taskList, "", " ")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(".task-cli.json", newData, 0777)

		if err != nil {
		panic(err)
	}

	fmt.Println("Task deleted")
}

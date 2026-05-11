package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func CreateTask(taskName string, taskList *[]Task) {

	fmt.Println("Create Task "+ taskName)
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Task Description:")
	scanner.Scan()
	descritpion := scanner.Text()

	id := 1
	taskLen := len(*taskList)
	
	if taskLen > 0 {
		taskLen++
		id = taskLen
	}

	newTask := Task{
		TaskName: taskName,
		Description: descritpion,
		Id: id,
		CreateAt: time.Now(),
	}
	*taskList = append(*taskList, newTask)
	
	newData, err := json.MarshalIndent(taskList, "", " ")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(".task-cli.json", newData, 0777)


	if err != nil {
		panic(err)
	}
	fmt.Println("Task added!")
}

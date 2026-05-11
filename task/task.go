package task

import (
	"time"
	"os"
	"encoding/json"
)

type Status int

const (
	Todo  Status = iota
	InProgress
	Done
)

type Task struct{
	TaskName string		`json:"task_name"`
	Id int				`json:"id"`
	Description string	`json:"description"`
	Status Status		`json:"status"`
	CreateAt time.Time	`json:"created_at"`
	UpdatedAt *time.Time	`json:"deleted_at,omitempty"`
}

func TaskFile(taskList *[]Task){

	file, err := os.ReadFile(".task-cli.json")

	if err != nil {

		//	CREATE FILE IF IT DOES NOT EXIST
		newFile, erro := os.OpenFile(".task-cli.json", os.O_CREATE, 0777)
		if erro != nil {
			panic(err)
		}

		newFile.Close()
	}

	if len(file) != 0 {
		err = json.Unmarshal(file, &taskList)
		
		if err != nil {
			panic(err)
		}
	}
}

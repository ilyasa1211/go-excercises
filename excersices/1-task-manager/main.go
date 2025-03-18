package main

import (
	"fmt"
	"reflect"
	"slices"
	"time"

	"github.com/ilyasa1211/go-exercises/excersices/1-task-manager/tasks/entities"
	"github.com/ilyasa1211/go-exercises/excersices/1-task-manager/tasks/interfaces"
)

type ITaskManager interface {
	// execute all tasks
	ExecuteAll()

	// get all logs from the executed tasks
	PrintLogs()

	// retry failed task
	RetryFailed()
}

type TaskManager struct {
	tasks []interfaces.Task
	logs  []entities.TaskLog
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) AddTask(t ...interfaces.Task) error {
	tm.tasks = append(tm.tasks, t...)

	return nil
}

func (tm *TaskManager) PrintLogs() {
	for _, log := range tm.logs {
		fmt.Println(log)
	}

	// clear logs
	tm.logs = nil
}

func (tm *TaskManager) ExecuteAll() {
	for i := 0; i < len(tm.tasks); i++ {
		task := tm.tasks[i]

		var log entities.TaskLog
		status := "OK"
		var desc string

		startTime := time.Now()

		// execute the task
		if err := task.Execute(); err != nil {
			status = "Error"
			desc = err.Error()

		} else {
			desc = reflect.ValueOf(task).Elem().FieldByName("Description").String()
			// remove element
			tm.tasks = slices.Delete(tm.tasks, i, i+1)
			i--
		}

		endTime := time.Now()
		deltaTime := endTime.Sub(startTime)

		// collecting log
		log = entities.TaskLog{
			Log: interfaces.Log{
				Type:        reflect.TypeOf(task).String(),
				Status:      status,
				Timestamp:   endTime,
				Description: desc,
			},
			StartTime: startTime,
			EndTime:   endTime,
			DeltaTime: deltaTime,
		}

		tm.logs = append(tm.logs, log)
	}
}

func (tm *TaskManager) RetryFailed() {
	fmt.Println("Retrying failed task")
	if len(tm.tasks) > 0 {
		tm.ExecuteAll()
	} else {
		fmt.Println("No task failed")
	}
}

func main() {
	taskManager := NewTaskManager()

	emailTask := entities.NewEmailTask("1", time.Now(), "Hello From Main!")
	reportTask := entities.NewReportTask("2", time.Now(), "Creating System Report")
	smsTask := entities.NewSmsTask("3", time.Now(), "Here is your verification code: 19283")

	taskManager.AddTask(emailTask, reportTask, smsTask)

	taskManager.ExecuteAll()
	taskManager.PrintLogs()

	taskManager.RetryFailed()
	taskManager.PrintLogs()
}

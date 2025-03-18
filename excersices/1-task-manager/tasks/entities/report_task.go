package entities

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ilyasa1211/go-exercises/excersices/1-task-manager/tasks/utils"
)

type ReportTask struct {
	BaseTask
}

func NewReportTask(id string, createdAt time.Time, description string) *ReportTask {
	return &ReportTask{
		BaseTask: BaseTask{
			ID:          id,
			CreatedAt:   createdAt,
			Description: description,
		},
	}
}

func (t *ReportTask) Execute() error {
	fmt.Println("Simulating report: ", t.Description)

	// heavy process
	time.Sleep(2 * time.Second)

	if err := utils.GetRandomFail("Failed to do report", 20); err != nil {
		return err
	}

	return nil
}

func (t *ReportTask) Info() string {
	name := reflect.TypeOf(t).String()
	desc := reflect.ValueOf(t).FieldByName("Description").String()

	return fmt.Sprintf("name=%s; desc=%s;\n", name, desc)
}

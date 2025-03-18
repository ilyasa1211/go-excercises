package entities

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ilyasa1211/go-exercises/excersices/1-task-manager/tasks/utils"
)

type EmailTask struct {
	BaseTask
}

func NewEmailTask(id string, createdAt time.Time, description string) *EmailTask {
	return &EmailTask{
		BaseTask: BaseTask{
			ID:          id,
			CreatedAt:   createdAt,
			Description: description,
		},
	}
}

func (t *EmailTask) Execute() error {
	fmt.Println("Simulating sending email: ", t.Description)

	// heavy process
	time.Sleep(3 * time.Second)

	if err := utils.GetRandomFail("Failed to send email", 20); err != nil {
		return err
	}

	return nil
}

func (t *EmailTask) Info() string {
	name := reflect.TypeOf(t).String()
	desc := reflect.ValueOf(t).FieldByName("Description").String()

	return fmt.Sprintf("name=%s; desc=%s;\n", name, desc)
}

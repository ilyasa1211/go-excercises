package entities

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ilyasa1211/go-exercises/excersices/1-task-manager/tasks/utils"
)

type SmsTask struct {
	BaseTask
}

func NewSmsTask(id string, createdAt time.Time, description string) *SmsTask {
	return &SmsTask{
		BaseTask: BaseTask{
			ID:          id,
			CreatedAt:   createdAt,
			Description: description,
		},
	}
}

func (t *SmsTask) Execute() error {
	fmt.Println("Simulating sending via sms: ", t.Description)

	// heavy process
	time.Sleep(4 * time.Second)

	if err := utils.GetRandomFail("Failed to send via sms", 20); err != nil {
		return err
	}

	return nil
}

func (t *SmsTask) Info() string {
	name := reflect.TypeOf(t).String()
	desc := reflect.ValueOf(t).FieldByName("Description").String()

	return fmt.Sprintf("name=%s; desc=%s;\n", name, desc)
}

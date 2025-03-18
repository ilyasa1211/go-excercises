package interfaces

type Task interface {

	// Execute the task
	Execute() error

	// Retrieve informations about the task
	Info() string
}

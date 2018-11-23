package abstractions

import (
	"fmt"
	"strings"
)

// TaskType alias
type TaskType int

const (
	// Send !
	Send TaskType = 0
	// Receive !
	Receive TaskType = 1
	// Noop !
	Noop TaskType = 2
	// Unknown !
	Unknown TaskType = -1
)

func (tt TaskType) String() string {
	if tt == Send {
		return "Send"

	} else if tt == Receive {
		return "Receive"

	} else if tt == Noop {
		return "Noop"

	} else {
		return "Unknown"
	}
}

// Task encapsulates info about something a process must do
type Task struct {
	// TaskType !
	TaskType TaskType
	// Dependent process
	DependentProcessID int
}

// ParseTaskType creates a taskType variable from the given string
// TODO support case insensitivity
func ParseTaskType(str string) TaskType {
	t := Unknown
	if strings.Compare(str, "Send") == 0 {
		t = Send

	} else if strings.Compare(str, "Receive") == 0 {
		t = Receive

	} else if strings.Compare(str, "Noop") == 0 {
		t = Noop

	}
	return t
}

// NewTask creates a task
func NewTask(taskType TaskType, dependentProcessID int) Task {
	return Task{taskType, dependentProcessID}
}

func (t Task) String() string {
	return fmt.Sprintf("Task of type '%v', depending on Process #%v", t.TaskType, t.DependentProcessID)
}

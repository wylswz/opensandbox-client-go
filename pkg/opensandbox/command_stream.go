package opensandbox

import "encoding/json"

// CommandStreamEventType is the event type emitted by command stream APIs.
type CommandStreamEventType string

const (
	CommandStreamEventInit              CommandStreamEventType = "init"
	CommandStreamEventStatus            CommandStreamEventType = "status"
	CommandStreamEventError             CommandStreamEventType = "error"
	CommandStreamEventStdout            CommandStreamEventType = "stdout"
	CommandStreamEventStderr            CommandStreamEventType = "stderr"
	CommandStreamEventResult            CommandStreamEventType = "result"
	CommandStreamEventExecutionComplete CommandStreamEventType = "execution_complete"
	CommandStreamEventExecutionCount    CommandStreamEventType = "execution_count"
	CommandStreamEventPing              CommandStreamEventType = "ping"
)

// CommandStreamError is structured error data attached to a stream event.
type CommandStreamError struct {
	Name      string
	Value     string
	Traceback []string
}

// CommandStreamEvent is a typed event emitted while a command is running.
// Raw contains the original JSON payload for forward compatibility.
type CommandStreamEvent struct {
	Type           CommandStreamEventType
	Text           string
	ExecutionCount *int32
	ExecutionTime  *int64
	Timestamp      *int64
	Results        map[string]interface{}
	Error          *CommandStreamError
	Raw            json.RawMessage
}

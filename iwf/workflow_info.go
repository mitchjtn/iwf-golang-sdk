package iwf

import "github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"

type WorkflowInfo struct {
	Status       iwfidl.WorkflowStatus
	CurrentRunId string
}

package iwf

import "github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"

type WorkflowStopOptions struct {
	StopType iwfidl.WorkflowStopType
	Reason   string
}

package iwf

import "github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"

type UnregisteredWorkflowOptions struct {
	WorkflowIdReusePolicy              *iwfidl.IDReusePolicy
	WorkflowCronSchedule               *string
	WorkflowRetryPolicy                *iwfidl.WorkflowRetryPolicy
	StartStateOptions                  *iwfidl.WorkflowStateOptions
	InitialSearchAttributes            []iwfidl.SearchAttribute
	WaitForCompletionStateExecutionIds []string
	WaitForCompletionStateIds          []string
}

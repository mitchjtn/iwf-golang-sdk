package iwf

import "github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"

type UnregisteredWorkflowOptions struct {
	WorkflowIdReusePolicy     *iwfidl.IDReusePolicy
	WorkflowCronSchedule      *string
	WorkflowStartDelaySeconds *int32
	WorkflowRetryPolicy       *iwfidl.WorkflowRetryPolicy
	StartStateOptions         *iwfidl.WorkflowStateOptions
	InitialSearchAttributes   []iwfidl.SearchAttribute
	WaitForCompletionState    []WorkflowState
}

func getWaitForCompletionStateExecutionIds(states []WorkflowState) []string {
	results := []string{}

	for _, state := range states {
		results = append(results, GetStateExecutionId(state, 1))
	}

	return results
}

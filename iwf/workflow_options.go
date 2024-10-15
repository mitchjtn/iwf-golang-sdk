package iwf

import "github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"

type WorkflowOptions struct {
	WorkflowIdReusePolicy *iwfidl.IDReusePolicy
	WorkflowCronSchedule  *string
	WorkflowRetryPolicy   *iwfidl.WorkflowRetryPolicy
	// InitialSearchAttributes set the initial search attributes to start a workflow
	// key is search attribute key, value much match with PersistenceSchema of the workflow definition
	// For iwfidl.DATETIME , the value can be either time.Time or a string value in format of DateTimeFormat
	InitialSearchAttributes map[string]interface{}
	WaitForCompletionState  []WorkflowState
}

func getWaitForCompletionStateExecutionIds(states []WorkflowState) []string {
	results := []string{}

	for _, state := range states {
		results = append(results, GetStateExecutionId(state, 1))
	}

	return results
}

func getWaitForCompletionStateIds(states []WorkflowState) []string {
	results := []string{}

	for _, state := range states {
		results = append(results, state.GetStateId())
	}

	return results
}

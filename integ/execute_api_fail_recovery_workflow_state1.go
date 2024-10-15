package integ

import (
	"errors"

	"github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
)

type executeApiFailRecoveryWorkflowState1 struct {
	iwf.WorkflowStateDefaultsNoWaitUntil
}

func (b executeApiFailRecoveryWorkflowState1) GetStateId() string {
	return "execute_api_fail_recovery_workflow_state1"
}

func (b executeApiFailRecoveryWorkflowState1) GetStateOptions() *iwf.StateOptions {
	options := &iwf.StateOptions{
		ExecuteApiRetryPolicy: &iwfidl.RetryPolicy{
			InitialIntervalSeconds: iwfidl.PtrInt32(1),
			MaximumAttempts:        iwfidl.PtrInt32(1),
		},
		ExecuteApiFailureProceedState: &executeApiFailRecoveryWorkflowState2{},
	}

	return options
}

func (b executeApiFailRecoveryWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return nil, errors.New("error")
}

package integ

import (
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
)

type interStateWorkflowState2 struct {
	iwf.WorkflowStateDefaults
}

func (b interStateWorkflowState2) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	communication.PublishInternalChannel(interStateChannel2, 2)
	return iwf.EmptyCommandRequest(), nil
}

func (b interStateWorkflowState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.GracefulCompletingWorkflow, nil
}

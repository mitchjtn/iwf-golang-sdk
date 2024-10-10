package integ

import (
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
	"time"
)

type timerWorkflowState1 struct {
	iwf.DefaultStateId
	iwf.DefaultStateOptions
}

func (b timerWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var i int
	input.Get(&i)
	return iwf.AllCommandsCompletedRequest(
		iwf.NewTimerCommandByDuration("", time.Duration(i)*time.Second),
	), nil
}

func (b timerWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	input.Get(&i)
	return iwf.GracefulCompleteWorkflow(i + 1), nil
}

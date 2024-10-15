package integ

import (
	"fmt"
	"github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
)

type interStateWorkflowState1 struct {
	iwf.WorkflowStateDefaults
}

func (b interStateWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.AnyCommandCompletedRequest(
			iwf.NewInternalChannelCommand("id1", interStateChannel1),
			iwf.NewInternalChannelCommand("id2", interStateChannel2)),
		nil
}

func (b interStateWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	cmd1 := commandResults.GetInternalChannelCommandResultById("id1")
	cmd2 := commandResults.GetInternalChannelCommandResultById("id2")
	cmd2.Value.Get(&i)
	if cmd1.Status == iwfidl.WAITING && i == 2 {
		return iwf.GracefulCompletingWorkflow, nil
	}
	return nil, fmt.Errorf("error in executing " + ctx.GetStateExecutionId())
}

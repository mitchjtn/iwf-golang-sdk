package integ

import (
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
)

type skipWaitUntilState2 struct {
	iwf.WorkflowStateDefaultsNoWaitUntil
}

func (b skipWaitUntilState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	input.Get(&i)
	return iwf.GracefulCompleteWorkflow(i + 1), nil
}

package integ

import "github.com/mitchjtn/iwf-golang-sdk/iwf"

type stateApiTimeoutWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyCommunicationSchema
	iwf.EmptyPersistenceSchema
}

func (b stateApiTimeoutWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&stateApiTimeoutWorkflowState1{}),
	}
}

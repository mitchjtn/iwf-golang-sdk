package integ

import "github.com/mitchjtn/iwf-golang-sdk/iwf"

type stateApiFailWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyCommunicationSchema
	iwf.EmptyPersistenceSchema
}

func (b stateApiFailWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&stateApiFailWorkflowState1{}),
	}
}

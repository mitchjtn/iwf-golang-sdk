package integ

import "github.com/mitchjtn/iwf-golang-sdk/iwf"

type forceFailWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyCommunicationSchema
	iwf.EmptyPersistenceSchema
}

func (b forceFailWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&forceFailWorkflowState1{}),
	}
}

package integ

import (
	"github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
)

type persistenceWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyCommunicationSchema
}

const (
	testDataObjectKey  = "test-data-object"
	testDataObjectKey2 = "test-data-object-2"

	testSearchAttributeInt      = "CustomIntField"
	testSearchAttributeDatetime = "CustomDatetimeField"
	testSearchAttributeBool     = "CustomBoolField"
	testSearchAttributeDouble   = "CustomDoubleField"
	testSearchAttributeText     = "CustomStringField"
	testSearchAttributeKeyword  = "CustomKeywordField"
)

func (b persistenceWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&persistenceWorkflowState1{}),
		iwf.NonStartingStateDef(&persistenceWorkflowState2{}),
	}
}

func (b persistenceWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return []iwf.PersistenceFieldDef{
		iwf.DataAttributeDef(testDataObjectKey),
		iwf.DataAttributeDef(testDataObjectKey2),
		iwf.SearchAttributeDef(testSearchAttributeInt, iwfidl.INT),
		iwf.SearchAttributeDef(testSearchAttributeDatetime, iwfidl.DATETIME),
		iwf.SearchAttributeDef(testSearchAttributeBool, iwfidl.BOOL),
		iwf.SearchAttributeDef(testSearchAttributeDouble, iwfidl.DOUBLE),
		iwf.SearchAttributeDef(testSearchAttributeText, iwfidl.TEXT),
		iwf.SearchAttributeDef(testSearchAttributeKeyword, iwfidl.KEYWORD),
	}
}

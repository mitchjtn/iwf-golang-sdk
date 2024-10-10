package integ

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/mitchjtn/iwf-golang-sdk/iwf"
	"github.com/stretchr/testify/assert"
)

func TestSignalWorkflow(t *testing.T) {
	wfId := "TestSignalWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &signalWorkflow{}, wfId, 10, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)
	err = client.SignalWorkflow(context.Background(), &signalWorkflow{}, wfId, "", testChannelName2, 10)
	assert.Nil(t, err)

	// wait for timer to be ready to be skipped
	time.Sleep(time.Second)
	err = client.SignalWorkflow(context.Background(), &signalWorkflow{}, wfId, "", testChannelName1, 100)
	assert.Nil(t, err)

	err = client.SkipTimerByCommandIndex(context.Background(), wfId, "", signalWorkflowState2{}, 1, 0)
	assert.Nil(t, err)

	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 100, output)

	err = client.SignalWorkflow(context.Background(), &signalWorkflow{}, "a wrong workflowId", "", testChannelName1, 100)
	assert.True(t, iwf.IsWorkflowNotExistsError(err))
}

func TestSignalWorkflowWithUntypedClient(t *testing.T) {
	unregisteredClient := iwf.NewUnregisteredClient(nil)

	wfType := iwf.GetFinalWorkflowType(&signalWorkflow{})
	wfId := "TestSignalWorkflowWithUntypedClient" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := unregisteredClient.StartWorkflow(context.Background(), wfType, iwf.GetFinalWorkflowStateId(signalWorkflowState1{}), wfId, 10, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)
	err = unregisteredClient.SignalWorkflow(context.Background(), wfId, "", testChannelName2, 10)
	assert.Nil(t, err)

	// wait for timer to be ready to be skipped
	time.Sleep(time.Second)
	err = unregisteredClient.SignalWorkflow(context.Background(), wfId, "", testChannelName1, 100)
	assert.Nil(t, err)

	err = unregisteredClient.SkipTimerByCommandIndex(context.Background(), wfId, "", iwf.GetFinalWorkflowStateId(signalWorkflowState2{}), 1, 0)
	assert.Nil(t, err)

	var output int
	err = unregisteredClient.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 100, output)
}

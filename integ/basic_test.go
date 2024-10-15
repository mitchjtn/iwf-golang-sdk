package integ

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
	"github.com/mitchjtn/iwf-golang-sdk/iwf/ptr"
	"github.com/stretchr/testify/assert"
)

func TestBasicWorkflow(t *testing.T) {
	wfId := "TestBasicWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &basicWorkflow{}, wfId, 10, 1, &iwf.WorkflowOptions{
		WorkflowIdReusePolicy: ptr.Any(iwfidl.DISALLOW_REUSE),
		WorkflowRetryPolicy: &iwfidl.WorkflowRetryPolicy{
			InitialIntervalSeconds: iwfidl.PtrInt32(10),
			MaximumAttempts:        iwfidl.PtrInt32(3),
			MaximumIntervalSeconds: iwfidl.PtrInt32(100),
			BackoffCoefficient:     iwfidl.PtrFloat32(3),
		},
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	// start the same workflowId again will fail
	_, err = client.StartWorkflow(context.Background(), &basicWorkflow{}, wfId, 10, nil, nil)
	assert.True(t, iwf.IsWorkflowAlreadyStartedError(err))

	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)

	err = client.GetSimpleWorkflowResult(context.Background(), "a wrong workflowId", "", &output)
	assert.True(t, iwf.IsWorkflowNotExistsError(err))
}

func TestProceedOnStateStartFailWorkflow(t *testing.T) {
	wfId := "TestProceedOnStateStartFailWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &proceedOnStateStartFailWorkflow{}, wfId, 10, "input", &iwf.WorkflowOptions{})
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	_, err = client.StartWorkflow(context.Background(), &basicWorkflow{}, wfId, 10, nil, nil)
	assert.True(t, iwf.IsWorkflowAlreadyStartedError(err))

	var output string
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Equal(t, "input_state1_start_state1_decide_state2_start_state2_decide", output)
	assert.Nil(t, err)
}

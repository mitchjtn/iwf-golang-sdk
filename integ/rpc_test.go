package integ

import (
	"context"
	"github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"
	"github.com/mitchjtn/iwf-golang-sdk/iwf"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRPCWorkflow(t *testing.T) {
	wfId := "TestRPCWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	wf := rpcWorkflow{}

	runId, err := client.StartWorkflow(context.Background(), wf, wfId, 10, 1, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	time.Sleep(time.Second)
	info, err := client.DescribeWorkflow(context.Background(), wfId, "")
	assert.Nil(t, err)
	assert.Equal(t, iwfidl.RUNNING, info.Status)

	err = client.InvokeRPC(context.Background(), wfId, "", wf.TestErrorRPC, 1, nil)
	assert.NotNil(t, err)
	assert.True(t, iwf.IsRPCError(err))
	rpcErr, _ := err.(*iwf.ApiError)
	assert.Equal(t, "worker API error, status:501, errorType:test-error-type", rpcErr.Response.GetDetail())

	var rpcOutput int
	err = client.InvokeRPC(context.Background(), wfId, "", wf.TestRPC, 1, &rpcOutput)
	assert.Nil(t, err)
	assert.Equal(t, 2, rpcOutput)

	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)
}

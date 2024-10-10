package iwf

import (
	"fmt"
	"github.com/mitchjtn/iwf-golang-sdk/gen/iwfidl"
	"log"
	"net/http"
	"runtime/debug"
)

type InvalidArgumentError struct {
	msg string
}

func (w InvalidArgumentError) Error() string {
	return fmt.Sprintf("WorkflowDefinitionError: %s", w.msg)
}

func NewInvalidArgumentError(msg string) error {
	return &InvalidArgumentError{
		msg: msg,
	}
}

func NewInvalidArgumentErrorFmt(tpl string, arg ...interface{}) error {
	return &WorkflowDefinitionError{
		msg: fmt.Sprintf(tpl, arg...),
	}
}

type WorkflowDefinitionError struct {
	msg string
}

func (w WorkflowDefinitionError) Error() string {
	return fmt.Sprintf("WorkflowDefinitionError: %s", w.msg)
}

func NewWorkflowDefinitionError(msg string) error {
	return &WorkflowDefinitionError{
		msg: msg,
	}
}

func NewWorkflowDefinitionErrorFmt(tpl string, arg ...interface{}) error {
	return &WorkflowDefinitionError{
		msg: fmt.Sprintf(tpl, arg...),
	}
}

type InternalError struct {
	Message string
}

func newInternalError(format string, args ...interface{}) error {
	return &InternalError{
		Message: fmt.Sprintf(format, args...),
	}
}

func (i InternalError) Error() string {
	return fmt.Sprintf("error in SDK or service: message:%v", i.Message)
}

type StateExecutionError struct {
	OriginalError error
	StackTrace    string
}

func newStateExecutionError(err error, stackTrace string) error {
	return &StateExecutionError{
		OriginalError: err,
		StackTrace:    stackTrace,
	}
}

func (i StateExecutionError) Error() string {
	return fmt.Sprintf("error message:%v, stacktrace: %v", i.OriginalError, i.StackTrace)
}

// for skipping the logging in testing code
var skipCaptureErrorLogging = false

// MUST be the result from calling recover, which MUST be done in a single level deep
// deferred function. The usual way of calling this is:
// - defer func() { captureStateExecutionError(recover(), logger, &err) }()
func captureStateExecutionError(errPanic interface{}, retError *error) {
	if errPanic != nil || *retError != nil {
		st := string(debug.Stack())

		var err error
		panicError, ok := errPanic.(error)
		if errPanic != nil {
			if ok && panicError != nil {
				err = newStateExecutionError(panicError, st)
			} else {
				err = newStateExecutionError(fmt.Errorf("errPanic is not an error %v", errPanic), st)
			}
		} else {
			err = newStateExecutionError(*retError, st)
		}

		if !skipCaptureErrorLogging && errPanic != nil {
			log.Printf("panic is captured: %v , stacktrace: %v", errPanic, st)
		}
		*retError = err
	}
}

type ApiError struct {
	StatusCode    int
	OriginalError error
	OpenApiError  *iwfidl.GenericOpenAPIError
	HttpResponse  *http.Response
	Response      *iwfidl.ErrorResponse
}

func (i *ApiError) Error() string {
	if i.Response != nil {
		return i.OriginalError.Error() + "\n" + i.Response.GetOriginalWorkerErrorDetail()
	}
	return i.OriginalError.Error()
}

func NewApiError(originalError error, openApiError *iwfidl.GenericOpenAPIError, httpResponse *http.Response, response *iwfidl.ErrorResponse) error {
	statusCode := 0
	if httpResponse != nil {
		statusCode = httpResponse.StatusCode
	}
	return &ApiError{
		StatusCode:    statusCode,
		OriginalError: originalError,
		OpenApiError:  openApiError,
		HttpResponse:  httpResponse,
		Response:      response,
	}
}

// GetOpenApiErrorBody retrieve the API error body into a string to be human-readable
func GetOpenApiErrorBody(err error) string {
	apiError, ok := err.(*ApiError)
	if !ok {
		return "not an ApiError"
	}
	return string(apiError.OpenApiError.Body())
}

func IsClientError(err error) bool {
	apiError, ok := err.(*ApiError)
	if !ok {
		return false
	}
	return apiError.StatusCode >= 400 && apiError.StatusCode < 500
}

func IsWorkflowAlreadyStartedError(err error) bool {
	apiError, ok := err.(*ApiError)
	if !ok || apiError.Response == nil {
		return false
	}
	return apiError.Response.GetSubStatus() == iwfidl.WORKFLOW_ALREADY_STARTED_SUB_STATUS
}

func IsWorkflowNotExistsError(err error) bool {
	apiError, ok := err.(*ApiError)
	if !ok || apiError.Response == nil {
		return false
	}
	return apiError.Response.GetSubStatus() == iwfidl.WORKFLOW_NOT_EXISTS_SUB_STATUS
}

func IsRPCError(err error) bool {
	apiError, ok := err.(*ApiError)
	if !ok || apiError.Response == nil {
		return false
	}
	return apiError.StatusCode == 420
}

type WorkflowUncompletedError struct {
	RunId        string
	ClosedStatus iwfidl.WorkflowStatus
	ErrorType    *iwfidl.WorkflowErrorType
	ErrorMessage *string
	StateResults []iwfidl.StateCompletionOutput
	Encoder      ObjectEncoder
}

func NewWorkflowUncompletedError(
	runId string, closedStatus iwfidl.WorkflowStatus, errorType *iwfidl.WorkflowErrorType,
	errorMessage *string, stateResults []iwfidl.StateCompletionOutput, encoder ObjectEncoder) error {
	return &WorkflowUncompletedError{
		RunId:        runId,
		ClosedStatus: closedStatus,
		ErrorType:    errorType,
		ErrorMessage: errorMessage,
		StateResults: stateResults,
		Encoder:      encoder,
	}
}
func (w *WorkflowUncompletedError) Error() string {
	errTypeMsg := "<nil>"
	message := "<nil>"
	if w.ErrorType != nil {
		errTypeMsg = fmt.Sprintf("%v", *w.ErrorType)
	}
	if w.ErrorMessage != nil {
		message = fmt.Sprintf("%v", *w.ErrorMessage)
	}
	return fmt.Sprintf("workflow is not completed successfully, closedStatus: %v, failedErrorType(applies if failed as closedStatus):%v, error message:%v",
		w.ClosedStatus, errTypeMsg, message)
}

// AsWorkflowUncompletedError will check if it's a WorkflowUncompletedError and convert it if so
func AsWorkflowUncompletedError(err error) (*WorkflowUncompletedError, bool) {
	wErr, ok := err.(*WorkflowUncompletedError)
	return wErr, ok
}

func (w *WorkflowUncompletedError) GetStateResult(index int, resultPtr interface{}) error {
	output := w.StateResults[index]
	return w.Encoder.Decode(output.CompletedStateOutput, resultPtr)
}

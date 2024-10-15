# iWF Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/mitchjtn/iwf-golang-sdk.svg)](https://pkg.go.dev/github.com/mitchjtn/iwf-golang-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/mitchjtn/iwf-golang-sdk)](https://goreportcard.com/report/github.com/mitchjtn/iwf-golang-sdk)
[![Coverage Status](https://codecov.io/github/indeedeng/iwf-golang-sdk/coverage.svg?branch=main)](https://app.codecov.io/gh/indeedeng/iwf-golang-sdk/branch/main)

[![Build status](https://github.com/mitchjtn/iwf-golang-sdk/actions/workflows/ci-integ-test.yml/badge.svg?branch=main)](https://github.com/mitchjtn/iwf-golang-sdk/actions/workflows/ci-integ-test.yml)

Golang SDK for [iWF workflow engine](https://github.com/mitchjtn/iwf)

See [samples](https://github.com/mitchjtn/iwf-golang-samples) for how to use this SDK.

## Contribution

See [contribution guide](CONTRIBUTION.md)

## Development Plan

### 1.0

- [x] Start workflow API
- [x] Executing `start`/`decide` APIs and completing workflow
- [x] Parallel execution of multiple states
- [x] Timer command
- [x] Signal command
- [x] SearchAttribute
- [x] DataAttributes
- [x] StateExecutionLocal
- [x] Signal workflow API
- [x] Get workflow result API
- [x] Search workflow API
- [x] Describe workflow API
- [x] Stop workflow API
- [x] Reset workflow API
- [x] Command type(s) for inter-state communications (e.g. internal channel)
- [x] More workflow start options: IdReusePolicy, cron schedule, retry
- [x] StateOption: Start/Decide API timeout and retry policy
- [x] Reset workflow by stateId/StateExecutionId
- [x] More workflow start options: initial search attributes

### 1.1

- [x] Skip timer API for testing/operation
- [x] Decider trigger type: any command combination

### 1.2

- [x] API improvements to reduce boilerplate code

### 1.3

- [x] Support failing workflow with results
- [x] Improve workflow uncompleted error return(canceled, failed, timeout, terminated)

### 1.4

- [x] Renaming some concepts/APIs with breaking changes(see release notes)
- [x] Support workflow RPC
- [x] PARTIAL_WITH_EXCLUSIVE_LOCK persistence loading type

/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package freeze

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2FreezeWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2FreezeWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2FreezeWebSocketClient) describeStagesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStagesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStagesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeWebSocketClient) DescribeStagesAsync(
	request *DescribeStagesRequest,
	callback chan<- DescribeStagesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "freeze",
			"component":   "stage",
			"function":    "describeStages",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeStagesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FreezeWebSocketClient) DescribeStages(
	request *DescribeStagesRequest,
) (*DescribeStagesResult, error) {
	callback := make(chan DescribeStagesAsyncResult, 1)
	go p.DescribeStagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FreezeWebSocketClient) getStageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStageResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStageAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeWebSocketClient) GetStageAsync(
	request *GetStageRequest,
	callback chan<- GetStageAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "freeze",
			"component":   "stage",
			"function":    "getStage",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StageName != nil && *request.StageName != "" {
		bodies["stageName"] = *request.StageName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getStageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FreezeWebSocketClient) GetStage(
	request *GetStageRequest,
) (*GetStageResult, error) {
	callback := make(chan GetStageAsyncResult, 1)
	go p.GetStageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FreezeWebSocketClient) promoteStageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PromoteStageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PromoteStageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PromoteStageResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PromoteStageAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PromoteStageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeWebSocketClient) PromoteStageAsync(
	request *PromoteStageRequest,
	callback chan<- PromoteStageAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "freeze",
			"component":   "stage",
			"function":    "promoteStage",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StageName != nil && *request.StageName != "" {
		bodies["stageName"] = *request.StageName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.promoteStageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FreezeWebSocketClient) PromoteStage(
	request *PromoteStageRequest,
) (*PromoteStageResult, error) {
	callback := make(chan PromoteStageAsyncResult, 1)
	go p.PromoteStageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FreezeWebSocketClient) rollbackStageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RollbackStageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RollbackStageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RollbackStageResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RollbackStageAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RollbackStageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeWebSocketClient) RollbackStageAsync(
	request *RollbackStageRequest,
	callback chan<- RollbackStageAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "freeze",
			"component":   "stage",
			"function":    "rollbackStage",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StageName != nil && *request.StageName != "" {
		bodies["stageName"] = *request.StageName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.rollbackStageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FreezeWebSocketClient) RollbackStage(
	request *RollbackStageRequest,
) (*RollbackStageResult, error) {
	callback := make(chan RollbackStageAsyncResult, 1)
	go p.RollbackStageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FreezeWebSocketClient) describeOutputsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeOutputsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeOutputsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeOutputsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeOutputsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeOutputsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeWebSocketClient) DescribeOutputsAsync(
	request *DescribeOutputsRequest,
	callback chan<- DescribeOutputsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "freeze",
			"component":   "output",
			"function":    "describeOutputs",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StageName != nil && *request.StageName != "" {
		bodies["stageName"] = *request.StageName
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeOutputsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FreezeWebSocketClient) DescribeOutputs(
	request *DescribeOutputsRequest,
) (*DescribeOutputsResult, error) {
	callback := make(chan DescribeOutputsAsyncResult, 1)
	go p.DescribeOutputsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FreezeWebSocketClient) getOutputAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetOutputAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetOutputAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetOutputResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetOutputAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetOutputAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeWebSocketClient) GetOutputAsync(
	request *GetOutputRequest,
	callback chan<- GetOutputAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "freeze",
			"component":   "output",
			"function":    "getOutput",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StageName != nil && *request.StageName != "" {
		bodies["stageName"] = *request.StageName
	}
	if request.OutputName != nil && *request.OutputName != "" {
		bodies["outputName"] = *request.OutputName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getOutputAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FreezeWebSocketClient) GetOutput(
	request *GetOutputRequest,
) (*GetOutputResult, error) {
	callback := make(chan GetOutputAsyncResult, 1)
	go p.GetOutputAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

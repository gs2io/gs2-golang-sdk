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

package deploy

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2DeployWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2DeployWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2DeployWebSocketClient) describeStacksAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStacksAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStacksAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStacksResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStacksAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStacksAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) DescribeStacksAsync(
	request *DescribeStacksRequest,
	callback chan<- DescribeStacksAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "describeStacks",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
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

	go p.describeStacksAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) DescribeStacks(
	request *DescribeStacksRequest,
) (*DescribeStacksResult, error) {
	callback := make(chan DescribeStacksAsyncResult, 1)
	go p.DescribeStacksAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) preCreateStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PreCreateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreCreateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreCreateStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreCreateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PreCreateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) PreCreateStackAsync(
	request *PreCreateStackRequest,
	callback chan<- PreCreateStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "preCreateStack",
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

	go p.preCreateStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) PreCreateStack(
	request *PreCreateStackRequest,
) (*PreCreateStackResult, error) {
	callback := make(chan PreCreateStackAsyncResult, 1)
	go p.PreCreateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) createStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) CreateStackAsync(
	request *CreateStackRequest,
	callback chan<- CreateStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "createStack",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go p.createStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) CreateStack(
	request *CreateStackRequest,
) (*CreateStackResult, error) {
	callback := make(chan CreateStackAsyncResult, 1)
	go p.CreateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) createStackFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateStackFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStackFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStackFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStackFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateStackFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) CreateStackFromGitHubAsync(
	request *CreateStackFromGitHubRequest,
	callback chan<- CreateStackFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "createStackFromGitHub",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
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

	go p.createStackFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) CreateStackFromGitHub(
	request *CreateStackFromGitHubRequest,
) (*CreateStackFromGitHubResult, error) {
	callback := make(chan CreateStackFromGitHubAsyncResult, 1)
	go p.CreateStackFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) preValidateAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PreValidateAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreValidateAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreValidateResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreValidateAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PreValidateAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) PreValidateAsync(
	request *PreValidateRequest,
	callback chan<- PreValidateAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "preValidate",
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

	go p.preValidateAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) PreValidate(
	request *PreValidateRequest,
) (*PreValidateResult, error) {
	callback := make(chan PreValidateAsyncResult, 1)
	go p.PreValidateAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) validateAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ValidateAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ValidateAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ValidateResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ValidateAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ValidateAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) ValidateAsync(
	request *ValidateRequest,
	callback chan<- ValidateAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "validate",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go p.validateAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) Validate(
	request *ValidateRequest,
) (*ValidateResult, error) {
	callback := make(chan ValidateAsyncResult, 1)
	go p.ValidateAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) getStackStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStackStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStackStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStackStatusResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStackStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStackStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) GetStackStatusAsync(
	request *GetStackStatusRequest,
	callback chan<- GetStackStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "getStackStatus",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.getStackStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) GetStackStatus(
	request *GetStackStatusRequest,
) (*GetStackStatusResult, error) {
	callback := make(chan GetStackStatusAsyncResult, 1)
	go p.GetStackStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) getStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) GetStackAsync(
	request *GetStackRequest,
	callback chan<- GetStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "getStack",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.getStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) GetStack(
	request *GetStackRequest,
) (*GetStackResult, error) {
	callback := make(chan GetStackAsyncResult, 1)
	go p.GetStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) preUpdateStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PreUpdateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PreUpdateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) PreUpdateStackAsync(
	request *PreUpdateStackRequest,
	callback chan<- PreUpdateStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "preUpdateStack",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.preUpdateStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) PreUpdateStack(
	request *PreUpdateStackRequest,
) (*PreUpdateStackResult, error) {
	callback := make(chan PreUpdateStackAsyncResult, 1)
	go p.PreUpdateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) updateStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) UpdateStackAsync(
	request *UpdateStackRequest,
	callback chan<- UpdateStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "updateStack",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go p.updateStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) UpdateStack(
	request *UpdateStackRequest,
) (*UpdateStackResult, error) {
	callback := make(chan UpdateStackAsyncResult, 1)
	go p.UpdateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) preChangeSetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PreChangeSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreChangeSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreChangeSetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreChangeSetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PreChangeSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) PreChangeSetAsync(
	request *PreChangeSetRequest,
	callback chan<- PreChangeSetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "preChangeSet",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.preChangeSetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) PreChangeSet(
	request *PreChangeSetRequest,
) (*PreChangeSetResult, error) {
	callback := make(chan PreChangeSetAsyncResult, 1)
	go p.PreChangeSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) changeSetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ChangeSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ChangeSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ChangeSetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ChangeSetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ChangeSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) ChangeSetAsync(
	request *ChangeSetRequest,
	callback chan<- ChangeSetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "changeSet",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go p.changeSetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) ChangeSet(
	request *ChangeSetRequest,
) (*ChangeSetResult, error) {
	callback := make(chan ChangeSetAsyncResult, 1)
	go p.ChangeSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) updateStackFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateStackFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStackFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStackFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStackFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateStackFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) UpdateStackFromGitHubAsync(
	request *UpdateStackFromGitHubRequest,
	callback chan<- UpdateStackFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "updateStackFromGitHub",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
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

	go p.updateStackFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) UpdateStackFromGitHub(
	request *UpdateStackFromGitHubRequest,
) (*UpdateStackFromGitHubResult, error) {
	callback := make(chan UpdateStackFromGitHubAsyncResult, 1)
	go p.UpdateStackFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) deleteStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) DeleteStackAsync(
	request *DeleteStackRequest,
	callback chan<- DeleteStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "deleteStack",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.deleteStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) DeleteStack(
	request *DeleteStackRequest,
) (*DeleteStackResult, error) {
	callback := make(chan DeleteStackAsyncResult, 1)
	go p.DeleteStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) forceDeleteStackAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ForceDeleteStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ForceDeleteStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ForceDeleteStackResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ForceDeleteStackAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ForceDeleteStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) ForceDeleteStackAsync(
	request *ForceDeleteStackRequest,
	callback chan<- ForceDeleteStackAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "forceDeleteStack",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.forceDeleteStackAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) ForceDeleteStack(
	request *ForceDeleteStackRequest,
) (*ForceDeleteStackResult, error) {
	callback := make(chan ForceDeleteStackAsyncResult, 1)
	go p.ForceDeleteStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) deleteStackResourcesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStackResourcesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStackResourcesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStackResourcesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStackResourcesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteStackResourcesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) DeleteStackResourcesAsync(
	request *DeleteStackResourcesRequest,
	callback chan<- DeleteStackResourcesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "deleteStackResources",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.deleteStackResourcesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) DeleteStackResources(
	request *DeleteStackResourcesRequest,
) (*DeleteStackResourcesResult, error) {
	callback := make(chan DeleteStackResourcesAsyncResult, 1)
	go p.DeleteStackResourcesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) deleteStackEntityAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStackEntityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStackEntityAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStackEntityResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStackEntityAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteStackEntityAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) DeleteStackEntityAsync(
	request *DeleteStackEntityRequest,
	callback chan<- DeleteStackEntityAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "deleteStackEntity",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.deleteStackEntityAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) DeleteStackEntity(
	request *DeleteStackEntityRequest,
) (*DeleteStackEntityResult, error) {
	callback := make(chan DeleteStackEntityAsyncResult, 1)
	go p.DeleteStackEntityAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) getServiceVersionAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetServiceVersionResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "stack",
			"function":    "getServiceVersion",
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

	go p.getServiceVersionAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) GetServiceVersion(
	request *GetServiceVersionRequest,
) (*GetServiceVersionResult, error) {
	callback := make(chan GetServiceVersionAsyncResult, 1)
	go p.GetServiceVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) describeResourcesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeResourcesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeResourcesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeResourcesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeResourcesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeResourcesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) DescribeResourcesAsync(
	request *DescribeResourcesRequest,
	callback chan<- DescribeResourcesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "resource",
			"function":    "describeResources",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.describeResourcesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) DescribeResources(
	request *DescribeResourcesRequest,
) (*DescribeResourcesResult, error) {
	callback := make(chan DescribeResourcesAsyncResult, 1)
	go p.DescribeResourcesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) getResourceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetResourceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetResourceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetResourceResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetResourceAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetResourceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) GetResourceAsync(
	request *GetResourceRequest,
	callback chan<- GetResourceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "resource",
			"function":    "getResource",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
	}
	if request.ResourceName != nil && *request.ResourceName != "" {
		bodies["resourceName"] = *request.ResourceName
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

	go p.getResourceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) GetResource(
	request *GetResourceRequest,
) (*GetResourceResult, error) {
	callback := make(chan GetResourceAsyncResult, 1)
	go p.GetResourceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) describeEventsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeEventsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEventsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEventsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEventsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeEventsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) DescribeEventsAsync(
	request *DescribeEventsRequest,
	callback chan<- DescribeEventsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "event",
			"function":    "describeEvents",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

	go p.describeEventsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) DescribeEvents(
	request *DescribeEventsRequest,
) (*DescribeEventsResult, error) {
	callback := make(chan DescribeEventsAsyncResult, 1)
	go p.DescribeEventsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) getEventAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetEventAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEventAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEventResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEventAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetEventAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployWebSocketClient) GetEventAsync(
	request *GetEventRequest,
	callback chan<- GetEventAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "event",
			"function":    "getEvent",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
	}
	if request.EventName != nil && *request.EventName != "" {
		bodies["eventName"] = *request.EventName
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

	go p.getEventAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DeployWebSocketClient) GetEvent(
	request *GetEventRequest,
) (*GetEventResult, error) {
	callback := make(chan GetEventAsyncResult, 1)
	go p.GetEventAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DeployWebSocketClient) describeOutputsAsyncHandler(
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

func (p Gs2DeployWebSocketClient) DescribeOutputsAsync(
	request *DescribeOutputsRequest,
	callback chan<- DescribeOutputsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "output",
			"function":    "describeOutputs",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

func (p Gs2DeployWebSocketClient) DescribeOutputs(
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

func (p Gs2DeployWebSocketClient) getOutputAsyncHandler(
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

func (p Gs2DeployWebSocketClient) GetOutputAsync(
	request *GetOutputRequest,
	callback chan<- GetOutputAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "deploy",
			"component":   "output",
			"function":    "getOutput",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StackName != nil && *request.StackName != "" {
		bodies["stackName"] = *request.StackName
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

func (p Gs2DeployWebSocketClient) GetOutput(
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

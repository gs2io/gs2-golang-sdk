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

package skillTree

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2SkillTreeWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2SkillTreeWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2SkillTreeWebSocketClient) describeNamespacesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeNamespacesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeNamespacesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeNamespacesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeNamespacesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "namespace",
			"function":    "describeNamespaces",
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

	go p.describeNamespacesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) DescribeNamespaces(
	request *DescribeNamespacesRequest,
) (*DescribeNamespacesResult, error) {
	callback := make(chan DescribeNamespacesAsyncResult, 1)
	go p.DescribeNamespacesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) createNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateNamespaceResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "namespace",
			"function":    "createNamespace",
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
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
	}
	if request.ReleaseScript != nil {
		bodies["releaseScript"] = request.ReleaseScript.ToDict()
	}
	if request.RestrainScript != nil {
		bodies["restrainScript"] = request.RestrainScript.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) CreateNamespace(
	request *CreateNamespaceRequest,
) (*CreateNamespaceResult, error) {
	callback := make(chan CreateNamespaceAsyncResult, 1)
	go p.CreateNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getNamespaceStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetNamespaceStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetNamespaceStatusResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetNamespaceStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetNamespaceStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "namespace",
			"function":    "getNamespaceStatus",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getNamespaceStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetNamespaceStatus(
	request *GetNamespaceStatusRequest,
) (*GetNamespaceStatusResult, error) {
	callback := make(chan GetNamespaceStatusAsyncResult, 1)
	go p.GetNamespaceStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetNamespaceResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "namespace",
			"function":    "getNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetNamespace(
	request *GetNamespaceRequest,
) (*GetNamespaceResult, error) {
	callback := make(chan GetNamespaceAsyncResult, 1)
	go p.GetNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) updateNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateNamespaceResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "namespace",
			"function":    "updateNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
	}
	if request.ReleaseScript != nil {
		bodies["releaseScript"] = request.ReleaseScript.ToDict()
	}
	if request.RestrainScript != nil {
		bodies["restrainScript"] = request.RestrainScript.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) UpdateNamespace(
	request *UpdateNamespaceRequest,
) (*UpdateNamespaceResult, error) {
	callback := make(chan UpdateNamespaceAsyncResult, 1)
	go p.UpdateNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) deleteNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteNamespaceResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "namespace",
			"function":    "deleteNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) DeleteNamespace(
	request *DeleteNamespaceRequest,
) (*DeleteNamespaceResult, error) {
	callback := make(chan DeleteNamespaceAsyncResult, 1)
	go p.DeleteNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) describeStatusModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStatusModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStatusModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) DescribeStatusModelsAsync(
	request *DescribeStatusModelsRequest,
	callback chan<- DescribeStatusModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModel",
			"function":    "describeStatusModels",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.describeStatusModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) DescribeStatusModels(
	request *DescribeStatusModelsRequest,
) (*DescribeStatusModelsResult, error) {
	callback := make(chan DescribeStatusModelsAsyncResult, 1)
	go p.DescribeStatusModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getStatusModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStatusModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetStatusModelAsync(
	request *GetStatusModelRequest,
	callback chan<- GetStatusModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModel",
			"function":    "getStatusModel",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		bodies["nodeModelName"] = *request.NodeModelName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getStatusModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetStatusModel(
	request *GetStatusModelRequest,
) (*GetStatusModelResult, error) {
	callback := make(chan GetStatusModelAsyncResult, 1)
	go p.GetStatusModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) describeStatusModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStatusModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStatusModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) DescribeStatusModelMastersAsync(
	request *DescribeStatusModelMastersRequest,
	callback chan<- DescribeStatusModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModelMaster",
			"function":    "describeStatusModelMasters",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
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

	go p.describeStatusModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) DescribeStatusModelMasters(
	request *DescribeStatusModelMastersRequest,
) (*DescribeStatusModelMastersResult, error) {
	callback := make(chan DescribeStatusModelMastersAsyncResult, 1)
	go p.DescribeStatusModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) createStatusModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStatusModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) CreateStatusModelMasterAsync(
	request *CreateStatusModelMasterRequest,
	callback chan<- CreateStatusModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModelMaster",
			"function":    "createStatusModelMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ReleaseConsumeActions != nil {
		var _releaseConsumeActions []interface{}
		for _, item := range request.ReleaseConsumeActions {
			_releaseConsumeActions = append(_releaseConsumeActions, item)
		}
		bodies["releaseConsumeActions"] = _releaseConsumeActions
	}
	if request.RestrainReturnRate != nil {
		bodies["restrainReturnRate"] = *request.RestrainReturnRate
	}
	if request.PremiseNodeNames != nil {
		var _premiseNodeNames []interface{}
		for _, item := range request.PremiseNodeNames {
			_premiseNodeNames = append(_premiseNodeNames, item)
		}
		bodies["premiseNodeNames"] = _premiseNodeNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createStatusModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) CreateStatusModelMaster(
	request *CreateStatusModelMasterRequest,
) (*CreateStatusModelMasterResult, error) {
	callback := make(chan CreateStatusModelMasterAsyncResult, 1)
	go p.CreateStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getStatusModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetStatusModelMasterAsync(
	request *GetStatusModelMasterRequest,
	callback chan<- GetStatusModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModelMaster",
			"function":    "getStatusModelMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		bodies["nodeModelName"] = *request.NodeModelName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getStatusModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetStatusModelMaster(
	request *GetStatusModelMasterRequest,
) (*GetStatusModelMasterResult, error) {
	callback := make(chan GetStatusModelMasterAsyncResult, 1)
	go p.GetStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) updateStatusModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStatusModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) UpdateStatusModelMasterAsync(
	request *UpdateStatusModelMasterRequest,
	callback chan<- UpdateStatusModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModelMaster",
			"function":    "updateStatusModelMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		bodies["nodeModelName"] = *request.NodeModelName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ReleaseConsumeActions != nil {
		var _releaseConsumeActions []interface{}
		for _, item := range request.ReleaseConsumeActions {
			_releaseConsumeActions = append(_releaseConsumeActions, item)
		}
		bodies["releaseConsumeActions"] = _releaseConsumeActions
	}
	if request.RestrainReturnRate != nil {
		bodies["restrainReturnRate"] = *request.RestrainReturnRate
	}
	if request.PremiseNodeNames != nil {
		var _premiseNodeNames []interface{}
		for _, item := range request.PremiseNodeNames {
			_premiseNodeNames = append(_premiseNodeNames, item)
		}
		bodies["premiseNodeNames"] = _premiseNodeNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateStatusModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) UpdateStatusModelMaster(
	request *UpdateStatusModelMasterRequest,
) (*UpdateStatusModelMasterResult, error) {
	callback := make(chan UpdateStatusModelMasterAsyncResult, 1)
	go p.UpdateStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) deleteStatusModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStatusModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) DeleteStatusModelMasterAsync(
	request *DeleteStatusModelMasterRequest,
	callback chan<- DeleteStatusModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "nodeModelMaster",
			"function":    "deleteStatusModelMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		bodies["nodeModelName"] = *request.NodeModelName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteStatusModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) DeleteStatusModelMaster(
	request *DeleteStatusModelMasterRequest,
) (*DeleteStatusModelMasterResult, error) {
	callback := make(chan DeleteStatusModelMasterAsyncResult, 1)
	go p.DeleteStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) markReleaseByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- MarkReleaseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReleaseByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReleaseByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReleaseByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- MarkReleaseByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) MarkReleaseByUserIdAsync(
	request *MarkReleaseByUserIdRequest,
	callback chan<- MarkReleaseByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "markReleaseByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.markReleaseByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) MarkReleaseByUserId(
	request *MarkReleaseByUserIdRequest,
) (*MarkReleaseByUserIdResult, error) {
	callback := make(chan MarkReleaseByUserIdAsyncResult, 1)
	go p.MarkReleaseByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) releaseAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReleaseAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReleaseAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReleaseResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReleaseAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReleaseAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) ReleaseAsync(
	request *ReleaseRequest,
	callback chan<- ReleaseAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "release",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.releaseAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) Release(
	request *ReleaseRequest,
) (*ReleaseResult, error) {
	callback := make(chan ReleaseAsyncResult, 1)
	go p.ReleaseAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) releaseByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReleaseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReleaseByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReleaseByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReleaseByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReleaseByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) ReleaseByUserIdAsync(
	request *ReleaseByUserIdRequest,
	callback chan<- ReleaseByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "releaseByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.releaseByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) ReleaseByUserId(
	request *ReleaseByUserIdRequest,
) (*ReleaseByUserIdResult, error) {
	callback := make(chan ReleaseByUserIdAsyncResult, 1)
	go p.ReleaseByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) markRestrainByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- MarkRestrainByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkRestrainByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkRestrainByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkRestrainByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- MarkRestrainByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) MarkRestrainByUserIdAsync(
	request *MarkRestrainByUserIdRequest,
	callback chan<- MarkRestrainByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "markRestrainByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.markRestrainByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) MarkRestrainByUserId(
	request *MarkRestrainByUserIdRequest,
) (*MarkRestrainByUserIdResult, error) {
	callback := make(chan MarkRestrainByUserIdAsyncResult, 1)
	go p.MarkRestrainByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) restrainAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RestrainAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RestrainAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RestrainResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RestrainAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RestrainAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) RestrainAsync(
	request *RestrainRequest,
	callback chan<- RestrainAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "restrain",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.restrainAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) Restrain(
	request *RestrainRequest,
) (*RestrainResult, error) {
	callback := make(chan RestrainAsyncResult, 1)
	go p.RestrainAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) restrainByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RestrainByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RestrainByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RestrainByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RestrainByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RestrainByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) RestrainByUserIdAsync(
	request *RestrainByUserIdRequest,
	callback chan<- RestrainByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "restrainByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.restrainByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) RestrainByUserId(
	request *RestrainByUserIdRequest,
) (*RestrainByUserIdResult, error) {
	callback := make(chan RestrainByUserIdAsyncResult, 1)
	go p.RestrainByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetStatusAsync(
	request *GetStatusRequest,
	callback chan<- GetStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "getStatus",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetStatus(
	request *GetStatusRequest,
) (*GetStatusResult, error) {
	callback := make(chan GetStatusAsyncResult, 1)
	go p.GetStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetStatusByUserIdAsync(
	request *GetStatusByUserIdRequest,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "getStatusByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetStatusByUserId(
	request *GetStatusByUserIdRequest,
) (*GetStatusByUserIdResult, error) {
	callback := make(chan GetStatusByUserIdAsyncResult, 1)
	go p.GetStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) resetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ResetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ResetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) ResetAsync(
	request *ResetRequest,
	callback chan<- ResetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "reset",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.resetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) Reset(
	request *ResetRequest,
) (*ResetResult, error) {
	callback := make(chan ResetAsyncResult, 1)
	go p.ResetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) resetByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ResetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ResetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) ResetByUserIdAsync(
	request *ResetByUserIdRequest,
	callback chan<- ResetByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "resetByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.resetByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) ResetByUserId(
	request *ResetByUserIdRequest,
) (*ResetByUserIdResult, error) {
	callback := make(chan ResetByUserIdAsyncResult, 1)
	go p.ResetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) markReleaseByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- MarkReleaseByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReleaseByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReleaseByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReleaseByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- MarkReleaseByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) MarkReleaseByStampSheetAsync(
	request *MarkReleaseByStampSheetRequest,
	callback chan<- MarkReleaseByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "markReleaseByStampSheet",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.markReleaseByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) MarkReleaseByStampSheet(
	request *MarkReleaseByStampSheetRequest,
) (*MarkReleaseByStampSheetResult, error) {
	callback := make(chan MarkReleaseByStampSheetAsyncResult, 1)
	go p.MarkReleaseByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) markRestrainByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- MarkRestrainByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkRestrainByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkRestrainByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkRestrainByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- MarkRestrainByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) MarkRestrainByStampTaskAsync(
	request *MarkRestrainByStampTaskRequest,
	callback chan<- MarkRestrainByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "status",
			"function":    "markRestrainByStampTask",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.markRestrainByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) MarkRestrainByStampTask(
	request *MarkRestrainByStampTaskRequest,
) (*MarkRestrainByStampTaskResult, error) {
	callback := make(chan MarkRestrainByStampTaskAsyncResult, 1)
	go p.MarkRestrainByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) exportMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExportMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExportMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ExportMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ExportMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "currentTreeMaster",
			"function":    "exportMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.exportMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) ExportMaster(
	request *ExportMasterRequest,
) (*ExportMasterResult, error) {
	callback := make(chan ExportMasterAsyncResult, 1)
	go p.ExportMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) getCurrentStatusMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentStatusMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentStatusMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentStatusMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentStatusMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCurrentStatusMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) GetCurrentStatusMasterAsync(
	request *GetCurrentStatusMasterRequest,
	callback chan<- GetCurrentStatusMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "currentTreeMaster",
			"function":    "getCurrentStatusMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getCurrentStatusMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) GetCurrentStatusMaster(
	request *GetCurrentStatusMasterRequest,
) (*GetCurrentStatusMasterResult, error) {
	callback := make(chan GetCurrentStatusMasterAsyncResult, 1)
	go p.GetCurrentStatusMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) updateCurrentStatusMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentStatusMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStatusMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStatusMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentStatusMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentStatusMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) UpdateCurrentStatusMasterAsync(
	request *UpdateCurrentStatusMasterRequest,
	callback chan<- UpdateCurrentStatusMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "currentTreeMaster",
			"function":    "updateCurrentStatusMaster",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateCurrentStatusMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) UpdateCurrentStatusMaster(
	request *UpdateCurrentStatusMasterRequest,
) (*UpdateCurrentStatusMasterResult, error) {
	callback := make(chan UpdateCurrentStatusMasterAsyncResult, 1)
	go p.UpdateCurrentStatusMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2SkillTreeWebSocketClient) updateCurrentStatusMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentStatusMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStatusMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeWebSocketClient) UpdateCurrentStatusMasterFromGitHubAsync(
	request *UpdateCurrentStatusMasterFromGitHubRequest,
	callback chan<- UpdateCurrentStatusMasterFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "skill_tree",
			"component":   "currentTreeMaster",
			"function":    "updateCurrentStatusMasterFromGitHub",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateCurrentStatusMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeWebSocketClient) UpdateCurrentStatusMasterFromGitHub(
	request *UpdateCurrentStatusMasterFromGitHubRequest,
) (*UpdateCurrentStatusMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentStatusMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentStatusMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

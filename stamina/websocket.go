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

package stamina

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2StaminaWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2StaminaWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2StaminaWebSocketClient) describeNamespacesAsyncHandler(
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
	callback <- DescribeNamespacesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "namespace",
    		"function": "describeNamespaces",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeNamespacesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeNamespaces(
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

func (p Gs2StaminaWebSocketClient) createNamespaceAsyncHandler(
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
	callback <- CreateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "namespace",
    		"function": "createNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.OverflowTriggerScript != nil {
        bodies["overflowTriggerScript"] = request.OverflowTriggerScript.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) CreateNamespace(
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

func (p Gs2StaminaWebSocketClient) getNamespaceStatusAsyncHandler(
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
	callback <- GetNamespaceStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "namespace",
    		"function": "getNamespaceStatus",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getNamespaceStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetNamespaceStatus(
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

func (p Gs2StaminaWebSocketClient) getNamespaceAsyncHandler(
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
	callback <- GetNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "namespace",
    		"function": "getNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetNamespace(
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

func (p Gs2StaminaWebSocketClient) updateNamespaceAsyncHandler(
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
	callback <- UpdateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "namespace",
    		"function": "updateNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.OverflowTriggerScript != nil {
        bodies["overflowTriggerScript"] = request.OverflowTriggerScript.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateNamespace(
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

func (p Gs2StaminaWebSocketClient) deleteNamespaceAsyncHandler(
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
	callback <- DeleteNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "namespace",
    		"function": "deleteNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DeleteNamespace(
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

func (p Gs2StaminaWebSocketClient) describeStaminaModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStaminaModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminaModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminaModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeStaminaModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeStaminaModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeStaminaModelMastersAsync(
	request *DescribeStaminaModelMastersRequest,
	callback chan<- DescribeStaminaModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModelMaster",
    		"function": "describeStaminaModelMasters",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeStaminaModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeStaminaModelMasters(
	request *DescribeStaminaModelMastersRequest,
) (*DescribeStaminaModelMastersResult, error) {
	callback := make(chan DescribeStaminaModelMastersAsyncResult, 1)
	go p.DescribeStaminaModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) createStaminaModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStaminaModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateStaminaModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) CreateStaminaModelMasterAsync(
	request *CreateStaminaModelMasterRequest,
	callback chan<- CreateStaminaModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModelMaster",
    		"function": "createStaminaModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.RecoverIntervalMinutes != nil {
        bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
    }
    if request.RecoverValue != nil {
        bodies["recoverValue"] = *request.RecoverValue
    }
    if request.InitialCapacity != nil {
        bodies["initialCapacity"] = *request.InitialCapacity
    }
    if request.IsOverflow != nil {
        bodies["isOverflow"] = *request.IsOverflow
    }
    if request.MaxCapacity != nil {
        bodies["maxCapacity"] = *request.MaxCapacity
    }
    if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
        bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
    }
    if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
        bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
    }
    if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
        bodies["recoverValueTableName"] = *request.RecoverValueTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createStaminaModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) CreateStaminaModelMaster(
	request *CreateStaminaModelMasterRequest,
) (*CreateStaminaModelMasterResult, error) {
	callback := make(chan CreateStaminaModelMasterAsyncResult, 1)
	go p.CreateStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getStaminaModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStaminaModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetStaminaModelMasterAsync(
	request *GetStaminaModelMasterRequest,
	callback chan<- GetStaminaModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModelMaster",
    		"function": "getStaminaModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getStaminaModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetStaminaModelMaster(
	request *GetStaminaModelMasterRequest,
) (*GetStaminaModelMasterResult, error) {
	callback := make(chan GetStaminaModelMasterAsyncResult, 1)
	go p.GetStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateStaminaModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStaminaModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateStaminaModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateStaminaModelMasterAsync(
	request *UpdateStaminaModelMasterRequest,
	callback chan<- UpdateStaminaModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModelMaster",
    		"function": "updateStaminaModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.RecoverIntervalMinutes != nil {
        bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
    }
    if request.RecoverValue != nil {
        bodies["recoverValue"] = *request.RecoverValue
    }
    if request.InitialCapacity != nil {
        bodies["initialCapacity"] = *request.InitialCapacity
    }
    if request.IsOverflow != nil {
        bodies["isOverflow"] = *request.IsOverflow
    }
    if request.MaxCapacity != nil {
        bodies["maxCapacity"] = *request.MaxCapacity
    }
    if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
        bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
    }
    if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
        bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
    }
    if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
        bodies["recoverValueTableName"] = *request.RecoverValueTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateStaminaModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateStaminaModelMaster(
	request *UpdateStaminaModelMasterRequest,
) (*UpdateStaminaModelMasterResult, error) {
	callback := make(chan UpdateStaminaModelMasterAsyncResult, 1)
	go p.UpdateStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) deleteStaminaModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStaminaModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteStaminaModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DeleteStaminaModelMasterAsync(
	request *DeleteStaminaModelMasterRequest,
	callback chan<- DeleteStaminaModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModelMaster",
    		"function": "deleteStaminaModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteStaminaModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DeleteStaminaModelMaster(
	request *DeleteStaminaModelMasterRequest,
) (*DeleteStaminaModelMasterResult, error) {
	callback := make(chan DeleteStaminaModelMasterAsyncResult, 1)
	go p.DeleteStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) describeMaxStaminaTableMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMaxStaminaTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMaxStaminaTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMaxStaminaTableMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMaxStaminaTableMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMaxStaminaTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeMaxStaminaTableMastersAsync(
	request *DescribeMaxStaminaTableMastersRequest,
	callback chan<- DescribeMaxStaminaTableMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "maxStaminaTableMaster",
    		"function": "describeMaxStaminaTableMasters",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeMaxStaminaTableMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeMaxStaminaTableMasters(
	request *DescribeMaxStaminaTableMastersRequest,
) (*DescribeMaxStaminaTableMastersResult, error) {
	callback := make(chan DescribeMaxStaminaTableMastersAsyncResult, 1)
	go p.DescribeMaxStaminaTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) createMaxStaminaTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateMaxStaminaTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateMaxStaminaTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) CreateMaxStaminaTableMasterAsync(
	request *CreateMaxStaminaTableMasterRequest,
	callback chan<- CreateMaxStaminaTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "maxStaminaTableMaster",
    		"function": "createMaxStaminaTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
        bodies["experienceModelId"] = *request.ExperienceModelId
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createMaxStaminaTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) CreateMaxStaminaTableMaster(
	request *CreateMaxStaminaTableMasterRequest,
) (*CreateMaxStaminaTableMasterResult, error) {
	callback := make(chan CreateMaxStaminaTableMasterAsyncResult, 1)
	go p.CreateMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getMaxStaminaTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMaxStaminaTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMaxStaminaTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetMaxStaminaTableMasterAsync(
	request *GetMaxStaminaTableMasterRequest,
	callback chan<- GetMaxStaminaTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "maxStaminaTableMaster",
    		"function": "getMaxStaminaTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
        bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getMaxStaminaTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetMaxStaminaTableMaster(
	request *GetMaxStaminaTableMasterRequest,
) (*GetMaxStaminaTableMasterResult, error) {
	callback := make(chan GetMaxStaminaTableMasterAsyncResult, 1)
	go p.GetMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateMaxStaminaTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMaxStaminaTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateMaxStaminaTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateMaxStaminaTableMasterAsync(
	request *UpdateMaxStaminaTableMasterRequest,
	callback chan<- UpdateMaxStaminaTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "maxStaminaTableMaster",
    		"function": "updateMaxStaminaTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
        bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
        bodies["experienceModelId"] = *request.ExperienceModelId
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateMaxStaminaTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateMaxStaminaTableMaster(
	request *UpdateMaxStaminaTableMasterRequest,
) (*UpdateMaxStaminaTableMasterResult, error) {
	callback := make(chan UpdateMaxStaminaTableMasterAsyncResult, 1)
	go p.UpdateMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) deleteMaxStaminaTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMaxStaminaTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteMaxStaminaTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DeleteMaxStaminaTableMasterAsync(
	request *DeleteMaxStaminaTableMasterRequest,
	callback chan<- DeleteMaxStaminaTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "maxStaminaTableMaster",
    		"function": "deleteMaxStaminaTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
        bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteMaxStaminaTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DeleteMaxStaminaTableMaster(
	request *DeleteMaxStaminaTableMasterRequest,
) (*DeleteMaxStaminaTableMasterResult, error) {
	callback := make(chan DeleteMaxStaminaTableMasterAsyncResult, 1)
	go p.DeleteMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) describeRecoverIntervalTableMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRecoverIntervalTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRecoverIntervalTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRecoverIntervalTableMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRecoverIntervalTableMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRecoverIntervalTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeRecoverIntervalTableMastersAsync(
	request *DescribeRecoverIntervalTableMastersRequest,
	callback chan<- DescribeRecoverIntervalTableMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverIntervalTableMaster",
    		"function": "describeRecoverIntervalTableMasters",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeRecoverIntervalTableMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeRecoverIntervalTableMasters(
	request *DescribeRecoverIntervalTableMastersRequest,
) (*DescribeRecoverIntervalTableMastersResult, error) {
	callback := make(chan DescribeRecoverIntervalTableMastersAsyncResult, 1)
	go p.DescribeRecoverIntervalTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) createRecoverIntervalTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRecoverIntervalTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateRecoverIntervalTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) CreateRecoverIntervalTableMasterAsync(
	request *CreateRecoverIntervalTableMasterRequest,
	callback chan<- CreateRecoverIntervalTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverIntervalTableMaster",
    		"function": "createRecoverIntervalTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
        bodies["experienceModelId"] = *request.ExperienceModelId
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createRecoverIntervalTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) CreateRecoverIntervalTableMaster(
	request *CreateRecoverIntervalTableMasterRequest,
) (*CreateRecoverIntervalTableMasterResult, error) {
	callback := make(chan CreateRecoverIntervalTableMasterAsyncResult, 1)
	go p.CreateRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getRecoverIntervalTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRecoverIntervalTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRecoverIntervalTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetRecoverIntervalTableMasterAsync(
	request *GetRecoverIntervalTableMasterRequest,
	callback chan<- GetRecoverIntervalTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverIntervalTableMaster",
    		"function": "getRecoverIntervalTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
        bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRecoverIntervalTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetRecoverIntervalTableMaster(
	request *GetRecoverIntervalTableMasterRequest,
) (*GetRecoverIntervalTableMasterResult, error) {
	callback := make(chan GetRecoverIntervalTableMasterAsyncResult, 1)
	go p.GetRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateRecoverIntervalTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRecoverIntervalTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateRecoverIntervalTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateRecoverIntervalTableMasterAsync(
	request *UpdateRecoverIntervalTableMasterRequest,
	callback chan<- UpdateRecoverIntervalTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverIntervalTableMaster",
    		"function": "updateRecoverIntervalTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
        bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
        bodies["experienceModelId"] = *request.ExperienceModelId
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateRecoverIntervalTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateRecoverIntervalTableMaster(
	request *UpdateRecoverIntervalTableMasterRequest,
) (*UpdateRecoverIntervalTableMasterResult, error) {
	callback := make(chan UpdateRecoverIntervalTableMasterAsyncResult, 1)
	go p.UpdateRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) deleteRecoverIntervalTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRecoverIntervalTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRecoverIntervalTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DeleteRecoverIntervalTableMasterAsync(
	request *DeleteRecoverIntervalTableMasterRequest,
	callback chan<- DeleteRecoverIntervalTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverIntervalTableMaster",
    		"function": "deleteRecoverIntervalTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
        bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteRecoverIntervalTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DeleteRecoverIntervalTableMaster(
	request *DeleteRecoverIntervalTableMasterRequest,
) (*DeleteRecoverIntervalTableMasterResult, error) {
	callback := make(chan DeleteRecoverIntervalTableMasterAsyncResult, 1)
	go p.DeleteRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) describeRecoverValueTableMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRecoverValueTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRecoverValueTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRecoverValueTableMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRecoverValueTableMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRecoverValueTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeRecoverValueTableMastersAsync(
	request *DescribeRecoverValueTableMastersRequest,
	callback chan<- DescribeRecoverValueTableMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverValueTableMaster",
    		"function": "describeRecoverValueTableMasters",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeRecoverValueTableMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeRecoverValueTableMasters(
	request *DescribeRecoverValueTableMastersRequest,
) (*DescribeRecoverValueTableMastersResult, error) {
	callback := make(chan DescribeRecoverValueTableMastersAsyncResult, 1)
	go p.DescribeRecoverValueTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) createRecoverValueTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRecoverValueTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateRecoverValueTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) CreateRecoverValueTableMasterAsync(
	request *CreateRecoverValueTableMasterRequest,
	callback chan<- CreateRecoverValueTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverValueTableMaster",
    		"function": "createRecoverValueTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
        bodies["experienceModelId"] = *request.ExperienceModelId
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createRecoverValueTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) CreateRecoverValueTableMaster(
	request *CreateRecoverValueTableMasterRequest,
) (*CreateRecoverValueTableMasterResult, error) {
	callback := make(chan CreateRecoverValueTableMasterAsyncResult, 1)
	go p.CreateRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getRecoverValueTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRecoverValueTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRecoverValueTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetRecoverValueTableMasterAsync(
	request *GetRecoverValueTableMasterRequest,
	callback chan<- GetRecoverValueTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverValueTableMaster",
    		"function": "getRecoverValueTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
        bodies["recoverValueTableName"] = *request.RecoverValueTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRecoverValueTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetRecoverValueTableMaster(
	request *GetRecoverValueTableMasterRequest,
) (*GetRecoverValueTableMasterResult, error) {
	callback := make(chan GetRecoverValueTableMasterAsyncResult, 1)
	go p.GetRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateRecoverValueTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRecoverValueTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateRecoverValueTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateRecoverValueTableMasterAsync(
	request *UpdateRecoverValueTableMasterRequest,
	callback chan<- UpdateRecoverValueTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverValueTableMaster",
    		"function": "updateRecoverValueTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
        bodies["recoverValueTableName"] = *request.RecoverValueTableName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
        bodies["experienceModelId"] = *request.ExperienceModelId
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateRecoverValueTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateRecoverValueTableMaster(
	request *UpdateRecoverValueTableMasterRequest,
) (*UpdateRecoverValueTableMasterResult, error) {
	callback := make(chan UpdateRecoverValueTableMasterAsyncResult, 1)
	go p.UpdateRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) deleteRecoverValueTableMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRecoverValueTableMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRecoverValueTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DeleteRecoverValueTableMasterAsync(
	request *DeleteRecoverValueTableMasterRequest,
	callback chan<- DeleteRecoverValueTableMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "recoverValueTableMaster",
    		"function": "deleteRecoverValueTableMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
        bodies["recoverValueTableName"] = *request.RecoverValueTableName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteRecoverValueTableMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DeleteRecoverValueTableMaster(
	request *DeleteRecoverValueTableMasterRequest,
) (*DeleteRecoverValueTableMasterResult, error) {
	callback := make(chan DeleteRecoverValueTableMasterAsyncResult, 1)
	go p.DeleteRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) exportMasterAsyncHandler(
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
	callback <- ExportMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "currentStaminaMaster",
    		"function": "exportMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.exportMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) ExportMaster(
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

func (p Gs2StaminaWebSocketClient) getCurrentStaminaMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentStaminaMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentStaminaMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentStaminaMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentStaminaMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentStaminaMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetCurrentStaminaMasterAsync(
	request *GetCurrentStaminaMasterRequest,
	callback chan<- GetCurrentStaminaMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "currentStaminaMaster",
    		"function": "getCurrentStaminaMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCurrentStaminaMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetCurrentStaminaMaster(
	request *GetCurrentStaminaMasterRequest,
) (*GetCurrentStaminaMasterResult, error) {
	callback := make(chan GetCurrentStaminaMasterAsyncResult, 1)
	go p.GetCurrentStaminaMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateCurrentStaminaMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentStaminaMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStaminaMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStaminaMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentStaminaMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentStaminaMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateCurrentStaminaMasterAsync(
	request *UpdateCurrentStaminaMasterRequest,
	callback chan<- UpdateCurrentStaminaMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "currentStaminaMaster",
    		"function": "updateCurrentStaminaMaster",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCurrentStaminaMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateCurrentStaminaMaster(
	request *UpdateCurrentStaminaMasterRequest,
) (*UpdateCurrentStaminaMasterResult, error) {
	callback := make(chan UpdateCurrentStaminaMasterAsyncResult, 1)
	go p.UpdateCurrentStaminaMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateCurrentStaminaMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentStaminaMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStaminaMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateCurrentStaminaMasterFromGitHubAsync(
	request *UpdateCurrentStaminaMasterFromGitHubRequest,
	callback chan<- UpdateCurrentStaminaMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "currentStaminaMaster",
    		"function": "updateCurrentStaminaMasterFromGitHub",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCurrentStaminaMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateCurrentStaminaMasterFromGitHub(
	request *UpdateCurrentStaminaMasterFromGitHubRequest,
) (*UpdateCurrentStaminaMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentStaminaMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentStaminaMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) describeStaminaModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStaminaModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminaModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminaModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeStaminaModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeStaminaModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeStaminaModelsAsync(
	request *DescribeStaminaModelsRequest,
	callback chan<- DescribeStaminaModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModel",
    		"function": "describeStaminaModels",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeStaminaModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeStaminaModels(
	request *DescribeStaminaModelsRequest,
) (*DescribeStaminaModelsResult, error) {
	callback := make(chan DescribeStaminaModelsAsyncResult, 1)
	go p.DescribeStaminaModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getStaminaModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStaminaModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStaminaModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStaminaModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetStaminaModelAsync(
	request *GetStaminaModelRequest,
	callback chan<- GetStaminaModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "staminaModel",
    		"function": "getStaminaModel",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getStaminaModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetStaminaModel(
	request *GetStaminaModelRequest,
) (*GetStaminaModelResult, error) {
	callback := make(chan GetStaminaModelAsyncResult, 1)
	go p.GetStaminaModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) describeStaminasAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStaminasAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminasAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminasResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeStaminasAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeStaminasAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeStaminasAsync(
	request *DescribeStaminasRequest,
	callback chan<- DescribeStaminasAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "describeStaminas",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeStaminasAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeStaminas(
	request *DescribeStaminasRequest,
) (*DescribeStaminasResult, error) {
	callback := make(chan DescribeStaminasAsyncResult, 1)
	go p.DescribeStaminasAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) describeStaminasByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStaminasByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminasByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminasByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeStaminasByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeStaminasByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DescribeStaminasByUserIdAsync(
	request *DescribeStaminasByUserIdRequest,
	callback chan<- DescribeStaminasByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "describeStaminasByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
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
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeStaminasByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DescribeStaminasByUserId(
	request *DescribeStaminasByUserIdRequest,
) (*DescribeStaminasByUserIdResult, error) {
	callback := make(chan DescribeStaminasByUserIdAsyncResult, 1)
	go p.DescribeStaminasByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getStaminaAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStaminaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStaminaAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStaminaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetStaminaAsync(
	request *GetStaminaRequest,
	callback chan<- GetStaminaAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "getStamina",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getStaminaAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetStamina(
	request *GetStaminaRequest,
) (*GetStaminaResult, error) {
	callback := make(chan GetStaminaAsyncResult, 1)
	go p.GetStaminaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) getStaminaByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStaminaByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) GetStaminaByUserIdAsync(
	request *GetStaminaByUserIdRequest,
	callback chan<- GetStaminaByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "getStaminaByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getStaminaByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) GetStaminaByUserId(
	request *GetStaminaByUserIdRequest,
) (*GetStaminaByUserIdResult, error) {
	callback := make(chan GetStaminaByUserIdAsyncResult, 1)
	go p.GetStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) updateStaminaByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStaminaByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateStaminaByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) UpdateStaminaByUserIdAsync(
	request *UpdateStaminaByUserIdRequest,
	callback chan<- UpdateStaminaByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "updateStaminaByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.Value != nil {
        bodies["value"] = *request.Value
    }
    if request.MaxValue != nil {
        bodies["maxValue"] = *request.MaxValue
    }
    if request.RecoverIntervalMinutes != nil {
        bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
    }
    if request.RecoverValue != nil {
        bodies["recoverValue"] = *request.RecoverValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.updateStaminaByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) UpdateStaminaByUserId(
	request *UpdateStaminaByUserIdRequest,
) (*UpdateStaminaByUserIdResult, error) {
	callback := make(chan UpdateStaminaByUserIdAsyncResult, 1)
	go p.UpdateStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) consumeStaminaAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ConsumeStaminaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeStaminaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeStaminaResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeStaminaAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ConsumeStaminaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) ConsumeStaminaAsync(
	request *ConsumeStaminaRequest,
	callback chan<- ConsumeStaminaAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "consumeStamina",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ConsumeValue != nil {
        bodies["consumeValue"] = *request.ConsumeValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.consumeStaminaAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) ConsumeStamina(
	request *ConsumeStaminaRequest,
) (*ConsumeStaminaResult, error) {
	callback := make(chan ConsumeStaminaAsyncResult, 1)
	go p.ConsumeStaminaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) consumeStaminaByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ConsumeStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeStaminaByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeStaminaByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ConsumeStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) ConsumeStaminaByUserIdAsync(
	request *ConsumeStaminaByUserIdRequest,
	callback chan<- ConsumeStaminaByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "consumeStaminaByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ConsumeValue != nil {
        bodies["consumeValue"] = *request.ConsumeValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.consumeStaminaByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) ConsumeStaminaByUserId(
	request *ConsumeStaminaByUserIdRequest,
) (*ConsumeStaminaByUserIdResult, error) {
	callback := make(chan ConsumeStaminaByUserIdAsyncResult, 1)
	go p.ConsumeStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) recoverStaminaByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RecoverStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RecoverStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RecoverStaminaByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RecoverStaminaByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RecoverStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) RecoverStaminaByUserIdAsync(
	request *RecoverStaminaByUserIdRequest,
	callback chan<- RecoverStaminaByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "recoverStaminaByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.RecoverValue != nil {
        bodies["recoverValue"] = *request.RecoverValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.recoverStaminaByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) RecoverStaminaByUserId(
	request *RecoverStaminaByUserIdRequest,
) (*RecoverStaminaByUserIdResult, error) {
	callback := make(chan RecoverStaminaByUserIdAsyncResult, 1)
	go p.RecoverStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) raiseMaxValueByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RaiseMaxValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RaiseMaxValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RaiseMaxValueByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RaiseMaxValueByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RaiseMaxValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) RaiseMaxValueByUserIdAsync(
	request *RaiseMaxValueByUserIdRequest,
	callback chan<- RaiseMaxValueByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "raiseMaxValueByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.RaiseValue != nil {
        bodies["raiseValue"] = *request.RaiseValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.raiseMaxValueByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) RaiseMaxValueByUserId(
	request *RaiseMaxValueByUserIdRequest,
) (*RaiseMaxValueByUserIdResult, error) {
	callback := make(chan RaiseMaxValueByUserIdAsyncResult, 1)
	go p.RaiseMaxValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setMaxValueByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetMaxValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaxValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaxValueByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetMaxValueByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetMaxValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetMaxValueByUserIdAsync(
	request *SetMaxValueByUserIdRequest,
	callback chan<- SetMaxValueByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setMaxValueByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.MaxValue != nil {
        bodies["maxValue"] = *request.MaxValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setMaxValueByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetMaxValueByUserId(
	request *SetMaxValueByUserIdRequest,
) (*SetMaxValueByUserIdResult, error) {
	callback := make(chan SetMaxValueByUserIdAsyncResult, 1)
	go p.SetMaxValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setRecoverIntervalByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRecoverIntervalByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverIntervalByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverIntervalByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRecoverIntervalByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRecoverIntervalByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetRecoverIntervalByUserIdAsync(
	request *SetRecoverIntervalByUserIdRequest,
	callback chan<- SetRecoverIntervalByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setRecoverIntervalByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.RecoverIntervalMinutes != nil {
        bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setRecoverIntervalByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetRecoverIntervalByUserId(
	request *SetRecoverIntervalByUserIdRequest,
) (*SetRecoverIntervalByUserIdResult, error) {
	callback := make(chan SetRecoverIntervalByUserIdAsyncResult, 1)
	go p.SetRecoverIntervalByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setRecoverValueByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRecoverValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverValueByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRecoverValueByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRecoverValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetRecoverValueByUserIdAsync(
	request *SetRecoverValueByUserIdRequest,
	callback chan<- SetRecoverValueByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setRecoverValueByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.RecoverValue != nil {
        bodies["recoverValue"] = *request.RecoverValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setRecoverValueByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetRecoverValueByUserId(
	request *SetRecoverValueByUserIdRequest,
) (*SetRecoverValueByUserIdResult, error) {
	callback := make(chan SetRecoverValueByUserIdAsyncResult, 1)
	go p.SetRecoverValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setMaxValueByStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetMaxValueByStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaxValueByStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaxValueByStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetMaxValueByStatusAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetMaxValueByStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetMaxValueByStatusAsync(
	request *SetMaxValueByStatusRequest,
	callback chan<- SetMaxValueByStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setMaxValueByStatus",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.SignedStatusBody != nil && *request.SignedStatusBody != "" {
        bodies["signedStatusBody"] = *request.SignedStatusBody
    }
    if request.SignedStatusSignature != nil && *request.SignedStatusSignature != "" {
        bodies["signedStatusSignature"] = *request.SignedStatusSignature
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setMaxValueByStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetMaxValueByStatus(
	request *SetMaxValueByStatusRequest,
) (*SetMaxValueByStatusResult, error) {
	callback := make(chan SetMaxValueByStatusAsyncResult, 1)
	go p.SetMaxValueByStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setRecoverIntervalByStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRecoverIntervalByStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverIntervalByStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverIntervalByStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRecoverIntervalByStatusAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRecoverIntervalByStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetRecoverIntervalByStatusAsync(
	request *SetRecoverIntervalByStatusRequest,
	callback chan<- SetRecoverIntervalByStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setRecoverIntervalByStatus",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.SignedStatusBody != nil && *request.SignedStatusBody != "" {
        bodies["signedStatusBody"] = *request.SignedStatusBody
    }
    if request.SignedStatusSignature != nil && *request.SignedStatusSignature != "" {
        bodies["signedStatusSignature"] = *request.SignedStatusSignature
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setRecoverIntervalByStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetRecoverIntervalByStatus(
	request *SetRecoverIntervalByStatusRequest,
) (*SetRecoverIntervalByStatusResult, error) {
	callback := make(chan SetRecoverIntervalByStatusAsyncResult, 1)
	go p.SetRecoverIntervalByStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setRecoverValueByStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRecoverValueByStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverValueByStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverValueByStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRecoverValueByStatusAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRecoverValueByStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetRecoverValueByStatusAsync(
	request *SetRecoverValueByStatusRequest,
	callback chan<- SetRecoverValueByStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setRecoverValueByStatus",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.SignedStatusBody != nil && *request.SignedStatusBody != "" {
        bodies["signedStatusBody"] = *request.SignedStatusBody
    }
    if request.SignedStatusSignature != nil && *request.SignedStatusSignature != "" {
        bodies["signedStatusSignature"] = *request.SignedStatusSignature
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setRecoverValueByStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetRecoverValueByStatus(
	request *SetRecoverValueByStatusRequest,
) (*SetRecoverValueByStatusResult, error) {
	callback := make(chan SetRecoverValueByStatusAsyncResult, 1)
	go p.SetRecoverValueByStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) deleteStaminaByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStaminaByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteStaminaByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) DeleteStaminaByUserIdAsync(
	request *DeleteStaminaByUserIdRequest,
	callback chan<- DeleteStaminaByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "deleteStaminaByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.StaminaName != nil && *request.StaminaName != "" {
        bodies["staminaName"] = *request.StaminaName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteStaminaByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) DeleteStaminaByUserId(
	request *DeleteStaminaByUserIdRequest,
) (*DeleteStaminaByUserIdResult, error) {
	callback := make(chan DeleteStaminaByUserIdAsyncResult, 1)
	go p.DeleteStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) recoverStaminaByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RecoverStaminaByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RecoverStaminaByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RecoverStaminaByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RecoverStaminaByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RecoverStaminaByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) RecoverStaminaByStampSheetAsync(
	request *RecoverStaminaByStampSheetRequest,
	callback chan<- RecoverStaminaByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "recoverStaminaByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.recoverStaminaByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) RecoverStaminaByStampSheet(
	request *RecoverStaminaByStampSheetRequest,
) (*RecoverStaminaByStampSheetResult, error) {
	callback := make(chan RecoverStaminaByStampSheetAsyncResult, 1)
	go p.RecoverStaminaByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) raiseMaxValueByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RaiseMaxValueByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RaiseMaxValueByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RaiseMaxValueByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RaiseMaxValueByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RaiseMaxValueByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) RaiseMaxValueByStampSheetAsync(
	request *RaiseMaxValueByStampSheetRequest,
	callback chan<- RaiseMaxValueByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "raiseMaxValueByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.raiseMaxValueByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) RaiseMaxValueByStampSheet(
	request *RaiseMaxValueByStampSheetRequest,
) (*RaiseMaxValueByStampSheetResult, error) {
	callback := make(chan RaiseMaxValueByStampSheetAsyncResult, 1)
	go p.RaiseMaxValueByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setMaxValueByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetMaxValueByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaxValueByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaxValueByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetMaxValueByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetMaxValueByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetMaxValueByStampSheetAsync(
	request *SetMaxValueByStampSheetRequest,
	callback chan<- SetMaxValueByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setMaxValueByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setMaxValueByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetMaxValueByStampSheet(
	request *SetMaxValueByStampSheetRequest,
) (*SetMaxValueByStampSheetResult, error) {
	callback := make(chan SetMaxValueByStampSheetAsyncResult, 1)
	go p.SetMaxValueByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setRecoverIntervalByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRecoverIntervalByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverIntervalByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverIntervalByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRecoverIntervalByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRecoverIntervalByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetRecoverIntervalByStampSheetAsync(
	request *SetRecoverIntervalByStampSheetRequest,
	callback chan<- SetRecoverIntervalByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setRecoverIntervalByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setRecoverIntervalByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetRecoverIntervalByStampSheet(
	request *SetRecoverIntervalByStampSheetRequest,
) (*SetRecoverIntervalByStampSheetResult, error) {
	callback := make(chan SetRecoverIntervalByStampSheetAsyncResult, 1)
	go p.SetRecoverIntervalByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) setRecoverValueByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRecoverValueByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverValueByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverValueByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRecoverValueByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRecoverValueByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) SetRecoverValueByStampSheetAsync(
	request *SetRecoverValueByStampSheetRequest,
	callback chan<- SetRecoverValueByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "setRecoverValueByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setRecoverValueByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) SetRecoverValueByStampSheet(
	request *SetRecoverValueByStampSheetRequest,
) (*SetRecoverValueByStampSheetResult, error) {
	callback := make(chan SetRecoverValueByStampSheetAsyncResult, 1)
	go p.SetRecoverValueByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2StaminaWebSocketClient) consumeStaminaByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ConsumeStaminaByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeStaminaByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeStaminaByStampTaskResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeStaminaByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ConsumeStaminaByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaWebSocketClient) ConsumeStaminaByStampTaskAsync(
	request *ConsumeStaminaByStampTaskRequest,
	callback chan<- ConsumeStaminaByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "stamina",
    		"component": "stamina",
    		"function": "consumeStaminaByStampTask",
            "contentType": "application/json",
    		"requestId": requestId,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.consumeStaminaByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2StaminaWebSocketClient) ConsumeStaminaByStampTask(
	request *ConsumeStaminaByStampTaskRequest,
) (*ConsumeStaminaByStampTaskResult, error) {
	callback := make(chan ConsumeStaminaByStampTaskAsyncResult, 1)
	go p.ConsumeStaminaByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
